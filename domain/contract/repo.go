package contract

import "github.com/JoaoLeal92/goals_backend/domain/entities"

// RepoManager repository manager interface
type RepoManager interface {
	Users() UserRepository
}

// UserRepository interface for users repo
type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) *entities.User
	FindUserByID(ID string) *entities.User
}
