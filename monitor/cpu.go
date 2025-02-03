package monitor

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCPUUsage récupère l'utilisation globale du CPU sur une courte période.
//
// Retourne :
// - **map[string]float64** : Contient une seule clé `cpu_usage` avec :
//   - `cpu_usage` (float64) : Pourcentage d'utilisation globale du CPU.
//
// - **error** : Retourne une erreur si la récupération des statistiques CPU échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/cpu.Percent(interval, false)` avec un délai de **500ms**.
// - Ne récupère **que l'utilisation globale du CPU**, pas par cœur.
// - Retourne les valeurs sous forme de dictionnaire clé-valeur.
func GetCPUUsage() (map[string]float64, error) {
	// Récupération de l'utilisation globale du CPU sur 500ms
	usage, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return nil, err
	}

	// Retourne un dictionnaire contenant le pourcentage d'utilisation CPU
	return map[string]float64{"cpu_usage": usage[0]}, nil
}

// GetCPUUsagePerCore récupère l'utilisation du CPU pour chaque cœur individuellement.
//
// Retourne :
// - **[]float64** : Un tableau contenant le pourcentage d'utilisation de chaque cœur du CPU.
// - **error** : Retourne une erreur si la récupération des statistiques CPU échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/cpu.Percent(interval, true)` avec un délai de **500ms**.
// - Retourne **un tableau avec l'utilisation CPU par cœur**, ce qui est utile pour voir la charge répartie.
func GetCPUUsagePerCore() ([]float64, error) {
	// Récupération de l'utilisation CPU par cœur sur 500ms
	usage, err := cpu.Percent(500*time.Millisecond, true)
	if err != nil {
		return nil, err
	}

	// Retourne un tableau contenant l'utilisation CPU pour chaque cœur
	return usage, nil
}
