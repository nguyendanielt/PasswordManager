package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"passwordservice/pkg/model"

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
	// create table based on struct in passwordModel
	db.AutoMigrate(&model.Password{})

	dbConnection = db
}

func GetPwd(passwordId string, userId string) *model.Password {
	var password model.Password
	result := dbConnection.Model(&model.Password{}).Where("password_id = ? AND user_id = ?", passwordId, userId).First(&password)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &password
}

func GetAllPwds(userId string) []model.Password {
	var passwords []model.Password
	result := dbConnection.Model(&model.Password{}).Where("user_id = ?", userId).Find(&passwords)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return passwords
}

func AddPwd(password *model.Password) error {
	if err := dbConnection.Create(password).Error; err != nil {
		fmt.Println("Error when adding password:", err)
		return err
	}
	fmt.Println("Successfully added password")
	return nil
}

func DeletePwd(passwordId string, userId string) error {
	var password model.Password
	result := dbConnection.Model(&model.Password{}).Where("password_id = ?", passwordId).First(&password)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Password to delete does not exist")
		return result.Error
	}
	result = dbConnection.Where("password_id = ? AND user_id = ?", passwordId, userId).Delete(&model.Password{})
	if result.Error != nil {
		fmt.Println("Error when deleting password:", result.Error)
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows affected in delete")
	}
	fmt.Println("Successfully deleted password")
	return nil
}

func UpdatePwd(updatedPwd *model.Password) error {
	result := dbConnection.Where("password_id = ? AND user_id = ?", updatedPwd.PasswordId.String(), updatedPwd.UserId.String()).UpdateColumns(updatedPwd)
	if result.Error != nil {
		fmt.Println("Error when updating password:", result.Error)
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows affected in update")
	}
	fmt.Println("Successfully updated password")
	return nil
}
