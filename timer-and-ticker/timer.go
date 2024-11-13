package main

import (
	"fmt"
	"time"
)

func timerFn() {
	timer := time.NewTimer(3 * time.Second)

	fmt.Printf("Timer expired at %s\n", <-timer.C)
}
