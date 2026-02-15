package router

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/api/handlers"
)

func Command(Conn net.Conn, h *handlers.Handlers) {
	buffer := make([]byte, 1024)
	defer Conn.Close()
	for {
		_, err := Conn.Read(buffer)
		if err != nil {
			fmt.Println("Клиент отключился:", err)
			return
		}
		line := strings.Split(string(buffer), "\n")
		log.Println("разбили", line)
		for _, v := range line {
			log.Println("получили", v)
			if strings.HasPrefix(v, "PING") {
				go h.Handler.PingHandler()
				continue
			}
			if strings.HasPrefix(v, "ECHO") {
				log.Println("вызвали echo")
				_, a, _ := strings.Cut(v, "ECHO")
				go h.Handler.EchoHandler(a)
				continue
			}
		}
	}

}
