package api

import (
	"encoding/json"
	"http_server/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id != "" {
		// Handle fetching a specific exhibition by ID if needed
		id, err := strconv.Atoi(id)
		if err != nil || id < 0 || id >= len(data.GetAll()) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
			return
		}
		json.NewEncoder(w).Encode(data.GetAll()[id])
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
