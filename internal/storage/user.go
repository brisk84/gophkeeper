package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/brisk84/gophkeeper/domain"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

func (s *storage) SaveUser(ctx context.Context, user domain.User) (string, error) {
	sql1 := `insert into users (login, password, token) values ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, sql1, user.Login, user.Hash, user.Token)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pgerrcode.UniqueViolation {
				return "", domain.ErrLoginIsBusy
			}
		}
		return "", err
	}
	return user.Token, nil
}

func (s *storage) GetByLogin(ctx context.Context, login string) (string, error) {
	var hash string
	sql1 := `select password from users where login = $1`
	row := s.db.QueryRowContext(ctx, sql1, login)
	err := row.Scan(&hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return hash, nil
}

func (s *storage) UpdateByLogin(ctx context.Context, login, token string) error {
	sql1 := `update users set token = $1 where login = $2`
	_, err := s.db.ExecContext(ctx, sql1, token, login)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) GetByToken(ctx context.Context, token string) (domain.User, error) {
	var user domain.User
	sql1 := `select id, login from users where token = $1`
	row := s.db.QueryRowContext(ctx, sql1, token)
	err := row.Scan(&user.ID, &user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUrerNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}
