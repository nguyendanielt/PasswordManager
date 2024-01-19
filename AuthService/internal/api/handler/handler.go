package handler

import (
	"net/http"

	"authservice/internal/api/helper"
	"authservice/pkg/database"
	"authservice/pkg/service"

	"github.com/google/uuid"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	reqBodyUser, err := helper.GetReqBodyUser(w, r)
	if err != nil {
		return
	}
	if err := database.AddUser(reqBodyUser); err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}
	helper.JsonSuccessResponse(w, map[string]string{
		"message": "Successfully created account",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	reqBodyUser, err := helper.GetReqBodyUser(w, r)
	if err != nil {
		return
	}
	user := database.GetUserByEmailAndPassword(reqBodyUser.Email, reqBodyUser.Password)
	if user == nil {
		http.Error(w, "Login data could not be fetched", http.StatusNotFound)
		return
	}
	token, err := service.GenerateJwt(user)
	if err != nil {
		http.Error(w, "Error occurred when generating JWT", http.StatusInternalServerError)
		return
	}
	helper.JsonSuccessResponse(w, map[string]string{
		"message": "Login was successful",
		"token":   token,
	})
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	userId := service.ValidateJwt(token)
	if userId == uuid.Nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	helper.JsonSuccessResponse(w, map[string]string{
		"message": "Authorization successful",
		"userid":  userId.String(),
	})
}
