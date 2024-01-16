package database

import (
	"log"
	"os"

	"passwordservice/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
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
	// create table based on struct in passwordModel
	db.AutoMigrate(&model.Password{})

	dbConnection = db
}

func GetPasswordByAcountName(acctName string) {

}

func GetAllPasswords(userId uuid.UUID) {

}

func AddPassword(password *model.Password) {

}

func DeletePassword(password *model.Password) {

}

func EditPassword() {

}
