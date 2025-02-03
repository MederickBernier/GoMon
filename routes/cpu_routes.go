package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

// RegisterCPURoutes enregistre les routes permettant de récupérer les statistiques d'utilisation du CPU.
//
// Routes HTTP :
// - **GET /system/cpu** : Retourne l'utilisation globale du CPU en pourcentage.
// - **GET /system/cpu/cores** : Retourne l'utilisation du CPU pour chaque cœur.
//
// Paramètres :
// - **mux** (*http.ServeMux) : Multiplexeur HTTP utilisé pour enregistrer les routes.
//
// Fonctionnalités :
// - **Log** chaque requête entrante via `utils.LogRequest`.
// - **Mesure** le temps d'exécution des requêtes avec `utils.MeasureExecutionTime`.
// - **Récupère** l'utilisation globale du CPU via `monitor.GetCPUUsage`.
// - **Récupère** l'utilisation par cœur via `monitor.GetCPUUsagePerCore`.
// - **Retourne** une réponse JSON avec les statistiques CPU ou une erreur en cas d'échec.
func RegisterCPURoutes(mux *http.ServeMux) {
	// Route pour l'utilisation globale du CPU
	mux.HandleFunc("/system/cpu", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête entrante
		utils.LogRequest(r)

		// Récupération de l'utilisation globale du CPU
		data, err := monitor.GetCPUUsage()
		if err != nil {
			// Gestion de l'erreur et envoi d'une réponse HTTP 500 en cas d'échec
			utils.HandleError(w, err, "Error getting CPU usage", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec l'utilisation globale du CPU
		utils.JSONResponse(w, data, http.StatusOK)
	}))

	// Route pour l'utilisation du CPU par cœur
	mux.HandleFunc("/system/cpu/cores", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Log de la requête entrante
		utils.LogRequest(r)

		// Récupération de l'utilisation CPU par cœur
		data, err := monitor.GetCPUUsagePerCore()
		if err != nil {
			// Gestion de l'erreur et envoi d'une réponse HTTP 500 en cas d'échec
			utils.HandleError(w, err, "Error getting CPU usage per core", http.StatusInternalServerError)
			return
		}

		// Envoi de la réponse JSON avec l'utilisation CPU par cœur
		utils.JSONResponse(w, map[string][]float64{"cores": data}, http.StatusOK)
	}))
}
