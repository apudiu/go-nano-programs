package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Starting timer at: %s", time.Now())
	timerFn()

	fmt.Printf("Starting ticker at: %s", time.Now())
	tickerFn()

}
