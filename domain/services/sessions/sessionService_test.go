package sessions

import (
	"testing"

	"github.com/JoaoLeal92/go_boilerplate/domain/services"
	"github.com/JoaoLeal92/go_boilerplate/infra/config"
	"github.com/JoaoLeal92/go_boilerplate/infra/hash"
	"github.com/JoaoLeal92/go_boilerplate/repositories"
	"github.com/JoaoLeal92/go_boilerplate/tests/fixtures"
	"github.com/stretchr/testify/suite"
)

type authenticationServiceTestSuite struct {
	suite.Suite
	db             *repositories.Connection
	sessionService *Service
}

func (s *authenticationServiceTestSuite) BeforeTest(suitename, testname string) {
	const sqlFilePath = "../../../tests/migrations/001-setup-db.sql"

	baseCfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	cfg := &config.Config{
		Db: config.DBConfig(baseCfg.Tests.Db),
	}

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

	sessionService := NewSessionService(service)

	s.db = testRepo
	s.sessionService = sessionService
}

func TestAuthenticationTestSuite(t *testing.T) {
	suite.Run(t, new(authenticationServiceTestSuite))
}

func (s *authenticationServiceTestSuite) TestAuthenticateUserService() {
	s.db.Users().CreateUser(fixtures.NewFixtureUser())

	tokenString, _, _ := s.sessionService.AuthenticateUserService("joao@teste.com", "12345678")

	s.Suite.Require().NotEmpty(tokenString, "Authentication error, Expected token string")
}

func (s *authenticationServiceTestSuite) TestAuthenticateWrongUser() {
	tokenString, _, err := s.sessionService.AuthenticateUserService("wrong@email.com", "12345678")

	s.Suite.Require().Empty(tokenString, "Authentication error, Expected token string to be empty")
	if s.Suite.Error(err) {
		s.Suite.Require().Equal("Wrong e-mail/password combination", err.Error(), "Expected error on the authentication of non existing user")
	}
}

func (s *authenticationServiceTestSuite) TestAuthenticateWrongPassword() {
	s.db.Users().CreateUser(fixtures.NewFixtureUser())

	tokenString, _, err := s.sessionService.AuthenticateUserService("joao@teste.com", "wrong password")

	s.Suite.Require().Empty(tokenString, "Authentication error, Expected token string to be empty")
	if s.Suite.Error(err) {
		s.Suite.Require().Equal("Wrong e-mail/password combination", err.Error(), "Expected error on the authentication of non existing user")
	}
}
