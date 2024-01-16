package helper

import (
	"encoding/json"
	"net/http"

	"authservice/pkg/model"
)

func GetReqBodyUser(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	var reqBodyUser model.User
	if err := json.NewDecoder(r.Body).Decode(&reqBodyUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &reqBodyUser, nil
}

func JsonSuccessResponse(w http.ResponseWriter, responseMap map[string]string) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonBytes))
}
