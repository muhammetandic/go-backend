package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/core/enums/errors"
	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/services"
)

func Login(c *gin.Context) {
	var auth models.Auth
	var errResponse models.ErrorResponse

	if err := c.ShouldBindJSON(&auth); err != nil {
		errResponse.Status = errors.ValidationErrorStatus
		errResponse.Code = errors.ValidationError
		errResponse.Error = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	response, err := services.Login(auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var register models.Register
	var errResponse models.ErrorResponse

	if err := c.ShouldBindJSON(&register); err != nil {
		errResponse.Status = errors.ValidationErrorStatus
		errResponse.Code = errors.ValidationError
		errResponse.Error = err.Error()
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if errResponse := services.Register(register); errResponse != nil {
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}
