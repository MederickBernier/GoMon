package routes

import "net/http"

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

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
