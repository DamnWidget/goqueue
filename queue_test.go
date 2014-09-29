/*
	Copyright (C) 2014  Oscar Campos <oscar.campos@member.fsf.org>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	See LICENSE file for more details.
*/

package goqueue_test

import (
	"runtime"

	queue "github.com/DamnWidget/goqueue"
)

var _ = Describe("Queue", func() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Describe("Call the New function", func() {
		Context("With a fixed size", func() {
			q := queue.New(10)
			It("Should return a pointer to a new Queue of that size", func() {
				Expect(q).ToNot(BeNil())
				Expect(q.Cap()).To(Equal(int64(10)))
			})
		})

		Context("With non size", func() {
			q := queue.New()
			It("Should return a pointer to a new Queue of unlimited size", func() {
				Expect(q).ToNot(BeNil())
				Expect(q.Cap()).To(Equal(int64(-1)))
			})
		})
	})

	Describe("Len method", func() {
		Describe("Should return the current Queue lenght", func() {
			Context("With ten elements", func() {
				q := queue.New()
				for i := 0; i < 10; i++ {
					q.Push(struct{}{})
				}
				It("Should return ten", func() {
					Expect(q.Len()).To(Equal(int64(10)))
				})
			})

			Context("With no elements", func() {
				q := queue.New()
				It("Should return zero", func() {
					Expect(q.Len()).To(BeZero())
				})
			})
		})
	})

	Describe("Push method", func() {
		Context("Push a new element at the end of the queue", func() {
			q := queue.New()
			q.Push(true)
			It("Should be a queue of unlimited capacity with one true element", func() {
				Expect(q.Cap()).To(Equal(int64(-1)))
				Expect(q.Len()).To(Equal(int64(1)))
				Expect(q.Pop()).To(BeTrue())
				Expect(q.Pop()).To(BeNil())
			})
		})

		Context("Push two elements at the end of the queue is FIFO", func() {
			q := queue.New()
			q.Push(true)
			q.Push(false)
			It("Should be an unlimited queue with two elements, first one true", func() {
				Expect(q.Cap()).To(Equal(int64(-1)))
				Expect(q.Len()).To(Equal(int64(2)))
				Expect(q.Pop()).To(BeTrue())
				Expect(q.Len()).To(Equal(int64(1)))
			})
		})

		Context("Push beyond limits is not allowed", func() {
			q := queue.New(1)
			q.Push(1)
			err := q.Push(2)
			It("Should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Thread Safety", func() {
		Context("Push is thread safe with two or more goroutines", func() {
			q := queue.New()
			for i := 0; i < 6; i++ {
				go func(q *queue.Queue) { q.Push(1) }(q)
			}
			It("Should be a queue of six element", func() {
				Expect(q.Len()).To(Equal(int64(6)))
			})
		})

		Context("Calling push and pop in goroutines is safe", func() {
			q := queue.New()
			for i := 1; i < 6001; i++ {
				go func(q *queue.Queue) { q.Push(1) }(q)
			}

			for i := 0; i < 3000; i++ {
				go func(q *queue.Queue) { q.Pop() }(q)
			}

			It("Should be a Queue of three thousand elements", func() {
				Expect(q.Len()).To(Equal(int64(3000)))
			})
		})
	})

	Describe("Values returns the whole Queue values", func() {
		q := queue.New()
		for i := 1; i < 6; i++ {
			q.Push(i * i)
		}
		It("Shouldn't mute the Queue", func() {
			v := q.Values()
			Expect(v).To(Equal([]interface{}{1, 4, 9, 16, 25}))
			Expect(q.Len()).To(Equal(int64(5)))
		})
	})

})
