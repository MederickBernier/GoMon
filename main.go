package main

import (
	"log"
	"net/http"

	"github.com/MederickBernier/GoMon/routes"
)

// main est le point d'entrée principal du programme.
// Cette fonction initialise les routes du serveur et démarre l'écoute sur le port défini.
func main() {
	// Initialisation du routeur avec les routes définies dans le package `routes`
	mux := routes.RegisterRoutes()

	// Définition du port sur lequel le serveur va écouter
	port := "8080"
	log.Printf("Server is running at http://localhost:%s\n", port)

	// Démarrage du serveur HTTP avec les routes configurées
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
