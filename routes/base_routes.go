package routes

import (
	"net/http"

	"github.com/MederickBernier/GoMon/utils"
)

func RegisterBaseRoutes(mux *http.ServeMux){
	mux.HandleFunc("/monitor", utils.MeasureExecutionTime(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"monitor.html")
	}))
}