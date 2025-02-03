package monitor

import (
	"runtime"

	"github.com/MederickBernier/GoMon/utils"
	"github.com/shirou/gopsutil/v3/host"
)

// GetSystemInfo récupère et retourne les informations du système hôte.
//
// Retourne :
// - **map[string]interface{}** : Contient les informations suivantes :
//   - `hostname` (string) : Nom de l'hôte du système.
//   - `os` (string) : Nom du système d'exploitation.
//   - `platform` (string) : Plateforme sur laquelle le programme s'exécute.
//   - `uptime` (string) : Durée de fonctionnement du système formatée en HH:MM:SS.
//   - `arch` (string) : Architecture du processeur (ex: amd64, arm64).
//   - `num_cores` (int) : Nombre total de cœurs du processeur.
//
// - **error** : Retourne une erreur si la récupération des informations échoue.
//
// Fonctionnalités :
// - Utilise `gopsutil/host.Info()` pour récupérer les informations système.
// - Utilise `runtime.GOARCH` pour obtenir l'architecture du processeur.
// - Utilise `runtime.NumCPU()` pour déterminer le nombre de cœurs.
// - Convertit l'uptime brut en format lisible via `utils.FormatUptime()`.
func GetSystemInfo() (map[string]interface{}, error) {
	// Récupération des informations sur l'hôte
	info, err := host.Info()
	if err != nil {
		return nil, err
	}

	// Construction de la réponse contenant les informations système
	return map[string]interface{}{
		"hostname":  info.Hostname,                   // Nom de l'hôte
		"os":        info.OS,                         // Système d'exploitation
		"platform":  info.Platform,                   // Plateforme (Windows, Linux, etc.)
		"uptime":    utils.FormatUptime(info.Uptime), // Temps depuis le dernier démarrage formaté
		"arch":      runtime.GOARCH,                  // Architecture du processeur
		"num_cores": runtime.NumCPU(),                // Nombre de cœurs CPU
	}, nil
}
