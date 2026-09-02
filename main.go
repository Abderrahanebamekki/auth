package main

import (
	"log"
	"net/http"
	"os"

	"auth/router"

	"github.com/joho/godotenv"
)

// @title           Auth Service API
// @version         1.0
// @description     Authentication service using Supabase Auth. Provides signup, login, logout and protected profile endpoints.
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	port := os.Getenv("PORT")

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
