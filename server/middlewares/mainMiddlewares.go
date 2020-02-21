package middlewares

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		fmt.Println("AuthHeader: " + authHeader)

		if authHeader == "" {
			c.Header("WWW-Authenticate", "JWT realm=")
			c.Abort()
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Not Have Authorization",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.Header("WWW-Authenticate", "JWT realm=")
			c.Abort()
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Invalid Header",
			})
			return
		}

		token := parts[1]
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return "secret token that long enoght", nil
		})

		if err == nil {
			c.Header("WWW-Authenticate", "JWT realm=")
			c.Abort()
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Invalid Token",
			})
			return
		}

		claims := jwt.MapClaims{}
		for key, value := range parsedToken.Claims.(jwt.MapClaims) {
			claims[key] = value
		}

		fmt.Println("Claims: " + fmt.Sprintf("%v", claims))

		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
