package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterAlertsRoutes enregistre la route permettant de récupérer la liste des alertes système.
//
// Route HTTP :
// - **GET /alerts** : Retourne une liste des alertes système détectées.
//
// Paramètres :
// - **mux** (*http.ServeMux) : Multiplexeur HTTP utilisé pour enregistrer la route.
//
// Fonctionnalités :
// - **Log** chaque requête entrante via `utils.LogRequest`.
// - **Mesure** le temps d'exécution avec `utils.MeasureExecutionTime`.
// - **Récupère** les alertes via `monitor.GetAlerts`.
// - **Retourne** une réponse JSON contenant les alertes ou une erreur en cas d'échec.
func RegisterAlertsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/alerts", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête entrante
		utils.LogRequest(r)

		// Récupération des alertes système
		alerts, err := monitor.GetAlerts()
		if err != nil {
			// Gestion de l'erreur et envoi d'une réponse HTTP 500 en cas d'échec
			utils.HandleError(w, err, "Error getting alerts", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec la liste des alertes
		utils.JSONResponse(w, map[string][]string{"alerts": alerts}, http.StatusOK)
	}))
}
