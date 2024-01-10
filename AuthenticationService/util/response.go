package util

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, responseMap map[string]string) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonBytes))
}
