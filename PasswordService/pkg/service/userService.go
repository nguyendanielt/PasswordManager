package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func ValidateUser(token string) (bool, uuid.UUID, error) {
	// create new http request to auth service with the token to verify in header
	url := fmt.Sprintf("%s%s", os.Getenv("AUTH_SERVICE_HOST"), "/api/user/authorization/validate")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Unable to create http request")
		return false, uuid.Nil, err
	}
	request.Header.Set("token", token)

	// send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request to auth service:", err)
		return false, uuid.Nil, err
	}

	// unauthorized user
	if response.StatusCode == http.StatusUnauthorized {
		return false, uuid.Nil, nil
	}

	resBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read response body")
		return false, uuid.Nil, err
	}

	authResponse := struct {
		Message string `json:"message"`
		UserId  string `json:"userid"`
	}{}

	if err := json.Unmarshal(resBodyBytes, &authResponse); err != nil {
		fmt.Println("Error decoding JSON")
		return false, uuid.Nil, err
	}

	uuidStrToUuid, _ := uuid.Parse(authResponse.UserId)
	return true, uuidStrToUuid, nil
}
