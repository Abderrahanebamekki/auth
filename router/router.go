package router

import (
	"auth/handler"
	"auth/middleware"

	_ "auth/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", handler.SingUp).Methods("POST")
	auth.HandleFunc("/login", handler.Login).Methods("POST")
	auth.HandleFunc("/signout", handler.SignOut).Methods("POST")

	r.HandleFunc("/public/info", handler.Public).Methods("GET")

	r.HandleFunc("/protected/profile", middleware.AuthGuard(handler.Profile)).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
