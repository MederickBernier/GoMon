package monitor

import "github.com/shirou/gopsutil/v3/disk"

func GetDiskStats()(map[string]interface{},error){
	usage, err := disk.Usage("/")
	if err != nil{
		return nil,err
	}
	return map[string]interface{}{
		"total_space":usage.Total,
		"used_space":usage.Used,
		"free_space":usage.Free,
		"used_percent":usage.UsedPercent,
	},nil
}