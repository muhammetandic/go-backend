package services

import (
	"testing"

	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/db"
)

// Login with correct email and password
func TestLoginWithCorrectEmailAndPassword(t *testing.T) {
	db.Connect(":memory:")
	db.Migrate()

	// Create a new user with the correct email and password
	user := &models.Register{
		Username: "test@example.com",
		Password: "password123",
	}
	// Add the user to the repository
	Register(*user)

	// Create a new login request with the correct email and password
	auth := models.Auth{
		Username: "test@example.com",
		Password: "password123",
	}

	// Call the Login function with the login request
	response, errResponse := Login(auth)

	// Check that the response is not nil
	if response == nil {
		t.Fatal("Login response is nil")
	}

	// Check that the access token is not empty
	if response.AccessToken == "" {
		t.Fatal("Access token is empty")
	}

	// Check that the refresh token is not empty
	if response.RefreshToken == "" {
		t.Fatal("Refresh token is empty")
	}

	// Check that the access token expires at is not zero
	if response.AccessTokenExpiresAt.IsZero() {
		t.Fatal("Access token expires at is zero")
	}

	// Check that the refresh token expires at is not zero
	if response.RefreshTokenExpiresAt.IsZero() {
		t.Fatal("Refresh token expires at is zero")
	}

	// Check that the error response is nil
	if errResponse != nil {
		t.Fatalf("Error response is not nil: %s", errResponse.Error)
	}
}
