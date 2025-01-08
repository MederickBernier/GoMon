package monitor

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUUsage() (map[string]float64, error) {
	usage, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return nil, err
	}
	return map[string]float64{"cpu_usage": usage[0]}, nil
}

func GetCPUUsagePerCore() ([]float64, error) {
	usage, err := cpu.Percent(500*time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	return usage, nil
}
