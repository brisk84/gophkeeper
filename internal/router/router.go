package router

import (
	"net/http"

	"github.com/brisk84/gophkeeper/api/gophserver"
	"github.com/brisk84/gophkeeper/internal/handler"
)

func New(h *handler.Handler) http.Handler {
	return gophserver.Handler(h)
}
