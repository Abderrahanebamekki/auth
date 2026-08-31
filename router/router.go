package router

import (
	"auth/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", handler.SingUp).Methods("POST")
	auth.HandleFunc("/login", handler.Login).Methods("POST")

	r.HandleFunc("/public/info", handler.Public).Methods("GET")
	r.HandleFunc("/protected/profile", handler.Protected).Methods("GET")
	return r
}
