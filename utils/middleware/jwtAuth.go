package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"

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
	method := c.Request.Method
	path := c.FullPath()
	path = strings.Replace(path, "/api/", "", -1)
	username, _ := c.Get("username")

	ctx := context.Background()

	userRepo := repository.UserRepo(db.Instance)
	uname, _ := username.(string)
	user := userRepo.GetWithRelated(&model.User{Email: uname}, "Roles.Role.Privileges", ctx)

	canDo := CanDo(method, path, user.Roles[:])
	if !canDo {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.StatusUnauthorized("you are not authorized to access the resource"))
	}

	c.Next()
}

func CanDo(method string, path string, roles []model.UserToRole) bool {
	switch method {
	case "GET":
		return slices.ContainsFunc(roles, func(role model.UserToRole) bool {
			return role.Role.Privileges.Endpoint == path && role.Role.Privileges.CanRead
		})
	case "POST":
		return slices.ContainsFunc(roles, func(role model.UserToRole) bool {
			return role.Role.Privileges.Endpoint == path && role.Role.Privileges.CanInsert
		})
	case "PUT":
		return slices.ContainsFunc(roles, func(role model.UserToRole) bool {
			return role.Role.Privileges.Endpoint == path && role.Role.Privileges.CanUpdate
		})
	case "DELETE":
		return slices.ContainsFunc(roles, func(role model.UserToRole) bool {
			return role.Role.Privileges.Endpoint == path && role.Role.Privileges.CanDelete
		})
	default:
		return false
	}
}
