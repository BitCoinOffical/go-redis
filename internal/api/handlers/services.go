package handlers

import (
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/adapters/secondary/storage"
	"github.com/codecrafters-io/redis-starter-go/internal/interfaces/http/repository"
	"github.com/codecrafters-io/redis-starter-go/internal/interfaces/http/services"
)

type Handlers struct {
	Handler *Handler
}

func NewHandlers(Conn net.Conn) *Handlers {
	srg := storage.NewStorage()
	repo := repository.NewRepository(srg)
	srv := services.NewService(repo)
	handler := NewHandler(Conn, srv)
	return &Handlers{Handler: handler}
}
