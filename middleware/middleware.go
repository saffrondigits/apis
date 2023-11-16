package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/apis/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return handler
}

func handler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	username, err := auth.AuthenticateToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	c.Set("username", username)
	c.Next()
}
