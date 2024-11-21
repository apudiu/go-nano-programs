package main

import (
	"bytes"
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/config"
	"github.com/apudiu/go-nano-programs/hwmonitor/internal/hardware"
	"github.com/apudiu/go-nano-programs/hwmonitor/internal/server"
	"github.com/apudiu/go-nano-programs/hwmonitor/templates"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	srv := server.New()

	go func(s *server.Server) {
		for {
			sysInfo, err := hardware.GetSystemInfo()
			if err != nil {
				fmt.Println("system info error", err)
				break
			}

			cpuInfo, err := hardware.GetCpuInfo()
			if err != nil {
				fmt.Println("cpu info error", err)
				break
			}

			diskInfo, err := hardware.GetDiskInfo()
			if err != nil {
				fmt.Println("disk info error", err)
				break
			}

			tmpl, err := templates.GetTemplate("components/sections.gohtml")
			if err != nil {
				fmt.Println("Failed to parse template:", err)
				break
			}

			data := map[string]any{
				"now":      time.Now().Format(time.DateTime),
				"sysInfo":  template.HTML(sysInfo),
				"cpuInfo":  template.HTML(cpuInfo),
				"diskInfo": template.HTML(diskInfo),
			}

			buff := new(bytes.Buffer)
			if err2 := tmpl.Execute(buff, data); err2 != nil {
				fmt.Println("Failed to execute template:", err2)
				break
			}

			s.Broadcast(buff.Bytes())
			time.Sleep(3 * time.Second)
		}
	}(srv)

	log.Fatalln(
		http.ListenAndServe(":"+config.Conf.Port, &srv.Mux),
	)
}
