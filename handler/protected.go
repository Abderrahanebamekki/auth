package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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
	token := strings.TrimPrefix(authHeader, "Bearer ")

	req, _ := http.NewRequest("GET", os.Getenv("SUPABASE_URL")+"/auth/v1/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid or expired token",
		})
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}
