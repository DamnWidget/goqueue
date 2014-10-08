# GoQueue

[![Build Status](https://travis-ci.org/DamnWidget/goqueue.png)](https://travis-ci.org/DamnWidget/sublime-text)

GoQueue is a simple thread safe empty interface Queue implementation in Golang

## Installation

To install GoQueue just use `go get` in your preferred interpreter

`go get github.com/DamnWidget/goqueue`

## Usage

GoQueue is pretty straightforward to use, you just have to create a new Queue
and push or pop whatever element you want from or to it.

### Unlimited Queues

Queues can have a maximum capacity size or be unlimited, to create an unlimited
queue just call the `New` method with no parameters:

```go
// create a new unlimited Queue
q := goqueue.New()
```

### Sized Queues

If you want to create a sized Queue, then pass the desired size of the queue as
first and only one parameter to the `New` function (if more than one parameter
is passed, only the first one is taken into account)

```go
// create a new Queue with a maximum size of 1000 elements
q := goqueue.New(1000)
```

### Queue size operations

The Queue size and capacity properties are exported via the `Len` and `Cap`
methods respectively.

```go
q := goqueue.New(1000)
fmt.Println(q.Len())
fmt.Println(q.Cap())
```
Will output:

```
0
1000
```

### Push elements to the end of the Queue

To push an element to the end of the Queue we just call the method `Push` using
the element as unique argument.

```go
q := goqueue.New()
err := q.Push("Test")
```

Remember that any type of element can be pushed to the GoQueue, including your
own defined types.

```go
type Customer struct {
    name        string
    customer_id int
}
err := q.Push(&Customer{"Test Customer", 1})
```

If we try to push an element beyond the size limit of the Queue the `Push`
method returns an error so you should always check the returned value. If
the `Push` is successful then, it returns `nil`.

### Pop elements from the end of the Queue

Elements are extracted from the end of the Queue (as this Queue is a FIFO data
structure).

```go
customer = q.Pop()
```

If the Queue is empty, `Pop` returns nil so you will always check the returned
value and use in your logic

```go
q := goqueue.Queue()
// push things here
...

for {
    item := q.Pop()
    if item == nil {
        break
    }
    // do something useful here
}
```

### Return all Queue elements at once

We can get all the elements from a Queue without extract them (and without
mutate the Queue) if we need to do so using the `Values` method.

```go
for item := range q.Values() {
    // do something clever here with item
}
```

## Thread Safe Operations

All the operations in the Queue are thread safe so we can share a Queue
instance safely between any number of `goroutines` as we need.
