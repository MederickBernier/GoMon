package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterPortsRoutes enregistre la route permettant de récupérer les ports en écoute.
//
// Route HTTP :
// - GET /system/ports : Retourne une liste des ports ouverts et en écoute sur la machine.
//
// Paramètres :
// - mux *http.ServeMux : Le routeur HTTP sur lequel enregistrer la route.
//
// Fonctionnalités :
// - Log chaque requête entrante.
// - Mesure le temps d'exécution de la requête.
// - Récupère la liste des ports en écoute via `monitor.GetListeningPorts`.
// - Retourne les données au format JSON ou une erreur en cas d'échec.
func RegisterPortsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/ports", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		// Récupération des ports en écoute
		data, err := monitor.GetListeningPorts()
		if err != nil {
			utils.HandleError(w, err, "Error getting listening ports", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
