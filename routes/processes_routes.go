package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterTopProcessesRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/top-processes", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetTopProcesses(5)
		if err != nil {
			utils.HandleError(w, err, "Error getting processes", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})
}
