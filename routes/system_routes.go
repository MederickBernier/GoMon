package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterSystemRoutes enregistre les routes permettant d'obtenir des informations sur le système.
//
// Route HTTP :
//   - **GET /system/info** : Retourne des informations détaillées sur le système telles que l'OS, le nom d'hôte,
//     l'architecture, le temps d'activité et le nombre de cœurs du processeur.
//
// Paramètres :
// - **mux** (*http.ServeMux) : Multiplexeur HTTP utilisé pour enregistrer la route.
//
// Fonctionnalités :
// - **Log** chaque requête entrante via `utils.LogRequest`.
// - **Mesure** le temps d'exécution de la requête avec `utils.MeasureExecutionTime`.
// - **Récupère** les informations système à l'aide de `monitor.GetSystemInfo`.
// - **Retourne** une réponse JSON avec les données système ou une erreur en cas d'échec.
func RegisterSystemRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/info", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête reçue
		utils.LogRequest(r)

		// Récupération des informations système
		data, err := monitor.GetSystemInfo()
		if err != nil {
			// Gestion de l'erreur et envoi d'une réponse HTTP 500 en cas d'échec
			utils.HandleError(w, err, "Error getting system info", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec les informations système
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
