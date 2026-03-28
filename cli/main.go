package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/parser"
)

func main() {
	var wg sync.WaitGroup
	for id := range 5 {
		wg.Go(func() {
			conn, err := net.Dial("tcp", "192.168.0.11:6379")
			if err != nil {
				log.Printf("Client %d: connect error: %v\n", id, err)
				return
			}
			defer conn.Close()
			//rpush key value ttl
			ping := "RpuSh key v1 px 2000000"
			_, err = conn.Write([]byte(parser.Array(ping)))
			if err != nil {
				log.Printf("Client %d: write error: %v\n", id, err)
				return
			}

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				log.Printf("Client %d: read error: %v\n", id, err)
				return
			}

			fmt.Printf("Client %d received: %q\n", id, buffer[:n])

			//rpush key value ttl
			ping = "RpuSh key v2 px 1"
			_, err = conn.Write([]byte(parser.Array(ping)))
			if err != nil {
				log.Printf("Client %d: write error: %v\n", id, err)
				return
			}

			buffer = make([]byte, 1024)
			n, err = conn.Read(buffer)
			if err != nil {
				log.Printf("Client %d: read error: %v\n", id, err)
				return
			}

			fmt.Printf("Client %d received: %q\n", id, buffer[:n])

			time.Sleep(time.Second)
			//v1, nil
			ping = "get key"
			_, err = conn.Write([]byte(parser.Array(ping)))
			if err != nil {
				log.Printf("Client %d: write error: %v\n", id, err)
				return
			}

			buffer = make([]byte, 1024)
			n, err = conn.Read(buffer)
			if err != nil {
				log.Printf("Client %d: read error: %v\n", id, err)
				return
			}

			fmt.Printf("Client %d received: %q\n", id, buffer[:n])

		})
	}
	wg.Wait()
}
