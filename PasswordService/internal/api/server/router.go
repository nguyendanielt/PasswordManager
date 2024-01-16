package server

import (
	"net/http"

	"passwordservice/internal/api/handler"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/password").Subrouter()
	s.HandleFunc("/tmp1", handler.GetPassword).Methods(http.MethodGet)
	s.HandleFunc("/tmp2", handler.GetAllPasswords).Methods(http.MethodGet)
	s.HandleFunc("/tmp3", handler.AddPassword).Methods(http.MethodPost)
	s.HandleFunc("/tmp4", handler.DeletePassword).Methods(http.MethodDelete)
	s.HandleFunc("/tmp5", handler.EditPassword).Methods(http.MethodPatch)
	return r
}
