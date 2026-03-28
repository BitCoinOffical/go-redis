package handlers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/dto"
	"github.com/codecrafters-io/redis-starter-go/internal/interfaces/http/services"
	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

type Handler struct {
	Conn    net.Conn
	service *services.Service
}

func NewHandler(Conn net.Conn, setvice *services.Service) *Handler {

	return &Handler{Conn: Conn, service: setvice}
}

func (h *Handler) PingHandler() {
	h.Conn.Write([]byte(parser.SimpleString("PONG")))
}

func (h *Handler) EchoHandler(responce string) {
	h.Conn.Write([]byte(responce))
}

func (h *Handler) Set(setDTO *dto.SetDTO) {
	status, err := h.service.SetService(setDTO)
	if err != nil {
		h.Conn.Write([]byte(err.Error()))
		return
	}
	h.Conn.Write([]byte(status))
}

func (h *Handler) Get(key string) {
	res, err := h.service.GetService(key)
	if err != nil {
		h.Conn.Write([]byte(err.Error()))
		return
	}
	value := fmt.Sprintf("%s", res)
	h.Conn.Write([]byte(value))
}
