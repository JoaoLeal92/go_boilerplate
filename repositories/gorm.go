package repositories

import (
	"fmt"

	"github.com/JoaoLeal92/goals_backend/domain/contract"
	"github.com/JoaoLeal92/goals_backend/infra/config"
	"github.com/JoaoLeal92/goals_backend/repositories/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection database connection
type Connection struct {
	Db *gorm.DB
}

// Instance returns an instance of the database
func Instance(cfg config.DBConfig) (*Connection, error) {
	var gormConfig = &gorm.Config{}
	if cfg.SilentLog {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		return &Connection{}, err
	}

	return &Connection{Db: db}, nil
}

// Users returns connection to the users repository
func (c *Connection) Users() contract.UserRepository {
	return users.NewRepository(c.Db)
}
