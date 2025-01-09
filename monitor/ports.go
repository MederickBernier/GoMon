package monitor

import "github.com/shirou/gopsutil/v3/net"

func GetListeningPorts() ([]map[string]interface{}, error) {
	connections, err := net.Connections("inet")
	if err != nil {
		return nil, err
	}

	var ports []map[string]interface{}
	for _, conn := range connections {
		if conn.Status == "LISTEN" {
			ports = append(ports, map[string]interface{}{
				"port":    conn.Laddr.Port,
				"address": conn.Laddr.IP,
				"process": conn.Pid,
			})
		}
	}

	return ports, nil
}
