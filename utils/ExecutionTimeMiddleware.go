package utils

import (
	"log"
	"net/http"
	"time"
)

// MeasureExecutionTime est un middleware qui mesure le temps d'exécution d'un gestionnaire HTTP.
// Il prend en paramètre `next` (une fonction http.HandlerFunc représentant le gestionnaire suivant) et retourne une fonction middleware.
// Cette fonction enregistre dans les logs le temps écoulé entre le début et la fin de l'exécution de la requête.
func MeasureExecutionTime(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()          // Capture du temps de début
		next(w, r)                   // Exécution du gestionnaire suivant
		elapsed := time.Since(start) // Calcul du temps écoulé
		log.Printf("Endpoint %s took %s", r.URL.Path, elapsed)
	}
}
