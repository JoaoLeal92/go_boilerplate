package user

import (
	"errors"
	"fmt"
	"testing"

	"github.com/JoaoLeal92/goals_backend/domain/entities"
	"github.com/JoaoLeal92/goals_backend/domain/services"
	"github.com/JoaoLeal92/goals_backend/infra/config"
	"github.com/JoaoLeal92/goals_backend/infra/hash"
	"github.com/JoaoLeal92/goals_backend/repositories"
	"github.com/JoaoLeal92/goals_backend/tests/fixtures"
	mockRepo "github.com/JoaoLeal92/goals_backend/tests/mocks/repositories"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userServiceTestSuite struct {
	suite.Suite
	db          *repositories.Connection
	userService *Service
}

func (s *userServiceTestSuite) BeforeTest(suitename, testname string) {
	baseCfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	cfg := &config.Config{
		Db: config.DBConfig(baseCfg.Tests.Db),
	}

	const sqlFilePath = "../../../tests/migrations/001-setup-db.sql"

	fakeHashProvider := hash.NewFakeProvider()

	testRepo, err := repositories.ConnectDataBase(cfg.Db)
	if err != nil {
		panic(err)
	}

	err = fixtures.PrepareDatabase(testRepo, sqlFilePath)
	if err != nil {
		panic(err)
	}

	service := services.NewService(testRepo, fakeHashProvider, cfg)

	userService := NewUserService(service)

	s.db = testRepo
	s.userService = userService
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(userServiceTestSuite))
}

func (s *userServiceTestSuite) TestCreateUserService() {
	var expectedResult entities.User
	expectedResult.Name = "João"
	expectedResult.Email = "joao@teste.com"
	expectedResult.Password = "12345678"

	user, err := s.userService.CreateUserService("João", "joao@teste.com", "12345678")

	s.Suite.Equal(expectedResult.Name, user.Name)
	s.Suite.Equal(expectedResult.Email, user.Email)
	s.Suite.Equal(expectedResult.Password, user.Password)
	s.Suite.Require().NotEmpty(user.ID)
	s.Suite.NoError(err)
}

func (s *userServiceTestSuite) TestCreateUserWithInvalidEmail() {
	expectedResult := entities.User{}

	s.db.Users().CreateUser(fixtures.NewFixtureUser())

	user, err := s.userService.CreateUserService("João 2", "joao@teste.com", "12345678")

	s.Suite.Equal(&expectedResult, user)
	s.Suite.Error(err)
	if s.Suite.Assert().Error(err) {
		s.Suite.Equal("Email already in use", err.Error(), "Expected error on the creation of user with e-mail already in use")
	}
}

func (s *userServiceTestSuite) TestCreateUserWithInvalidPassword() {
	expectedResult := entities.User{}

	user, err := s.userService.CreateUserService("João", "joao@teste.com", "")

	s.Suite.Equal(user, &expectedResult, fmt.Sprintf("Expected created user to be %+v", expectedResult))
	s.Suite.Error(err)
	if s.Suite.Assert().Error(err) {
		s.Suite.Equal("Password required", err.Error(), "Expected error on the creation of user with invalid password")
	}
}

func (s *userServiceTestSuite) TestCreateUserError() {
	mockRepository := mockRepo.RepoManager{}
	mockUserRepo := mockRepo.UserRepository{}
	s.userService.svc.Db = &mockRepository

	mockRepository.On("Users").Return(&mockUserRepo)
	mockUserRepo.On("CreateUser", mock.Anything).Return(errors.New("User creation error"))
	mockUserRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(&entities.User{}, errors.New("record not found"))

	expectedResult := entities.User{}

	user, err := s.userService.CreateUserService("João", "joao@teste.com", "12345678")

	s.Suite.Equal(user, &expectedResult, fmt.Sprintf("Expected created user to be %+v", expectedResult))
	s.Suite.Error(err)
	if s.Suite.Assert().Error(err) {
		s.Suite.Equal("User creation error", err.Error(), "Expected error on the insertion of user in repository")
	}
}
