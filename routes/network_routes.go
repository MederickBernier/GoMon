package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterNetworkRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/network", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetNetworkStats()
		if err != nil {
			utils.HandleError(w, err, "Error getting network stats", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	}))
}
