package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Authenticate(c *gin.Context) {
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
	c.Next()
}

func Authorize(c *gin.Context) {
	ctx := context.Background()
	username, _ := c.Get("username")
	userRepo := repository.UserRepo(db.Instance)
	uname, _ := username.(string)
	user := userRepo.Get(&model.User{Email: uname}, ctx)
	println(user.Roles)
}
