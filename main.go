package main

import (
	"log"
	"net/http"
	"os"

	"auth/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	port := os.Getenv("PORT")

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
