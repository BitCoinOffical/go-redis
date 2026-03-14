package parser

import "fmt"

func Error(err string) error {
	return fmt.Errorf("-ERR %s\r\n", err)
}
