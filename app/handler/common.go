package main

import (
	"encoding/json"
	"net/http"
)

// respondJSON makes response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, error := json.Marshal(payload)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
	}

	w.Header().Set("Content=Type", "application/sion")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
