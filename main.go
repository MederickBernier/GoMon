package main

import (
	"log"
	"net/http"

	"github.com/MederickBernier/GoMon/routes"
)

func main() {
	
	mux := routes.RegisterRoutes()

	port := "8080"
	log.Printf("Server is running at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, mux)
}
