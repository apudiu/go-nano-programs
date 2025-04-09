package main

func main() {

    ch := make(chan Foo, 1)
    s := "x" // escapes to heap as points by Foo
    bar := Foo{s: &s}
    ch <- bar
}

type Foo struct{ s *string }

// check heap alloc using cmd: go build -gcflags "-m=2" .

// General reasons for heap allocation
// 1. If a local variable is too large to fit on the stack.
// 2. If the size of a local variable is unknown. For example, s := make([]int, 10) may not escape to the heap, but s := make([]int, n) will, because its size is based on a variable.
// 3. If the backing array of a slice is reallocated using append.

// In general, sharing down (in stack frames) stays on the stack, whereas sharing up escapes to the heap
