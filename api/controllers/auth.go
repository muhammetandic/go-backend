package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/services"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Login(c *gin.Context) {
	var auth models.Auth

	if err := c.ShouldBindJSON(&auth); err != nil {
		errResponse := helpers.StatusUnvalidated(err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	response, errResponse := services.Login(auth)
	if errResponse != nil {
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var register models.Register

	if err := c.ShouldBindJSON(&register); err != nil {
		errResponse := helpers.StatusUnvalidated(err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if errResponse := services.Register(register); errResponse != nil {
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
