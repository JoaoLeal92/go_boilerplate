package repositories

import (
	"github.com/JoaoLeal92/goals_backend/infra/config"
)

// ConnectDataBase Connects to database
func ConnectDataBase(cfg config.DBConfig) (*Connection, error) {
	return Instance(cfg)
}
