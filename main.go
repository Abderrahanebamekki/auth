package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	supabase "github.com/supabase-community/supabase-go"
)

var client *supabase.Client

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	port := os.Getenv("PORT")

	var err error
	client, err = supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatal("failed to init supabase client:", err)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
