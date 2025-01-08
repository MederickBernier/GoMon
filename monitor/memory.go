package monitor

import "github.com/shirou/gopsutil/v3/mem"

func GetMemoryStats() (map[string]interface{}, error) {
	memoryStats, err := mem.VirtualMemory()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_memory": memoryStats.Total,
		"used_memory":  memoryStats.Used,
		"free_memory":  memoryStats.Free,
		"used_percent": memoryStats.UsedPercent,
	}, nil
}
