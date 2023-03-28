package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	token := c.GetHeader("Bearer")
}
