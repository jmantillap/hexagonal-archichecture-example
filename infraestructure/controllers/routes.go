package controllers

import "github.com/gorilla/mux"

func NewRouter(userHandler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.UpdateUser).Methods("PUT")

	return r
}