package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetAlerts() ([]string, error) {
	var alerts []string

	cpuUsage, err := cpu.Percent(500, false)
	if err != nil {
		return nil, fmt.Errorf("error getting CPU usage : %v", err)
	}
	if len(cpuUsage) > 0 && cpuUsage[0] > 90 {
		alerts = append(alerts, "High CPU usage: > 90 %")
	}

	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error getting memory usage: %v", err)
	}
	if memoryStats.UsedPercent > 85 {
		alerts = append(alerts, "High memory usage: > 85%")
	}

	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error getting disk usage: %v", err)
	}
	if diskStats.Free < diskStats.Total/10 {
		alerts = append(alerts, "Low Disk Space: < 10% free")
	}
	return alerts, nil
}
