package server

import (
	"log"
	"net/http"

	"passwordservice/internal/api/asyncmessaging"
	"passwordservice/pkg/database"
)

func CreateServer() {
	database.InitDatabase()
	asyncmessaging.ProducerSetup()
	log.Fatal(http.ListenAndServe(":8081", newRouter()))
}
