package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/services"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Login(c echo.Context) error {
	var auth models.Auth

	if err := c.Bind(&auth); err != nil {
		errResponse := helpers.StatusInvalidated(err.Error())
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	response, errResponse := services.Login(auth)
	if errResponse != nil {
		return c.JSON(http.StatusUnauthorized, errResponse)
	}

	return c.JSON(http.StatusOK, response)
}

func Register(c echo.Context) error {
	var register models.Register

	if err := c.Bind(&register); err != nil {
		errResponse := helpers.StatusInvalidated(err.Error())
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	if errResponse := services.Register(register); errResponse != nil {
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func RefreshToken(c echo.Context) error {
	var token models.RefreshToken

	if err := c.Bind(&token); err != nil {
		errResponse := helpers.StatusInvalidated(err.Error())
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	tokens, errResponse := services.RefreshToken(token.RefreshToken)
	if errResponse != nil {
		return c.JSON(http.StatusUnauthorized, errResponse)
	}
	return c.JSON(http.StatusOK, tokens)
}
