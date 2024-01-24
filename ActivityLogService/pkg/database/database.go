package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"activitylogservice/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

var dbConnection *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// establish database connection and initialize GORM instance
	dsn := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// enable uuid-ossp extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	// create table based on struct in activityModel
	db.AutoMigrate(&model.Activity{})

	dbConnection = db
}

func AddMessage(message *sarama.ConsumerMessage) error {
	var activity model.Activity
	err := json.Unmarshal(message.Value, &activity)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return err
	}
	if err := dbConnection.Create(&activity).Error; err != nil {
		fmt.Println("Error when adding activity:", err)
		return err
	}
	fmt.Println("Successfully added activity")
	return nil
}
