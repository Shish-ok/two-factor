package auth

import (
	"time"
	"two-factor-auth/internal/config"
)

type Storage interface {
}

func NewService(cfg config.ServiceConfiguration, storage Storage) *Service {
	return &Service{
		codeTTL: cfg.AuthConfig.TTL,
		codeLen: cfg.AuthConfig.CodeLen,
		storage: storage,
	}
}

type Service struct {
	codeTTL time.Duration
	codeLen uint
	storage Storage
}
