package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	buff := bytes.NewReader(make([]byte, 1024*5))

	//syncReader(buff)
	asyncReader(buff)
}

func asyncReader(r io.Reader) {
	var count int64 = 0
	workers := 5

	ch := make(chan []byte, workers)

	var wg sync.WaitGroup
	wg.Add(workers)

	fmt.Printf("%#v \n", 1)

	// read & send to chan for precessing
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("err: %#v \n", err)
			return
		}

		fmt.Printf("%#v \n", 6)
		ch <- b
	}

	fmt.Printf("%#v \n", 5)

	// execute tasks with (n) workers
	// when receive from chan
	for i := 0; i < workers; i++ {
		fmt.Printf("%#v \n", 2)
		go func(workerNo int) {
			defer wg.Done()
			fmt.Printf("%#v W(%d) \n", 3, workerNo)

			for buf := range ch {
				v := task(buf)
				atomic.AddInt64(&count, int64(v))
				fmt.Printf("%#v W(%d) \n", 4, workerNo)
			}

			// different impl
			//v := task(<-ch)
			//atomic.AddInt64(&count, int64(v))
			//fmt.Printf("%#v W(%d) \n", 4, workerNo)

		}(i)
	}

	fmt.Printf("%#v \n", 7)
	// after read all, close chan
	close(ch)
	fmt.Printf("%#v \n", 8)
	// wait for tasks to be completed
	wg.Wait()

	fmt.Printf("%#v \n", 9)

	fmt.Printf("Async Tasks: %#v \n", count)
}

func syncReader(buff io.Reader) {
	count := 0
	for {
		b := make([]byte, 1024)
		_, err := buff.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("err: %#v \n", err)
			return
		}
		count += task(b)
	}

	fmt.Printf("Sync Tasks: %#v \n", count)
}

func task(b []byte) int {
	time.Sleep(time.Second)
	return 1
}
