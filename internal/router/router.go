package router

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/api/handlers"
	"github.com/codecrafters-io/redis-starter-go/internal/dto"
	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

func Command(Conn net.Conn, h *handlers.Handlers) {
	buffer := make([]byte, 1024)
	defer Conn.Close()
	for {
		n, err := Conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("client disconnect")
			} else {
				log.Printf("%s: %s\n", ErrorRead.Error(), err.Error())
			}
			break
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

			typettl := cmd[len(cmd)-2]
			if strings.ToUpper(typettl) == "PX" || strings.ToUpper(typettl) == "EX" {
				valuettl, err := strconv.Atoi(cmd[len(cmd)-1])
				if err != nil {
					Conn.Write([]byte(ErrorBadAgument.Error()))
					continue
				}
				dto := dto.SetDTO{
					Key:      cmd[1],
					Typettl:  typettl,
					Valuettl: valuettl,
					Values:   cmd[2:(len(cmd) - 2)],
				}

				go h.Handler.Set(&dto)
				continue
			}

			dto := dto.SetDTO{
				Key:    cmd[1],
				Values: cmd[2:],
			}

			go h.Handler.Set(&dto)
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

		if strings.Contains(strings.ToUpper(cmd[0]), "RPUSH") {
			//rpush key v
			//rpush key v px 20
			////rpush key value1 value2 value3 px 20
			if len(cmd) < 3 {
				Conn.Write([]byte(ErrorBadAgument.Error()))
				continue
			}
			typettl := cmd[len(cmd)-2]
			if strings.ToUpper(typettl) == "PX" || strings.ToUpper(typettl) == "EX" {
				valuettl, err := strconv.Atoi(cmd[len(cmd)-1])
				if err != nil {
					Conn.Write([]byte(ErrorBadAgument.Error()))
					continue
				}
				dto := dto.SetDTO{
					Key:      cmd[1],
					Typettl:  typettl,
					Valuettl: valuettl,
					Values:   cmd[2:(len(cmd) - 2)],
				}

				go h.Handler.Set(&dto)
				continue
			}
			dto := dto.SetDTO{
				Key:    cmd[1],
				Values: cmd[2:],
			}

			go h.Handler.Set(&dto)
			continue
		}
	}
}
