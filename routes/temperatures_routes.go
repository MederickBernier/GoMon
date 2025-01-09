package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterTemperatureRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/temperature", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		data, err := monitor.GetTemperatures()
		if err != nil {
			utils.JSONResponse(w,map[string]string{
				"error": "Temperature monitoring failed",
				"message": err.Error(),
			},http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})
}
