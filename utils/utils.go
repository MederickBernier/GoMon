package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func LogRequest(r *http.Request) {
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func HandleError(w http.ResponseWriter, err error, message string, status int) {
	log.Printf("Error: %s, Message: %s", err, message)
	JSONResponse(w, map[string]string{"error": message}, status)
}
