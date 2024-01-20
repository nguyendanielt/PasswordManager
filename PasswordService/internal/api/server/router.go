package server

import (
	"net/http"

	"passwordservice/internal/api/handler"
	"passwordservice/internal/api/middleware"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/password").Subrouter()
	s.HandleFunc("/one", handler.GetPassword).Methods(http.MethodGet)
	s.HandleFunc("/all", handler.GetAllPasswords).Methods(http.MethodGet)
	s.HandleFunc("/add", handler.AddPassword).Methods(http.MethodPost)
	s.HandleFunc("/delete", handler.DeletePassword).Methods(http.MethodDelete)
	s.HandleFunc("/update", handler.UpdatePassword).Methods(http.MethodPut)
	r.Use(middleware.Authorization)
	return r
}
