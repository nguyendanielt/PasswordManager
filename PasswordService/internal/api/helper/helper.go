package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"passwordservice/internal/api/asyncmessaging"
	"passwordservice/pkg/model"

	"github.com/google/uuid"
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

func SendActivityProducerMessage(id string, activityType string) {
	userId, _ := uuid.Parse(id)
	activityJson, err := json.Marshal(model.Activity{
		UserId:       userId,
		ActivityType: activityType,
		DateAndTime:  time.Now(),
	})
	if err != nil {
		fmt.Println("Error occurred during marshal")
		return
	}
	asyncmessaging.SendActivityMessage(string(activityJson))
}
