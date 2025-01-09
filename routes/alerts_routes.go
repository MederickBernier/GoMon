package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterAlertsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		alerts, err := monitor.GetAlerts()
		if err != nil {
			utils.HandleError(w, err, "Error getting alerts", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, map[string][]string{"alerts": alerts}, http.StatusOK)
	})
}
