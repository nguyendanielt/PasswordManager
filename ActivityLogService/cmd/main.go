package main

import (
	"fmt"

	"activitylogservice/internal/api/asyncmessaging"
	"activitylogservice/pkg/database"
)

func main() {
	database.InitDatabase()
	err := asyncmessaging.ReadActivityMessages()
	if err != nil {
		fmt.Println("Message consumer failed", err)
	}
}
