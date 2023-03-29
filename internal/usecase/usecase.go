package usecase

import (
	"context"

	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

type storage interface {
	UserRegister(user domain.User) (domain.Bearer, error)
	UserLogin(user domain.User) (domain.Bearer, error)

	SaveUser(ctx context.Context, user domain.User) (string, error)
	GetByLogin(ctx context.Context, login string) (string, error)
	UpdateByLogin(ctx context.Context, login, token string) error
	GetByToken(ctx context.Context, token string) (domain.User, error)
}

type service struct {
	lg      logger.Logger
	storage storage
}

func New(lg logger.Logger, cfg config.Config, stor storage) (*service, error) {
	return &service{
		lg:      lg,
		storage: stor,
	}, nil
}

func (s *service) UserRegister(user domain.User) (domain.Bearer, error) {
	return s.storage.UserRegister(user)
}

func (s *service) UserLogin(user domain.User) (domain.Bearer, error) {
	return s.storage.UserLogin(user)
}
