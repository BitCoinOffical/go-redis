package handlers

import (
	"fmt"
	"net"
)

type Handler struct {
	Conn net.Conn
}

func NewHandler(Conn net.Conn) *Handler {

	return &Handler{Conn: Conn}
}

func (h *Handler) PingHandler() {
	h.Conn.Write([]byte("+PONG\r\n"))
}

func (h *Handler) EchoHandler(a string) {
	text := fmt.Sprintf(string([]byte("+%v\r\n")), a)
	h.Conn.Write([]byte(text))
}
