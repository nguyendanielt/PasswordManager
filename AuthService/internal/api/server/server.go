package server

import (
	"log"
	"net/http"

	"authservice/pkg/database"
)

func CreateServer() {
	database.InitDatabase()
	log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
