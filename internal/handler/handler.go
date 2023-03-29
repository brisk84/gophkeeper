package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

type Handler struct {
	lg      logger.Logger
	useCase useCase
}

type useCase interface {
	UserRegister(user domain.User) (domain.Bearer, error)
	UserLogin(user domain.User) (domain.Bearer, error)
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

func (h *Handler) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	user, err := parseRequest[domain.User](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}

	bearer, err := h.useCase.UserRegister(user)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusBadRequest)
		return
	}
	sendResponse(w, bearer, http.StatusOK)
}

func (h *Handler) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	user, err := parseRequest[domain.User](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}

	bearer, err := h.useCase.UserLogin(user)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusBadRequest)
		return
	}
	sendResponse(w, bearer, http.StatusOK)
}
