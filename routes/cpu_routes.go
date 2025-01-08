package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterCPURoutes(mux *http.ServeMux){
	mux.HandleFunc("/system/cpu", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetCPUUsage()
		if err != nil {
			utils.HandleError(w, err, "Error getting CPU usage", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})

	mux.HandleFunc("/system/cpu/cores", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetCPUUsagePerCore()
		if err != nil {
			utils.HandleError(w, err, "Error getting CPU usage per core", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, map[string][]float64{"cores": data}, http.StatusOK)
	})
}