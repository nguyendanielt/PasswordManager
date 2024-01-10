package main

import (
	"log"
	"net/http"

	"authservice/database"
	"authservice/handler"
	"authservice/repository"
	"authservice/service"

	"github.com/gorilla/mux"
)

func main() {
	db := database.CreateConnection()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	jwtService := service.NewJwtService()
	handler := handler.NewHandler(userService, jwtService)

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/signup", handler.SignUpUser).Methods(http.MethodPost)
	s.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)
	s.HandleFunc("/authorization/validate", handler.ValidateToken).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
