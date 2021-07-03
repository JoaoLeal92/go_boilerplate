package middlewares

import (
	"net/http"
	"strings"

	"github.com/JoaoLeal92/goals_backend/infra/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type header struct {
	Authorization string
}

// EnsureAuthenticated ensures the user is authenticated and can access a given route
func EnsureAuthenticated(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := header{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		tokenParts := strings.Split(h.Authorization, " ")

		if len(tokenParts) == 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
		}

		tokenString := tokenParts[1]

		// Parse cookie to extract content
		claims := claim{}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.Global.SecretKey), nil
		})

		// Check token validity
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
