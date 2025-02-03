package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/utils"
)

// RegisterBaseRoutes enregistre les routes de base pour le serveur.
//
// Routes HTTP :
// - **GET /monitor** : Sert la page HTML `monitor.html`, qui affiche l'interface de monitoring.
//
// Paramètres :
// - **mux** (*http.ServeMux) : Multiplexeur HTTP utilisé pour enregistrer la route.
//
// Fonctionnalités :
// - **Log** chaque requête entrante via `utils.MeasureExecutionTime`.
// - **Fournit** le fichier HTML statique `monitor.html` via `http.ServeFile`.
func RegisterBaseRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/monitor", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		// Servir le fichier HTML statique pour l'interface de monitoring
		http.ServeFile(w, r, "monitor.html")
	}))
}
