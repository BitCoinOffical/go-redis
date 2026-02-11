package handlers

import (
	"bufio"
	"log"
	"net"
)

type PingHandler struct {
	Conn net.Conn
}

func NewPingHandler(Conn net.Conn) *PingHandler {

	return &PingHandler{Conn: Conn}
}

func (h *PingHandler) PingHandler() {
	defer h.Conn.Close()
	for {
		go func() {

			scanner := bufio.NewScanner(h.Conn)

			for scanner.Scan() {
				text := scanner.Text()
				log.Println(text)
				if text == "PING" {
					h.Conn.Write([]byte("+PONG\r\n"))
				}
			}

		}()
	}

}
