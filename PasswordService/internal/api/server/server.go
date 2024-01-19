package server

import (
	"log"
	"net/http"

	"passwordservice/pkg/database"
)

func CreateServer() {
	database.InitDatabase()
	log.Fatal(http.ListenAndServe(":8081", newRouter()))
}
