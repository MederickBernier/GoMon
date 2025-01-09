package monitor

import (
	"runtime"

	"github.com/MederickBernier/GoMon/utils"
	"github.com/shirou/gopsutil/v3/host"
)

func GetSystemInfo() (map[string]interface{}, error) {
	info, err := host.Info()
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"hostname":  info.Hostname,
		"os":        info.OS,
		"platform":  info.Platform,
		"uptime":    utils.FormatUptime(info.Uptime),
		"arch":      runtime.GOARCH,
		"num_cores": runtime.NumCPU(),
	}, nil
}
