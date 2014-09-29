# goqueue

[![Build Status](https://travis-ci.org/DamnWidget/goqueue.png)](https://travis-ci.org/DamnWidget/sublime-text)

GoQueue is a simple thread safe empty interface Queue implementation in Golang

## Installation

To install GoQueue just use `go get` in your preferred interpreter

`go get github.com/DamnWidget/goqueue`

## Usage

GoQueue is pretty straightforward to use, you just have to create a new Queue
and push or pop whatever element you want from or to it.

Queues can have a maximum capacity size or be unlimited, to create an unlimited
queue just call the `New` method with no parameters:

```go
// create a new unlimited Queue
q := goqueue.New()
```

If you want to create a sized Queue, then pass the desired size of the queue as
first and only one parameters to the `New` function (if more than one parameter
is passed, only the first one is taken into account)

```go
// create a new Queue with a maximum size of 1000 elements
q := goqueue.New(1000)
```

...