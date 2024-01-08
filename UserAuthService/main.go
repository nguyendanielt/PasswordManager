package main

import (
	"net/http"

	"userauthservice/database"
	"userauthservice/handler"
	"userauthservice/repository"
	"userauthservice/service"

	"github.com/gorilla/mux"
)

func main() {
	db := database.CreateConnection()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	jwtService := service.NewJwtService()
	handler := handler.NewHandler(userService, jwtService)

	r := mux.NewRouter()
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/signup", handler.SignUpUser).Methods(http.MethodPost)
	s.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)
	s.HandleFunc("/validate", handler.ValidateToken).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
