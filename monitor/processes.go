package monitor

import "github.com/shirou/gopsutil/v3/process"

func GetTopProcesses(limit int) ([]map[string]interface{}, error) {
	var topProcesses []map[string]interface{}

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, p := range processes {
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryInfo()
		name, err := p.Name()
		if err != nil {
			name = "unknown"
		}
		if cpu > 1 || mem != nil {
			topProcesses = append(topProcesses, map[string]interface{}{
				"pid":         p.Pid,
				"name":        name,
				"cpu_percent": cpu,
				"memory":      mem.RSS,
			})
		}
		if len(topProcesses) >= limit {
			break
		}
	}
	return topProcesses, nil
}
