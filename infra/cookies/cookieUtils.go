package cookies

import (
	"github.com/dgrijalva/jwt-go"
)

type claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GetUserIDFromCookie returns user ID stored in cookie
func GetUserIDFromCookie(cookie string, secretKey string) (string, error) {
	claims := &claim{}

	_, errClaim := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if errClaim != nil {
		return "", errClaim
	}

	return claims.UserID, nil
}
