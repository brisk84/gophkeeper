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

type useCase interface {
	Register(ctx context.Context, user domain.User) (string, error)
	Login(ctx context.Context, user domain.User) (bool, string, error)
	Auth(ctx context.Context, token string) (domain.User, error)
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
