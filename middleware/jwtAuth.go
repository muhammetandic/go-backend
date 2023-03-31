package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return
	}

}
