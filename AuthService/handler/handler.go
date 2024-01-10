package handler

import (
	"encoding/json"
	"net/http"

	"authservice/model"
	"authservice/service"
	"authservice/util"

	"github.com/google/uuid"
)

type Handler struct {
	userService *service.UserService
	jwtService  *service.JwtService
}

func NewHandler(userService *service.UserService, jwtService *service.JwtService) *Handler {
	return &Handler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (h *Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var reqBodyUser model.User
	if err := json.NewDecoder(r.Body).Decode(&reqBodyUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.userService.CreateUser(&reqBodyUser); err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	util.SuccessResponse(w, map[string]string{
		"message": "Successfully created account",
	})
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqBodyUser model.User
	if err := json.NewDecoder(r.Body).Decode(&reqBodyUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := h.userService.GetUser(&reqBodyUser)
	if user == nil {
		http.Error(w, "Login data could not be fetched", http.StatusInternalServerError)
		return
	}

	token, err := h.jwtService.GenerateJwt(user)
	if err != nil {
		http.Error(w, "Error occurred when generating JWT", http.StatusInternalServerError)
		return
	}

	util.SuccessResponse(w, map[string]string{
		"message": "Login was successful",
		"token":   token,
	})
}

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	userId := h.jwtService.ValidateJwt(token)
	if userId == uuid.Nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	util.SuccessResponse(w, map[string]string{
		"message": "Authorization successful",
		"userid":  userId.String(),
	})
}
