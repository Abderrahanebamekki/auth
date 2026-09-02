package handler

import (
	"encoding/json"
	"net/http"
)

// Public godoc
// @Summary      Public info
// @Description  Returns a public welcome message that requires no authentication
// @Tags         Public
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /public/info [get]
func Public(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message": "Welcome stranger! This info is public.",
	}

	json.NewEncoder(w).Encode(response)

}
