package handler

import (
	"auth/middleware"

	"net/http"
)

// Profile godoc
// @Summary      Get current user profile
// @Description  Returns the authenticated user's profile data (requires a valid token)
// @Tags         Protected
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} object
// @Failure      401 {object} map[string]string
// @Router       /protected/profile [get]
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
