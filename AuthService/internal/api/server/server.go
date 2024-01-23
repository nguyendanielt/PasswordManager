package server

import (
	"log"
	"net/http"

	"authservice/internal/api/asyncmessaging"
	"authservice/pkg/database"
)

func CreateServer() {
	database.InitDatabase()
	asyncmessaging.ProducerSetup()
	log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
