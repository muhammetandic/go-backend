package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Authenticate(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		errorResponse := helpers.StatusUnauthorized("no authorization header provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
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
	path = strings.TrimPrefix(path, "/api/")
	pattern := `\/.*`
	regEx := regexp.MustCompile(pattern)
	path = regEx.ReplaceAllString(path, "")

	username, _ := c.Get("username")

	ctx := c.Request.Context()

	userRepo := repository.NewUserRepo()
	uname, _ := username.(string)
	user := userRepo.GetWithRelated(&model.User{Email: uname}, "Roles.Role.Privileges", ctx)

	canDo := CanDo(method, path, user.Roles[:])
	if !canDo {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.StatusUnauthorized("you are not authorized to access the resource"))
	}

	c.Next()
}

func CanDo(method string, path string, roles []model.UserToRole) bool {
	for _, role := range roles {
		for _, privilege := range role.Role.Privileges {
			if privilege.Endpoint == path {
				switch method {
				case "GET":
					if privilege.CanRead {
						return true
					}
				case "POST":
					if privilege.CanInsert {
						return true
					}
				case "PUT":
					if privilege.CanUpdate {
						return true
					}
				case "DELETE":
					if privilege.CanDelete {
						return true
					}
				}
			}
		}
	}
	return false
}
