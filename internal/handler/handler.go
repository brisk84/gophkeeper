package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

type Handler struct {
	api.UnimplementedGophKeeperServer
	lg      logger.Logger
	useCase useCase
}

//go:generate mockery --name=useCase --structname=useCaseMock --filename=usecase_mock.go --inpackage
type useCase interface {
	Register(ctx context.Context, user domain.User) (string, error)
	Login(ctx context.Context, user domain.User) (bool, string, error)
	Auth(ctx context.Context, token string) (domain.User, error)
	SaveData(ctx context.Context, userId int, title string, data []byte) (int, error)
	SaveLogin(ctx context.Context, userId int, title string, login, pass string) (int, error)
	SaveText(ctx context.Context, userId int, title string, text string) (int, error)
	SaveBankCard(ctx context.Context, userId int, title string, card domain.CardInfo) (int, error)
	ListData(ctx context.Context, userId int) ([]domain.Data, error)
	GetData(ctx context.Context, userId int, dataId int) ([]byte, error)
}

func New(lg logger.Logger, useCase useCase) *Handler {
	return &Handler{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *Handler) PostHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
