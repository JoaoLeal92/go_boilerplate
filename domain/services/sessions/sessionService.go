package sessions

import (
	"errors"
	"time"

	"github.com/JoaoLeal92/go_boilerplate/domain/entities"
	"github.com/JoaoLeal92/go_boilerplate/domain/services"
	"github.com/dgrijalva/jwt-go"
)

// Service session service struct
type Service struct {
	svc *services.Service
}

type claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewSessionService creates new user service
func NewSessionService(svc *services.Service) *Service {
	return &Service{
		svc: svc,
	}
}

// AuthenticateUserService authenticates registered user
func (s *Service) AuthenticateUserService(user *entities.User) (string, *entities.User, error) {
	userRepo := s.svc.Db.Users()

	userData := userRepo.FindUserByEmail(user.Email)

	if userData.Email == "" {
		return "", &entities.User{}, errors.New("Wrong e-mail/password combination")
	}

	// Check if passwords match
	err := s.svc.Hash.CompareHashAndPassword(userData.Password, user.Password)
	if err != nil {
		return "", &entities.User{}, errors.New("Wrong e-mail/password combination")
	}

	// Create JWT token for user
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &claim{
		UserID: userData.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	secretKey := s.svc.Cfg.Global.SecretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", &entities.User{}, errors.New("Internal server error")
	}

	return tokenString, userData, nil
}
