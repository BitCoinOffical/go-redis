package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	conn, err := net.Dial("tcp", "192.168.0.11:6379")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	for i := range 5 {
		wg.Go(func() {
			fmt.Fprintln(conn, "PING ECHO  PING PING ergoijkgdf gpreg pdfkjg dikfgpdrkgedk ECHO echo echo echo 4393094i3409#(*$%@&#%*@#%&@)(*&%!_!$_!$+@#%*&$@*&%!////////)")
			fmt.Println("send", i)

		})
	}
	wg.Wait()

	for range 5 {
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		msg := strings.TrimRight(string(buf[:n]), "\r\n")
		fmt.Println(msg)
	}

}
