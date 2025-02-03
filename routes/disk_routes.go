package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterDiskRoutes enregistre la route permettant de récupérer les statistiques d'utilisation du disque.
//
// Route HTTP :
// - **GET /system/disk** : Retourne les statistiques d'utilisation du disque, incluant :
//   - L'espace total
//   - L'espace utilisé
//   - L'espace libre
//   - Le pourcentage d'utilisation du disque
//
// Paramètres :
// - **mux** (*http.ServeMux) : Multiplexeur HTTP utilisé pour enregistrer la route.
//
// Fonctionnalités :
// - **Log** chaque requête entrante via `utils.LogRequest`.
// - **Mesure** le temps d'exécution de la requête avec `utils.MeasureExecutionTime`.
// - **Récupère** les statistiques d'utilisation du disque à l'aide de `monitor.GetDiskStats`.
// - **Retourne** une réponse JSON avec les données du disque ou une erreur en cas d'échec.
func RegisterDiskRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/disk", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête entrante
		utils.LogRequest(r)

		// Récupération des statistiques du disque
		data, err := monitor.GetDiskStats()
		if err != nil {
			// Gestion de l'erreur et envoi d'une réponse HTTP 500 en cas d'échec
			utils.HandleError(w, err, "Error getting disk stats", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec les statistiques du disque
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
