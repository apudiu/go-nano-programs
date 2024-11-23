package main

import (
	"fmt"
	"runtime"
)

func main() {
	foos := make([]Foo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	// notice memory is not free, other 998 items still in memory
	//two := keepFirstTwoElementsOnlyProblem(foos)

	two := keepFirstTwoElementsOnly(foos)

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

type Foo struct {
	v []byte
}

func keepFirstTwoElementsOnlyProblem(foos []Foo) []Foo {
	// slicing just returns a new slice but underlying array still exists
	return foos[:2]
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	// here we're copying 2 items from existing slice to a new slice with len 2
	// after we're done there's no reference to existing slice (so to array)
	// so GC will collect existing slice

	s := make([]Foo, 2)
	copy(s, foos)
	return s
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
