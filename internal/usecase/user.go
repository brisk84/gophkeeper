package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/brisk84/gophkeeper/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, user domain.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	user.Hash = string(hash)
	user.Token, err = s.genToken()
	if err != nil {
		return "", err
	}
	return s.storage.SaveUser(ctx, user)
}

func (s *service) genToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)
	return token, nil
}

func (s *service) Login(ctx context.Context, user domain.User) (bool, string, error) {
	var err error
	user.Hash, err = s.storage.GetByLogin(ctx, user.Login)
	if err != nil {
		return false, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(user.Password))
	if err != nil {
		return false, "", nil
	}

	user.Token, err = s.genToken()
	if err != nil {
		return false, "", err
	}

	err = s.storage.UpdateByLogin(ctx, user.Login, user.Token)
	if err != nil {
		return false, "", err
	}

	return true, user.Token, nil
}

func (s *service) Auth(ctx context.Context, token string) (domain.User, error) {
	return s.storage.GetByToken(ctx, token)
}
