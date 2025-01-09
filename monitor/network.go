package monitor

import (
	"github.com/shirou/gopsutil/v3/net"
)

func GetNetworkStats() (map[string]interface{}, error) {
	ioCounters, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}
	if len(ioCounters) == 0 {
		return nil, nil
	}

	stats := ioCounters[0]

	return map[string]interface{}{
		"bytes_sent":     stats.BytesSent,
		"bytes_received": stats.BytesRecv,
		"packets_sent":   stats.PacketsSent,
		"packets_recv":   stats.PacketsRecv,
	}, nil
}
