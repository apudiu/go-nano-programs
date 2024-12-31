package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    wg := sync.WaitGroup{}
    var v int64

    maxCount := 10

    for i := 0; i < maxCount; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            atomic.AddInt64(&v, 1)
        }()
    }

    wg.Wait()
    fmt.Printf("%#v \n", v)
}

//todo: add 9.12 example
