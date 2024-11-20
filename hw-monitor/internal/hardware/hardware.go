package hardware

import (
	"bytes"
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/templates"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"html/template"
	"runtime"
	"strconv"
)

const megabyteDiv uint64 = 1024 * 1024
const gigabyteDiv uint64 = megabyteDiv * 1024

func GetSystemInfo() (string, error) {
	currentOs := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostInfo, err := host.Info()
	if err != nil {
		return "", err
	}

	tmpl, err := templates.GetTemplate("components/system.gohtml")
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"currentOs":         currentOs,
		"platform":          hostInfo.Platform,
		"hostname":          hostInfo.Hostname,
		"processCount":      strconv.FormatUint(hostInfo.Procs, 10),
		"memoryTotal":       strconv.FormatUint(vmStat.Total/megabyteDiv, 10) + " MB",
		"memoryFree":        strconv.FormatUint(vmStat.Free/megabyteDiv, 10) + " MB",
		"memoryUsedPercent": strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
	}

	buff := bytes.NewBufferString("")
	if err2 := tmpl.Execute(buff, data); err2 != nil {
		return "", err2
	}

	return buff.String(), nil
}

func GetCpuInfo() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting CPU info", err)

	}
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return "", err
	}

	// just to make 2 groups out of available
	firstCpus := percentage[:len(percentage)/2]
	secondCpus := percentage[len(percentage)/2:]

	tmpl, err := templates.GetTemplate("components/cpu.gohtml", template.FuncMap{
		"calculateCpuIndex": func(a int) int {
			return a + len(firstCpus)
		},
	})
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"cores":      len(cpuInfo),
		"modelName":  cpuInfo[0].ModelName,
		"family":     cpuInfo[0].Family,
		"frequency":  cpuInfo[0].Mhz,
		"firstCpus":  firstCpus,
		"secondCpus": secondCpus,
	}

	buff := bytes.NewBufferString("")
	if err2 := tmpl.Execute(buff, data); err2 != nil {
		return "", err2
	}

	return buff.String(), nil
}

func GetDiskInfo() (string, error) {
	diskInfo, err := disk.Usage("/")
	if err != nil {
		return "", err
	}

	html := "<div class='disk-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html = html + "<tr><td>Total disk space:</td><td>" + strconv.FormatUint(diskInfo.Total/gigabyteDiv, 10) + " GB</td></tr>"
	html = html + "<tr><td>Used disk space:</td><td>" + strconv.FormatUint(diskInfo.Used/gigabyteDiv, 10) + " GB</td></tr>"
	html = html + "<tr><td>Free disk space:</td><td>" + strconv.FormatUint(diskInfo.Free/gigabyteDiv, 10) + " GB</td></tr>"
	html = html + "<tr><td>Percentage disk space usage:</td><td>" + strconv.FormatFloat(diskInfo.UsedPercent, 'f', 2, 64) + "%</td></tr>"
	return html, nil
}
