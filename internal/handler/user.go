package handler

import (
	"errors"
	"net/http"

	"github.com/brisk84/gophkeeper/domain"
	"go.uber.org/zap"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	user, err := parseRequest[domain.User](r)
	if err != nil {
		h.lg.Errorln("parse request", zap.Error(err))
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}
	if !user.IsValid() {
		sendResponse(w, nil, http.StatusBadRequest)
		return
	}

	token, err := h.useCase.Register(r.Context(), user)
	if err != nil {
		if errors.Is(err, domain.ErrLoginIsBusy) {
			sendResponse(w, nil, http.StatusConflict)
			return
		}
		h.lg.Errorln("Register failed", zap.Error(err))
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Authorization", "Bearer "+token)

	sendResponse(w, nil, http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := parseRequest[domain.User](r)
	if err != nil {
		h.lg.Errorln("parse request", zap.Error(err))
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}
	if !user.IsValid() {
		sendResponse(w, nil, http.StatusBadRequest)
		return
	}

	success, token, err := h.useCase.Login(r.Context(), user)
	if err != nil {
		h.lg.Errorln("Login failed", zap.Error(err))
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}

	if !success {
		sendResponse(w, nil, http.StatusUnauthorized)
		return
	}

	w.Header().Add("Authorization", "Bearer "+token)
	sendResponse(w, nil, http.StatusOK)
}
