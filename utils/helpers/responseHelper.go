package helpers

import (
	"github.com/muhammetandic/go-backend/main/core/enums/errors"
	"github.com/muhammetandic/go-backend/main/core/models"
)

func StatusUnauthorized(message string) models.ErrorResponse {
	return models.ErrorResponse{
		Code:   errors.AuthorizationError,
		Status: errors.AuthorizationErrorStatus,
		Error:  message,
	}
}

func StatusUnauthenticated(message string) models.ErrorResponse {
	return models.ErrorResponse{
		Code:   errors.AuthenticationError,
		Status: errors.AuthenticationErrorStatus,
		Error:  message,
	}
}

func StatusUnvalidated(message string) models.ErrorResponse {
	return models.ErrorResponse{
		Code:   errors.ValidationError,
		Status: errors.ValidationErrorStatus,
		Error:  message,
	}
}

func StatusInternalServerError(message string) models.ErrorResponse {
	return models.ErrorResponse{
		Code:   errors.InternalServerError,
		Status: errors.InternalServerErrorStatus,
		Error:  message,
	}
}
