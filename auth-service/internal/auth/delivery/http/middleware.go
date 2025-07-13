package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/StevieAdrian/Fyn-API/auth-service/pkg/token"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		// Remove Bearer
		authHeader = strings.TrimPrefix(authHeader, "Bearer")

		claims, err := token.ValidateToken(authHeader)

		if err != nil {
			log.Printf("Token validation error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}

	// Bearer <token>
}
