package api

import (
	"encoding/json"
	"http_server/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Handle POST request to add a new exhibition
	var exhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&exhibition)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	data.Add(exhibition)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
