package routes

import "net/http"

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	RegisterCPURoutes(mux)
	RegisterMemoryRoutes(mux)
	RegisterDiskRoutes(mux)
	RegisterNetworkRoutes(mux)
	RegisterSystemRoutes(mux)
	RegisterAlertsRoutes(mux)
	RegisterPortsRoutes(mux)
	RegisterTemperatureRoutes(mux)
	return mux
}
