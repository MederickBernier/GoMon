package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterSystemRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/info", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetSystemInfo()
		if err != nil {
			utils.HandleError(w, err, "Error getting system info", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})
}
