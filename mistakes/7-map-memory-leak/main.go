package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()

	// NOTICE: memory is not fully reduced, indicates memory leak
	printAlloc()

	runtime.KeepAlive(m)
}

func randBytes() [128]byte {
	return [128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

// In GO maps can't shrink so even after clearing the items will not
// give back total used memory by that map.
// The solution could be using a pointer for the value which yields lower memory usage and
// create new maps in specific intervals so previous one gets collected by GC
