package helper

import (
	"encoding/json"
	"net/http"

	"passwordservice/pkg/model"
)

func GetReqBody(w http.ResponseWriter, r *http.Request) (*model.Password, error) {
	var reqBody model.Password
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &reqBody, nil
}

func JsonSuccessResponse(w http.ResponseWriter, responseMap map[string]interface{}) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonBytes))
}
