package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/monitor"
	"github.com/MederickBernier/GoMon/utils"
)

func RegisterDiskRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/system/disk", func(w http.ResponseWriter, r *http.Request){
		utils.LogRequest(r)

		data,err := monitor.GetDiskStats()
		if err != nil{
			utils.HandleError(w,err,"Error getting disk stats", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w,data,http.StatusOK)
	})
}
