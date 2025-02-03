package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// GetAlerts vérifie l'état des ressources système (CPU, mémoire, disque) et génère des alertes si des seuils critiques sont atteints.
//
// Retourne :
// - **[]string** : Une liste d'alertes sous forme de messages si des seuils critiques sont dépassés.
//   - "High CPU usage: > 90%" si l'utilisation CPU dépasse 90%.
//   - "High memory usage: > 85%" si l'utilisation de la mémoire dépasse 85%.
//   - "Low Disk Space: < 10% free" si l'espace disque disponible est inférieur à 10%.
//
// - **error** : Retourne une erreur si la récupération des données système échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/cpu.Percent()` pour surveiller la charge CPU.
// - Utilise `gopsutil/mem.VirtualMemory()` pour surveiller l'utilisation mémoire.
// - Utilise `gopsutil/disk.Usage("/")` pour surveiller l'utilisation disque.
// - Ajoute des messages d'alerte si les seuils définis sont dépassés.
func GetAlerts() ([]string, error) {
	var alerts []string

	// Vérification de l'utilisation du CPU
	cpuUsage, err := cpu.Percent(500, false)
	if err != nil {
		return nil, fmt.Errorf("error getting CPU usage: %v", err)
	}
	if len(cpuUsage) > 0 && cpuUsage[0] > 90 {
		alerts = append(alerts, "High CPU usage: > 90%")
	}

	// Vérification de l'utilisation de la mémoire
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error getting memory usage: %v", err)
	}
	if memoryStats.UsedPercent > 85 {
		alerts = append(alerts, "High memory usage: > 85%")
	}

	// Vérification de l'espace disque disponible
	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error getting disk usage: %v", err)
	}
	if diskStats.Free < diskStats.Total/10 {
		alerts = append(alerts, "Low Disk Space: < 10% free")
	}

	// Retourne les alertes détectées
	return alerts, nil
}
