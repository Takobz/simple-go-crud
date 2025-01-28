package helpers

import (
	"simple-go-crud/dtos"
)

func CreateErrorResponse(message string) dtos.ErrorResponse {
	return dtos.ErrorResponse{
		Message: message,
	}
}
