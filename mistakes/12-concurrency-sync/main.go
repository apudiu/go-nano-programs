package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    withAtomic()
    withMutex()
    withChannel()
}

func withAtomic() {
    var n int64 = 0

    go func() {
        atomic.AddInt64(&n, 1)
    }()

    go func() {
        atomic.AddInt64(&n, 1)
    }()

    // if we need n here after goroutines are finished, we can wait using sync.WaitGroup
}

func withMutex() {
    var m sync.Mutex
    var n int64 = 0

    go func() {
        m.Lock()
        defer m.Unlock()
        n++
    }()

    go func() {
        m.Lock()
        defer m.Unlock()
        n++
    }()
}

func withChannel() {
    var n int64 = 0

    c := make(chan int)

    go func() {
        c <- 0
    }()

    go func() {
        c <- 0
    }()

    <-c
    n++
    <-c
    n++

    fmt.Printf("%#v \n", n)
}
