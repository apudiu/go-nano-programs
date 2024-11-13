package main

import (
	"log"
	"runtime"
)

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	b2Mb := func(b uint64) uint64 {
		return b / 1000 / 1000
	}

	log.Printf("Heap: %d mb, Total(so far): %d mb, Heap+Stack: %d mb \n", b2Mb(m.Alloc), b2Mb(m.TotalAlloc), b2Mb(m.Sys))
}

func main() {
	log.Println("Memory usage before")
	printMemoryUsage()

	list := make([]int, 2_000_000)
	for i := range list {
		list[i] = i
	}

	log.Println("Memory usage after")
	printMemoryUsage()
}
