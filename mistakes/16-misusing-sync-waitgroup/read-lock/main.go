package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//usingMutex()
	//usingChan()
	usingSyncCond()
}

/**
  The main issue — and what makes this a terrible implementation—is the busy loop.
  Each listener goroutine keeps looping until its donation goal is met,
  which wastes a lot of CPU cycles and makes the CPU usage gigantic.
*/

func usingMutex() {
	type Donation struct {
		mu      sync.RWMutex
		balance int
	}

	donation := &Donation{}

	f := func(goal int) {
		donation.mu.RLock()
		defer donation.mu.RUnlock()

		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal met \n", donation.balance)
	}

	go f(10)
	go f(20)

	go func() {
		for {
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()

	time.Sleep(time.Second * 2)
}

/**
  The Issue - Each message is received by a single goroutine.
  Therefore, the first goroutine didn’t receive the $10 message in this example,
  but the second one did. Only a channel closure event can be broadcast to multiple goroutines.
  But here we don’t want to close the channel,
  because then the updater goroutine couldn’t send messages.

  There’s another issue with using channels in this situation.
  The listener goroutines return whenever their donation goal is met.
  Hence, the updater goroutine has to know when all the listeners stop receiving messages
  to the channel. Otherwise, the channel will eventually become full and block the sender.
  This will lead to a deadlock.
  A possible solution could be to add a sync.WaitGroup to the mix,
  but doing so would make the solution more complex.
*/

func usingChan() {
	type Donation struct {
		ch      chan int
		balance int
	}

	donation := &Donation{
		ch: make(chan int),
	}

	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal met \n", balance)
				return
			}
		}
	}

	go f(5)
	go f(7)

	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}
}

/**
 * This is the most optimal solution, but it has a problem: if
 * no receiver receives the broadcast, then the msg is missed
 */

func usingSyncCond() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}
	donation := &Donation{cond: sync.NewCond(&sync.Mutex{})}

	f := func(goal int) {
		donation.cond.L.Lock()
		defer donation.cond.L.Unlock()

		for donation.balance < goal {
			donation.cond.Wait()
		}

		fmt.Printf("$%d goal met \n", donation.balance)
	}

	go f(5)
	go f(7)

	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()

		donation.cond.Broadcast()
	}
}
