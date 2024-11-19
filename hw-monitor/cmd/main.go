package main

import (
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/internal/hardware"
	"time"
)

func main() {
	go func() {
		for {
			sysInfo, err := hardware.GetSystemInfo()
			if err != nil {
				fmt.Println(err)
			}

			cpuInfo, err := hardware.GetCpuInfo()
			if err != nil {
				fmt.Println(err)
			}

			diskInfo, err := hardware.GetDiskInfo()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(sysInfo)
			fmt.Println(cpuInfo)
			fmt.Println(diskInfo)

			time.Sleep(3 * time.Second)
		}
	}()

	time.Sleep(5 * time.Minute)
}
