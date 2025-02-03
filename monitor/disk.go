package monitor

import "github.com/shirou/gopsutil/v3/disk"

// GetDiskStats récupère les statistiques d'utilisation du disque.
//
// Retourne :
// - **map[string]interface{}** : Contient les informations suivantes :
//   - `total_space` (uint64) : Espace disque total en octets.
//   - `used_space` (uint64) : Espace disque actuellement utilisé en octets.
//   - `free_space` (uint64) : Espace disque libre en octets.
//   - `used_percent` (float64) : Pourcentage de l'espace disque utilisé.
//
// - **error** : Retourne une erreur si la récupération des statistiques disque échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/disk.Usage("/")` pour récupérer les métriques d'utilisation du disque racine `/`.
// - Retourne un dictionnaire contenant les **statistiques du disque** sous une forme exploitable.
func GetDiskStats() (map[string]interface{}, error) {
	// Récupération des statistiques d'utilisation du disque pour la partition racine "/"
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	// Retourne un dictionnaire contenant les métriques disque
	return map[string]interface{}{
		"total_space":  usage.Total,       // Taille totale du disque en octets
		"used_space":   usage.Used,        // Espace disque utilisé en octets
		"free_space":   usage.Free,        // Espace disque libre en octets
		"used_percent": usage.UsedPercent, // Pourcentage de l'espace disque utilisé
	}, nil
}
