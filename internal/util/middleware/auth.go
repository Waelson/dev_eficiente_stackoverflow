package middleware

import (
	"github.com/Waelson/internal/util/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			c.JSON(401, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}

		jwt := strings.Split(header, " ")[1]

		claims, err := token.Validate(jwt)
		if err != nil {
			c.JSON(401, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}

		c.Request.Header.Add("user", claims.Username)
		c.Next()
	}
}
