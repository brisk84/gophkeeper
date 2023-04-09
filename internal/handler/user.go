// Package handler contains handlers
package handler

import (
	"context"
	"errors"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
)

// Register creates new user and returns token
func (h *Handler) Register(ctx context.Context, req *api.RegisterLoginReq) (*api.RegisterLoginResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user := domain.User{
		Login:    req.Login,
		Password: req.Password,
	}
	token, err := h.useCase.Register(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrLoginIsBusy) {
			return &api.RegisterLoginResp{
				Success: false,
			}, nil
		}
		// h.lg.Errorln("Register failed", err)
		return nil, err
	}
	return &api.RegisterLoginResp{
		Token:   token,
		Success: true,
	}, nil
}

// Login checks user's login, password and return token, or error
func (h *Handler) Login(ctx context.Context, req *api.RegisterLoginReq) (*api.RegisterLoginResp, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	user := domain.User{
		Login:    req.Login,
		Password: req.Password,
	}
	success, token, err := h.useCase.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	return &api.RegisterLoginResp{
		Token:   token,
		Success: success,
	}, nil
}
