package controllers

import "github.com/gorilla/mux"

func NewRouter(userHandler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	return r
}