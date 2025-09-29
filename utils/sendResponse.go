package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}

func SendError(writer http.ResponseWriter, statusCode int, errMsg string) {
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(errMsg)
}
