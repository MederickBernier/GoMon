package monitor

import "github.com/shirou/gopsutil/v3/net"

// GetListeningPorts récupère la liste des ports actuellement en écoute sur la machine.
//
// Retourne :
// - **[]map[string]interface{}** : Une liste contenant les informations suivantes pour chaque port en écoute :
//   - `port` (int) : Numéro du port en écoute.
//   - `address` (string) : Adresse IP associée au port.
//   - `process` (int) : ID du processus utilisant le port.
//
// - **error** : Retourne une erreur si la récupération des connexions réseau échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/net.Connections("inet")` pour récupérer la liste des connexions réseau actives.
// - Filtre les connexions pour ne conserver que celles ayant le statut `"LISTEN"`.
// - Retourne un tableau de ports sous forme de **dictionnaires** avec les détails associés.
func GetListeningPorts() ([]map[string]interface{}, error) {
	// Récupération de toutes les connexions réseau actives (IPv4 et IPv6)
	connections, err := net.Connections("inet")
	if err != nil {
		return nil, err
	}

	// Initialisation du tableau de ports en écoute
	var ports []map[string]interface{}

	// Parcours des connexions pour filtrer celles en écoute (LISTEN)
	for _, conn := range connections {
		if conn.Status == "LISTEN" {
			ports = append(ports, map[string]interface{}{
				"port":    conn.Laddr.Port, // Numéro du port
				"address": conn.Laddr.IP,   // Adresse IP locale associée
				"process": conn.Pid,        // ID du processus qui écoute sur ce port
			})
		}
	}

	// Retourne la liste des ports ouverts et en écoute
	return ports, nil
}
