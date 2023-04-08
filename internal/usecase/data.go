package usecase

import (
	"context"
	"encoding/json"

	"github.com/brisk84/gophkeeper/domain"
)

func (s *service) SaveData(ctx context.Context, userId int, title string, data []byte) (int, error) {
	return s.storage.SaveData(ctx, userId, title, data, "binary")
}

func (s *service) SaveLogin(ctx context.Context, userId int, title string, login, pass string) (int, error) {
	m := make(map[string]string)
	m["login"] = login
	m["pass"] = pass
	data, err := json.Marshal(m)
	if err != nil {
		return 0, nil
	}
	return s.storage.SaveData(ctx, userId, title, data, "login")
}

func (s *service) SaveBankCard(ctx context.Context, userId int, title string, card domain.CardInfo) (int, error) {
	data, err := json.Marshal(card)
	if err != nil {
		return 0, nil
	}
	return s.storage.SaveData(ctx, userId, title, data, "card")
}

func (s *service) SaveText(ctx context.Context, userId int, title string, text string) (int, error) {
	return s.storage.SaveData(ctx, userId, title, []byte(text), "text")
}

func (s *service) ListData(ctx context.Context, userId int) ([]domain.Data, error) {
	return s.storage.ListData(ctx, userId)
}

func (s *service) GetData(ctx context.Context, userId int, dataId int) ([]byte, error) {
	return s.storage.GetData(ctx, userId, dataId)
}
