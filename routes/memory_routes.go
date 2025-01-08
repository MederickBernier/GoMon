package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterMemoryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/memory", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		data, err := monitor.GetMemoryStats()
		if err != nil {
			utils.HandleError(w, err, "Error getting memory stats", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, data, http.StatusOK)
	})
}
