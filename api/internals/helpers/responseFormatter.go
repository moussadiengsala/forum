package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
This function has as purpose formatted the
response from the server & send it the client
*/
func ResponseFormatter(data any, w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
	log.Println("Request URL:", r.URL.Path, "Status Code:", statusCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	return
}
