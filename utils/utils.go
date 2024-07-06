package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test-dep-prod/types"

	"github.com/go-playground/validator/v10"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func ValidateRegisterUserPayload(payload *types.RegisterUserPayload) error {
	err := validator.New().Struct(*payload)
	if err != nil {
		errorMessages := make([]string, 0)
		for _, validationError := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, validationError.Field())
		}

		if len(errorMessages) == 1 {
			log.Println("this field is required:", errorMessages)
		} else if len(errorMessages) > 1 {
			log.Println("these field are required:", errorMessages)
		} else {
			log.Println("validation error:", errorMessages)
		}

		return fmt.Errorf("%v is required", errorMessages[0])
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errorResponse := struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}{
		Success: false,
		Error:   err.Error(),
	}

	WriteJSON(w, status, errorResponse)
}
