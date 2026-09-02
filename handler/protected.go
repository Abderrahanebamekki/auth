package handler

import (
	"auth/middleware"

	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserContextKey).([]byte)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "User data not found in context"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userData)
}
