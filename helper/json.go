package helper

import (
	"encoding/json"
	"golang-with-mongo/internal/payload"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := payload.WebResponse{
		Code:    code,
		Status:  "error",
		Message: message,
		Data:    nil,
	}

	PanicIfError(json.NewEncoder(w).Encode(response))
}

func WriteSuccessResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := payload.WebResponse{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	}

	PanicIfError(json.NewEncoder(w).Encode(response))
}
