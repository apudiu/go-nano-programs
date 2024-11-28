package main

import "fmt"

func main() {

	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		close(ch2)
	}()

	ch := ch1
	for v := range ch {
		fmt.Printf("-> %#v \n", v)
		// NOTICE: in next iteration it'll still print from ch1, due to expr once eval on range
		ch = ch2
	}

}
