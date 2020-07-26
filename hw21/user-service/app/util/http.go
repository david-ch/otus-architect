package util

import (
	"encoding/json"
	"net/http"
)

// Resp writes response with specified code and data
func Resp(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// ErrorResp is the standard format of error response in the app
type ErrorResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// FromError creates ErrorResp from an error
func FromError(e error) ErrorResp {
	return ErrorResp{Message: e.Error()}
}
