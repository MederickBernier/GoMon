package monitor

import (
	"github.com/shirou/gopsutil/v3/net"
)

// GetNetworkStats récupère les statistiques globales du réseau.
//
// Retourne :
// - **map[string]interface{}** : Contient les statistiques réseau suivantes :
//   - `bytes_sent` (uint64) : Nombre total d'octets envoyés.
//   - `bytes_received` (uint64) : Nombre total d'octets reçus.
//   - `packets_sent` (uint64) : Nombre total de paquets envoyés.
//   - `packets_recv` (uint64) : Nombre total de paquets reçus.
//
// - **error** : Retourne une erreur si la récupération des statistiques échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/net.IOCounters(false)` pour récupérer les **statistiques réseau globales**.
// - Vérifie si des données ont été obtenues avant de les retourner.
// - Retourne un dictionnaire avec les métriques réseau importantes.
func GetNetworkStats() (map[string]interface{}, error) {
	// Récupération des statistiques globales du réseau
	ioCounters, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	// Vérification si aucune donnée n'est disponible
	if len(ioCounters) == 0 {
		return nil, nil
	}

	// Sélection du premier résultat (statistiques globales)
	stats := ioCounters[0]

	// Retourne un dictionnaire avec les statistiques réseau
	return map[string]interface{}{
		"bytes_sent":     stats.BytesSent,   // Nombre total d'octets envoyés
		"bytes_received": stats.BytesRecv,   // Nombre total d'octets reçus
		"packets_sent":   stats.PacketsSent, // Nombre total de paquets envoyés
		"packets_recv":   stats.PacketsRecv, // Nombre total de paquets reçus
	}, nil
}
