package contract

import (
	"github.com/JoaoLeal92/go_boilerplate/domain/entities"
)

// UserService user service interface
type UserService interface {
	CreateUserService(user *entities.User) (*entities.User, error)
}

// SessionService user authentication service interface
type SessionService interface {
	AuthenticateUserService(user *entities.User) (string, *entities.User, error)
}
