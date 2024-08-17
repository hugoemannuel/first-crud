package routes

import (
	"first-crud/controller"
	"first-crud/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func User() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controller.UserGetAll).Methods("GET")
	r.HandleFunc("/users/{id}", controller.UserFindOne).Methods("GET")
	r.Handle("/users", middleware.ValidCreate(http.HandlerFunc(controller.UserCreate))).Methods("POST")
	r.Handle("/users/{id}", middleware.ValidExclude(http.HandlerFunc(controller.UserExclude))).Methods("DELETE")
	r.Handle("/users/{id}", middleware.ValidUpdate(http.HandlerFunc(controller.Update))).Methods("PUT")
	return r
}
