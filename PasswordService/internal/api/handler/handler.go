package handler

import (
	"net/http"

	"passwordservice/internal/api/helper"
	"passwordservice/pkg/database"

	"github.com/google/uuid"
)

func GetPassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	password := database.GetOnePwd(reqBody.AccountName, userId)
	if password == nil {
		http.Error(w, "Password could not be fetched", http.StatusNotFound)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"password": *password,
	})
}

func GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	passwords := database.GetAllPwds(userId)
	if passwords == nil {
		http.Error(w, "Password could not be fetched", http.StatusNotFound)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"password": passwords,
	})
}

func AddPassword(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(r.Header.Get("userid"))
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	reqBody.UserId = userId
	if err := database.AddPwd(reqBody); err != nil {
		http.Error(w, "Failed to add password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully added password",
	})
}

func DeletePassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userid")
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	if err := database.DeletePwd(reqBody.AccountName, userId); err != nil {
		http.Error(w, "Failed to delete password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully deleted password",
	})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(r.Header.Get("userid"))
	reqBody, err := helper.GetReqBody(w, r)
	if err != nil {
		return
	}
	reqBody.UserId = userId
	if err := database.UpdatePwd(reqBody, userId.String()); err != nil {
		http.Error(w, "Failed to update password", http.StatusBadRequest)
		return
	}
	helper.JsonSuccessResponse(w, map[string]interface{}{
		"message": "Successfully updated password",
	})
}
