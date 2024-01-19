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

func GetOnePwd(acctName string, userId string) *model.Password {
	var password model.Password
	result := dbConnection.Model(&model.Password{}).Where("account_name = ? AND user_id = ?", acctName, userId).First(&password)
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

func DeletePwd(acctName string, userId string) error {
	if err := dbConnection.Where("account_name = ? AND user_id = ?", acctName, userId).Delete(&model.Password{}).Error; err != nil {
		fmt.Println("Error when deleting password:", err)
		return err
	}
	fmt.Println("Successfully deleted password with account name:", acctName)
	return nil
}

func UpdatePwd(updatedPwd *model.Password, userId string) error {
	err := dbConnection.Where("account_name = ? AND user_id = ?", updatedPwd.AccountName, userId).Save(updatedPwd).Error
	if err != nil {
		fmt.Println("Error when updating password:", err)
		return err
	}
	fmt.Println("Successfully updated password")
	return nil
}
