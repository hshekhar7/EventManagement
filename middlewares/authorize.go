package middlewares

import (
	"net/http"

	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
)

func Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userclaims, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		claims, ok := userclaims.(*models.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
			c.Abort()
			return
		}
		if claims.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
