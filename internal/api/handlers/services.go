package handlers

import "net"

type Handlers struct {
	Ping *PingHandler
}

func NewHandlers(Conn net.Conn) *Handlers {
	ping := NewPingHandler(Conn)
	return &Handlers{Ping: ping}
}
