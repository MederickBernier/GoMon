package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JSONResponse envoie une réponse HTTP au format JSON.
// Elle prend en paramètre `w` (le ResponseWriter), `data` (les données à encoder en JSON) et `status` (le code HTTP à retourner).
func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// LogRequest enregistre les informations d'une requête HTTP dans les logs.
// Elle affiche la méthode, l'URL et l'adresse IP du client.
func LogRequest(r *http.Request) {
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

// HandleError gère une erreur en loggant un message et en envoyant une réponse JSON avec le statut approprié.
// Elle prend en paramètre `w` (le ResponseWriter), `err` (l'erreur capturée), `message` (un message d'erreur personnalisé) et `status` (le code HTTP).
func HandleError(w http.ResponseWriter, err error, message string, status int) {
	log.Printf("Error: %s, Message: %s", err, message)
	JSONResponse(w, map[string]string{"error": message}, status)
}

// FormatUptime convertit une durée en secondes en une chaîne formatée HH:mm:ss.
// Elle prend en paramètre `seconds` (nombre total de secondes) et retourne une chaîne formatée.
func FormatUptime(seconds uint64) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02dh:%02dm:%02ds", hours, minutes, seconds)
}
