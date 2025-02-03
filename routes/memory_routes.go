package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterMemoryRoutes enregistre la route permettant de récupérer les statistiques mémoire.
//
// Route HTTP :
// - GET /system/memory : Retourne les informations d'utilisation de la mémoire.
//
// Paramètres :
// - mux *http.ServeMux : Le routeur HTTP utilisé pour enregistrer la route.
//
// Fonctionnalités :
// - Log chaque requête entrante.
// - Mesure le temps d'exécution de la requête.
// - Récupère les statistiques mémoire via `monitor.GetMemoryStats`.
// - Retourne les données au format JSON ou une erreur en cas d'échec.
func RegisterMemoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/memory", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête entrante
		utils.LogRequest(r)

		// Récupération des statistiques mémoire
		data, err := monitor.GetMemoryStats()
		if err != nil {
			// Gestion de l'erreur en cas d'échec de récupération des données
			utils.HandleError(w, err, "Error getting memory stats", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec les statistiques mémoire
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
