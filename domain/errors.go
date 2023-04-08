package domain

import "errors"

var (
	ErrLoginIsBusy  = errors.New("login is busy")
	ErrUrerNotFound = errors.New("user not found")
)
