package routes

import "net/http"

func RegisterRoutes() *http.ServeMux{
	mux := http.NewServeMux()

	RegisterCPURoutes(mux)
	RegisterMemoryRoutes(mux)
	
	return mux
}