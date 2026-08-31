package handler

import (
	"encoding/json"
	"net/http"
)

func Public(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message": "Welcome stranger! This info is public.",
	}

	json.NewEncoder(w).Encode(response)

}
