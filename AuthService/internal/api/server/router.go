package server

import (
	"net/http"

	"authservice/internal/api/handler"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/user").Subrouter()
	s.HandleFunc("/signup", handler.SignUpUser).Methods(http.MethodPost)
	s.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)
	s.HandleFunc("/authorization/validate", handler.ValidateToken).Methods(http.MethodGet)
	return r
}
