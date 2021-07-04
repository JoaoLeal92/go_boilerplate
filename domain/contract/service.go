package contract

import (
	"github.com/JoaoLeal92/go_boilerplate/domain/entities"
)

// UserService user service interface
type UserService interface {
	CreateUserService(userName string, email string, password string) (*entities.User, error)
}

// SessionService user authentication service interface
type SessionService interface {
	AuthenticateUserService(email string, password string) (string, *entities.User, error)
}
