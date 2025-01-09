package utils

import (
	"log"
	"net/http"
	"time"
)

func MeasureExecutionTime(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Printf("Endpoint %s took %s", r.URL.Path, elapsed)
	}
}
