package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tickerFn() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	alertTimerActive := false
	var alertChan <-chan time.Time

	go func() {
		for {
			select {
			case <-ticker.C:
				cpuUsage := rand.Intn(100) + 100 // simulated high CPU usage
				fmt.Printf("CPU Usage: %d%% \n", cpuUsage)

				if cpuUsage > 80 {
					// if alert is not already active, activate it
					if !alertTimerActive {
						fmt.Println("High CPU usage detected, activating alert timer")

						alertTimerActive = true

						alertTimer := time.NewTimer(10 * time.Second)
						alertChan = alertTimer.C
					}
				} else {
					// stop alert timer as cpu usage decreased
					if alertTimerActive {
						fmt.Println("CPU usage is normal, de-activating alert timer")
						alertTimerActive = false
						alertChan = nil
					}

				}

			case <-alertChan:
				if alertTimerActive {
					fmt.Println("[ALERT]: CPU usage is high for 10 seconds")

					// deactivate the timer
					alertTimerActive = false
					alertChan = nil
				}
			}

		}
	}()

	// block indefinitely
	select {}

}
