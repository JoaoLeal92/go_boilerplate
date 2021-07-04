package user

import (
	"errors"

	"github.com/JoaoLeal92/go_boilerplate/domain/entities"
	"github.com/JoaoLeal92/go_boilerplate/domain/services"
)

// Service user service struct
type Service struct {
	svc *services.Service
}

// NewUserService creates new user service
func NewUserService(svc *services.Service) *Service {
	return &Service{
		svc: svc,
	}
}

// CreateUserService creates new user, if it doesn't already exists
func (s *Service) CreateUserService(userName string, email string, password string) (*entities.User, error) {
	userRepo := s.svc.Db.Users()

	user := userRepo.FindUserByEmail(email)

	if user.Email != "" {
		return &entities.User{}, errors.New("Email already in use")
	}

	hashedPassword, err := s.svc.Hash.GenerateHash(password)
	if err != nil {
		return &entities.User{}, err
	}

	user.Password = string(hashedPassword)
	user.Name = string(userName)
	user.Email = string(email)

	err = userRepo.CreateUser(user)
	if err != nil {
		return &entities.User{}, err
	}

	return user, nil
}
