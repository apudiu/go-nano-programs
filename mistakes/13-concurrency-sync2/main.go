package main

import "fmt"

func main() {
	chanSend()
	chanBuffSend()

	chanClose()
}

func chanSend() {
	i := 0

	ch := make(chan any)

	go func() {
		<-ch
		fmt.Printf("[send] I is: %d \n", i)
	}()

	i++
	ch <- 'x'
}

func chanBuffSend() {
	i := 0

	//ch := make(chan any, 1) // this makes race condition
	ch := make(chan any) // this doesn't as receive happens before the send

	go func() {
		i = 5
		<-ch
	}()

	ch <- 'x'
	fmt.Printf("[buff send] I is: %d \n", i)
}

func chanClose() {
	i := 0

	ch := make(chan any)

	go func() {
		<-ch
		fmt.Printf("[close] I is: %d \n", i)
	}()

	i++
	close(ch)
}
