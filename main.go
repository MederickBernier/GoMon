package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
)

func main(){
	// Define a basic route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		response := map[string]string{"Status":"OK"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Define a route to get CPU usage
	http.HandleFunc("/system/cpu", func(w http.ResponseWriter, r *http.Request){
		// get the CPU usage
		usage, err := cpu.Percent(0,false)
		if err != nil{
			http.Error(w, "Error getting CPU usage", http.StatusInternalServerError)
			return
		}

		// return the json response
		response := map[string]float64{"cpu_usage": usage[0]}
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