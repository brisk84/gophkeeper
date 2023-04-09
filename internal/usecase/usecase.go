// Package usecase is a business logic
package usecase

import (
	"context"

	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

//go:generate mockery --name=storage --structname=storageMock --filename=storage_mock.go --inpackage
type storage interface {
	SaveUser(ctx context.Context, user domain.User) (string, error)
	GetByLogin(ctx context.Context, login string) (string, error)
	UpdateByLogin(ctx context.Context, login, token string) error
	GetByToken(ctx context.Context, token string) (domain.User, error)
	SaveData(ctx context.Context, userId int, title string, data []byte, dataType string) (int, error)
	SaveLogin(ctx context.Context, userId int, title string, login, pass string) (int, error)
	ListData(ctx context.Context, userId int) ([]domain.Data, error)
	GetData(ctx context.Context, userId int, dataId int) ([]byte, error)
}

type service struct {
	lg      logger.Logger
	storage storage
}

// New creates a service
func New(lg logger.Logger, stor storage) (*service, error) {
	return &service{
		lg:      lg,
		storage: stor,
	}, nil
}
