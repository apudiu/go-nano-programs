package hardware

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"runtime"
)

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

	output := fmt.Sprintf(
		"Hostname: %s\nTotal Memory: %s\nUsed Memory: %s\nOS: %s",
		hostInfo.Hostname, toGb(vmStat.Total), toGb(vmStat.Used), currentOs,
	)
	return output, nil
}

func GetCpuInfo() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("CPU: %s\nCores: %d", cpuInfo[0].ModelName, len(cpuInfo))
	return output, nil
}

func GetDiskInfo() (string, error) {
	diskInfo, err := disk.Usage("/")
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("Total Disk Space: %s\nFree Disk Space: %s", toGb(diskInfo.Total), toGb(diskInfo.Free))
	return output, nil
}

// helpers
func toGb(n uint64) string {
	return fmt.Sprintf("%.2fGB", float64(n)/float64(1024*1024*1024))
}
