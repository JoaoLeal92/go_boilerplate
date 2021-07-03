package services

import (
	"github.com/JoaoLeal92/goals_backend/domain/contract"
	"github.com/JoaoLeal92/goals_backend/infra/config"
	"github.com/JoaoLeal92/goals_backend/infra/hash"
)

// Service Generic service struct
type Service struct {
	Db   contract.RepoManager
	Hash hash.HashProvider
	Cfg  *config.Config
}

// NewService Returns instance of generic service struct
func NewService(db contract.RepoManager, hash hash.HashProvider, cfg *config.Config) *Service {
	return &Service{
		Db:   db,
		Hash: hash,
		Cfg:  cfg,
	}
}
