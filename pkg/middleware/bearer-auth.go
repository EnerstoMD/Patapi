package middleware

import (
	"lupus/patapi/pkg/auth"
	"lupus/patapi/pkg/model"
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
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)

		c.Next()
	}
}
