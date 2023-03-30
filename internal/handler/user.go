package handler

import (
	"errors"
	"net/http"

	"github.com/brisk84/gophkeeper/api/gophserver"
	"github.com/brisk84/gophkeeper/domain"
	"go.uber.org/zap"
)

func (h *Handler) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest[gophserver.RegisterLoginReq](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}
	var user domain.User
	user.FromReq(req)
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
		h.lg.Errorln("Register failed", err)
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}
	resp := gophserver.RegisterLoginResp{
		Token: token,
	}
	sendResponse(w, resp, http.StatusOK)
}

func (h *Handler) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest[gophserver.RegisterLoginReq](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse(w, nil, http.StatusInternalServerError)
		return
	}
	var user domain.User
	user.FromReq(req)
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
	resp := gophserver.RegisterLoginResp{
		Token: token,
	}
	sendResponse(w, resp, http.StatusOK)
}
