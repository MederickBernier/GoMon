package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterPortsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/ports", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetListeningPorts()
		if err != nil {
			utils.HandleError(w, err, "Error getting listening ports", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})
}
