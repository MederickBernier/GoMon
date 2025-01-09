package monitor

import (
	"errors"
	"runtime"

	"github.com/shirou/gopsutil/v3/host"
)

func GetTemperatures() ([]map[string]interface{}, error) {
	if runtime.GOOS != "linux" {
		return nil, errors.New("temperature monitoring is not supported on this OS")
	}

	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}

	var temps []map[string]interface{}
	for _, sensor := range sensors {
		temps = append(temps, map[string]interface{}{
			"sensor":     sensor.SensorKey,
			"temperature": sensor.Temperature,
		})
	}

	return temps, nil
}