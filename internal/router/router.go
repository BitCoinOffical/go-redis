package router

import (
	"errors"
	"io"
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/api/handlers"
	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

func Command(Conn net.Conn, h *handlers.Handlers) {
	buffer := make([]byte, 1024)
	defer Conn.Close()
	for {
		n, err := Conn.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			log.Printf(">>> %w: %v\n", ErrorRead, err)
			continue
		}
		text := string(buffer[:n])
		cmd := parser.Decode(text)

		if len(cmd) == 0 {
			Conn.Write([]byte(ErrorEmptyCommand.Error()))
			continue
		}

		if strings.Contains(strings.ToUpper(cmd[0]), "PING") {
			go h.Handler.PingHandler()
			continue
		}

		if strings.Contains(strings.ToUpper(cmd[0]), "ECHO") {
			if len(cmd) > 2 {
				Conn.Write([]byte(ErrorBadAgument.Error()))
				continue
			}
			go h.Handler.EchoHandler(cmd[0])
			continue
		}

		if strings.Contains(strings.ToUpper(cmd[0]), "SET") {
			if len(cmd) < 3 {
				Conn.Write([]byte(ErrorBadAgument.Error()))
				continue
			}
			//set key v px 100 | len = 5
			//set s s | len = 3/ idx = 2
			if len(cmd) > 3 {
				//set we we we | fail!
				//set we we we we| fail!
				//set we we PX 10000 | ok! len = 5 / idx = 4
				if len(cmd) < 5 {
					Conn.Write([]byte(ErrorBadAgument.Error()))
					continue
				}
				go h.Handler.Set(cmd[1], cmd[2], cmd[3], cmd[4])
				continue
			}

			go h.Handler.Set(cmd[1], cmd[2])
			continue
		}

		if strings.Contains(strings.ToUpper(cmd[0]), "GET") {
			if len(cmd) < 2 {
				Conn.Write([]byte(ErrorBadAgument.Error()))
				continue
			}
			go h.Handler.Get(cmd[1])
			continue
		}
	}
}
