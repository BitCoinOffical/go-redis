package handlers

import "net"

type Handlers struct {
	Handler *Handler
}

func NewHandlers(Conn net.Conn) *Handlers {
	handler := NewHandler(Conn)
	return &Handlers{Handler: handler}
}
