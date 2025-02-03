package monitor

import "github.com/shirou/gopsutil/v3/mem"

// GetMemoryStats récupère les statistiques de la mémoire vive du système.
//
// Retourne :
// - **map[string]interface{}** : Contient les informations suivantes :
//   - `total_memory` (uint64) : Quantité totale de mémoire vive (RAM) disponible en octets.
//   - `used_memory` (uint64) : Quantité de mémoire actuellement utilisée en octets.
//   - `free_memory` (uint64) : Quantité de mémoire libre en octets.
//   - `used_percent` (float64) : Pourcentage de la mémoire utilisée.
//
// - **error** : Retourne une erreur si la récupération des statistiques mémoire échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/mem.VirtualMemory()` pour obtenir les métriques de mémoire vive.
// - Retourne un dictionnaire contenant les **statistiques RAM** sous une forme exploitable.
func GetMemoryStats() (map[string]interface{}, error) {
	// Récupération des statistiques de la mémoire RAM
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Retourne un dictionnaire contenant les métriques mémoire
	return map[string]interface{}{
		"total_memory": memoryStats.Total,       // RAM totale en octets
		"used_memory":  memoryStats.Used,        // RAM utilisée en octets
		"free_memory":  memoryStats.Free,        // RAM libre en octets
		"used_percent": memoryStats.UsedPercent, // Pourcentage de RAM utilisée
	}, nil
}
