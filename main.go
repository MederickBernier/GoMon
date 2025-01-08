package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main(){
	// Define a basic route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		response := map[string]string{"Status":"OK"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Start the server
	port := "8080"
	fmt.Printf("Server is running at http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}