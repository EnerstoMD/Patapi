package middleware

import (
	"log"
	"lupus/patapi/pkg/auth"
	"lupus/patapi/pkg/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BearerAuth(a auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := model.JwtWrapper{
			SecretKey: "secret",
			Issuer:    "lupus",
		}

		claims, err := a.ValidateToken(c, clientToken, jwtWrapper)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("userId", claims.Subject)
		c.Next()
	}
}
