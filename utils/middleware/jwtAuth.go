package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Authorize(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		errorResponse := helpers.StatusUnauthorized("no authorization header provided")
		c.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		errorResponse := helpers.StatusUnauthorized("could not find bearer token in authorization header")
		c.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
		return
	}

	user, err := jwtAuth.ValidateToken(token)
	if err != nil {
		errorResponse := helpers.StatusUnauthorized(err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, errorResponse)
		return
	}

	c.Set("username", user.Username)
	c.Set("userId", user.Uid)
	c.Set("role", user.UserType)
	c.Next()
}
