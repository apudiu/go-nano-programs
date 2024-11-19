package main

import (
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/internal/hardware"
	"log"
	"net/http"
	"time"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
}

func newServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,
	}

	s.mux.Handle("/", http.FileServer(http.Dir("./htmx/")))

	return s
}

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

	srv := newServer()
	log.Fatalln(
		http.ListenAndServe(":8000", &srv.mux),
	)
}
