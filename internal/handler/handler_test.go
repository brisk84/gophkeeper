package handler

import (
	"context"
	"log"
	"testing"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	uc := new(useCaseMock)
	ctx := context.Background()
	mockToken := "b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4"
	mockReq := api.RegisterLoginReq{Login: "user01", Password: "pass"}
	mockUser := domain.User{
		Login:    "user01",
		Password: "pass",
	}

	uc.On("Register", ctx, mockUser).Return(mockToken, nil).Once()
	uc.On("Register", ctx, mockUser).Return("", domain.ErrLoginIsBusy).Once()

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}
	h := New(lg, uc)

	t.Run("success", func(t *testing.T) {
		ret, err := h.Register(ctx, &mockReq)
		assert.NoError(t, err)
		assert.Equal(t, true, ret.Success)
		assert.Equal(t, mockToken, ret.Token)
	})

	t.Run("busy", func(t *testing.T) {
		ret, err := h.Register(ctx, &mockReq)
		assert.NoError(t, err)
		assert.Equal(t, false, ret.Success)
		assert.Equal(t, "", ret.Token)
	})
}

func TestLogin(t *testing.T) {
	uc := new(useCaseMock)
	ctx := context.Background()
	mockToken := "b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4"
	mockReq1 := api.RegisterLoginReq{Login: "user01", Password: "pass"}
	mockReq2 := api.RegisterLoginReq{Login: "user02", Password: "pass"}

	mockUser1 := domain.User{
		Login:    "user01",
		Password: "pass",
	}
	mockUser2 := domain.User{
		Login:    "user02",
		Password: "pass",
	}

	uc.On("Login", ctx, mockUser1).Return(true, mockToken, nil).Once()
	uc.On("Login", ctx, mockUser2).Return(false, "", nil).Once()

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}
	h := New(lg, uc)

	t.Run("success", func(t *testing.T) {
		ret, err := h.Login(ctx, &mockReq1)
		assert.NoError(t, err)
		assert.Equal(t, true, ret.Success)
		assert.Equal(t, mockToken, ret.Token)
	})

	t.Run("denied", func(t *testing.T) {
		ret, err := h.Login(ctx, &mockReq2)
		assert.NoError(t, err)
		assert.Equal(t, false, ret.Success)
		assert.Equal(t, "", ret.Token)
	})
}
