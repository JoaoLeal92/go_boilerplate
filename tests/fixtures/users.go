package fixtures

import "github.com/JoaoLeal92/goals_backend/domain/entities"

// NewFixtureUser returns a fixture user entity
func NewFixtureUser() *entities.User {
	var user entities.User
	user.Name = "João"
	user.Email = "joao@teste.com"
	user.Password = "12345678"

	return &user
}
