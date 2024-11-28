package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc2()

	for i := 0; i < n; i++ {
		m[i] = randBytes2()
	}
	printAlloc2()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()

	// NOTICE: memory is not fully reduced, indicates memory leak
	printAlloc2()

	runtime.KeepAlive(m)
}

func randBytes2() *[128]byte {
	return &[128]byte{}
}

func printAlloc2() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

//todo: 4
