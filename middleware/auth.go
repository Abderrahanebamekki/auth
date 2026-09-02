package middleware

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		supabaseURL := os.Getenv("SUPABASE_URL")
		supabaseKey := os.Getenv("SUPABASE_KEY")

		req, err := http.NewRequest("GET", supabaseURL+"/auth/v1/user", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Failed to create verification request",
			})
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("apikey", supabaseKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid or expired token",
			})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid or expired token",
			})
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Failed to read user data",
			})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserContextKey, body)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
