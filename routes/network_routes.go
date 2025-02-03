package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterNetworkRoutes enregistre la route permettant de récupérer les statistiques réseau.
//
// Route HTTP :
// - GET /system/network : Retourne les statistiques réseau de la machine (bytes envoyés/reçus).
//
// Paramètres :
// - mux *http.ServeMux : Le routeur HTTP sur lequel enregistrer la route.
//
// Fonctionnalités :
// - Log chaque requête entrante.
// - Mesure le temps d'exécution de la requête.
// - Récupère les statistiques réseau via `monitor.GetNetworkStats`.
// - Retourne les données au format JSON ou une erreur en cas d'échec.
func RegisterNetworkRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/network", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête reçue
		utils.LogRequest(r)

		// Récupération des statistiques réseau
		data, err := monitor.GetNetworkStats()
		if err != nil {
			// Gestion de l'erreur en cas d'échec de récupération des données
			utils.HandleError(w, err, "Error getting network stats", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec les statistiques réseau
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
