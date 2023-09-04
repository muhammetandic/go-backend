package middleware

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"

	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		auth := req.Header.Get("Authorization")
		if auth == "" {
			errorResponse := helpers.StatusUnauthorized("no authorization header provided")
			return c.JSON(http.StatusForbidden, errorResponse)
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			errorResponse := helpers.StatusUnauthorized("could not find bearer token in authorization header")
			return c.JSON(http.StatusForbidden, errorResponse)
		}

		user, err := jwtAuth.ValidateToken(token)
		if err != nil {
			errorResponse := helpers.StatusUnauthorized(err.Error())
			return c.JSON(http.StatusForbidden, errorResponse)
		}

		c.Set("username", user.Username)
		next(c)
		return nil
	}
}

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		method := req.Method

		path := req.URL.Path
		path = strings.Replace(path, "/api/", "", -1)
		pattern := `\/.*`
		regEx := regexp.MustCompile(pattern)
		path = regEx.ReplaceAllString(path, "")

		username := c.Get("username")

		ctx := context.Background()

		userRepo := repository.NewUserRepo()
		uname, _ := username.(string)
		user := userRepo.GetWithRelated(&model.User{Email: uname}, "Roles.Role.Privileges", ctx)

		canDo := CanDo(method, path, user.Roles[:])
		if !canDo {
			return c.JSON(http.StatusUnauthorized, helpers.StatusUnauthorized("you are not authorized to access the resource"))
		}

		next(c)
		return nil
	}
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
