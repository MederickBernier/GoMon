package routes

import "net/http"

// RegisterRoutes initialise et enregistre toutes les routes du serveur.
// Elle retourne un *http.ServeMux qui gère les routes définies.
//
// Routes enregistrées :
// - Base routes
// - CPU monitoring routes
// - Memory monitoring routes
// - Disk monitoring routes
// - Network monitoring routes
// - System information routes
// - Alerts routes
// - Ports monitoring routes
//
// Retourne :
// - *http.ServeMux : un gestionnaire de routes HTTP.
func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Enregistrement des routes spécifiques
	RegisterBaseRoutes(mux)
	RegisterCPURoutes(mux)
	RegisterMemoryRoutes(mux)
	RegisterDiskRoutes(mux)
	RegisterNetworkRoutes(mux)
	RegisterSystemRoutes(mux)
	RegisterAlertsRoutes(mux)
	RegisterPortsRoutes(mux)

	return mux
}
