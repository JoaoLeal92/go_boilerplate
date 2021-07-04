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
func (s *Service) CreateUserService(user *entities.User) (*entities.User, error) {
	userRepo := s.svc.Db.Users()

	userData := userRepo.FindUserByEmail(user.Email)

	if userData.Email != "" {
		return &entities.User{}, errors.New("Email already in use")
	}

	hashedPassword, err := s.svc.Hash.GenerateHash(user.Password)
	if err != nil {
		return &entities.User{}, err
	}

	userData.Password = string(hashedPassword)
	userData.Name = string(user.Name)
	userData.Email = string(user.Email)

	err = userRepo.CreateUser(userData)
	if err != nil {
		return &entities.User{}, err
	}

	return userData, nil
}
