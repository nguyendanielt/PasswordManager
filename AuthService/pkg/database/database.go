package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"authservice/pkg/model"
	"authservice/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	// create table based on struct in userModel
	db.AutoMigrate(&model.User{})

	dbConnection = db
}

func GetUserByEmailAndPassword(email string, password string) *model.User {
	var user model.User
	result := dbConnection.Model(&model.User{}).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || !util.CompareHashAndPwdString(user.Password, password) {
		return nil
	}
	return &user
}

func AddUser(user *model.User) error {
	user.Password = util.GenerateHashedPwdString(user.Password)
	if err := dbConnection.Create(user).Error; err != nil {
		fmt.Println("Error when adding user:", err)
		return err
	}
	fmt.Println("Successfully added user")
	return nil
}
