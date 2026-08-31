package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access token required",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

}
