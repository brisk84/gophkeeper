// Storage package uses postgres database
package storage

import (
	"context"
	"encoding/json"

	"github.com/brisk84/gophkeeper/domain"
)

// SaveData stores a data to database
func (s *storage) SaveData(ctx context.Context, userId int, title string, data []byte, dataType string) (int, error) {
	var ret int
	sql1 := `insert into data (user_id, title, data, type) values ($1, $2, $3, $4) returning id`
	row := s.db.QueryRowContext(ctx, sql1, userId, title, data, dataType)
	err := row.Scan(&ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

// SaveLogin stores a login and password to database
func (s *storage) SaveLogin(ctx context.Context, userId int, title string, login, pass string) (int, error) {
	var ret int
	m := make(map[string]string)
	m["login"] = login
	m["pass"] = pass
	data, err := json.Marshal(m)
	if err != nil {
		return 0, nil
	}
	sql1 := `insert into data (user_id, title, login, pass, data, type) values ($1, $2, $3, $4, $5, 'login') returning id`
	row := s.db.QueryRowContext(ctx, sql1, userId, title, login, pass, data)
	err = row.Scan(&ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

// ListData returns list of records by user id
func (s *storage) ListData(ctx context.Context, userId int) ([]domain.Data, error) {
	var ret []domain.Data
	sql1 := `select id, title, type from data where user_id = $1`
	rows, err := s.db.QueryContext(ctx, sql1, userId)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	for rows.Next() {
		var data domain.Data
		err = rows.Scan(&data.Id, &data.Title, &data.Type)
		if err != nil {
			return nil, err
		}
		ret = append(ret, data)
	}
	return ret, nil
}

// GetData returns a data by data id
func (s *storage) GetData(ctx context.Context, userId int, dataId int) ([]byte, error) {
	var ret []byte
	sql1 := `select data from data where id = $1`
	row := s.db.QueryRowContext(ctx, sql1, dataId)
	err := row.Scan(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
