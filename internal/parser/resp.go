package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func BulkString(text string) string {
	line := strings.Fields(text)
	for i := 0; i <= len(line)-1; i++ {
		line[i] = fmt.Sprintf("$%d\r\n%s\r\n", len(line[i]), line[i])
	}
	return strings.Join(line, "")
}

func Array(text string) string {
	line := strings.Fields(text)
	for i := 0; i <= len(line)-1; i++ {
		line[i] = fmt.Sprintf("$%d\r\n%s\r\n", len(line[i]), line[i])
	}
	line = append([]string{fmt.Sprintf("*%d\r\n", len(line))}, line...)
	return strings.Join(line, "")
}

func SimpleString(text string) string {
	return fmt.Sprintf("+%s\r\n", text)
}

func Integer(integer int) string {
	return fmt.Sprintf(":%d\r\n", integer)
}

func Decode(text string) []string {
	res := []string{}
	r := bufio.NewReader(strings.NewReader(text))
	lenArr := 0
	str, _ := r.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	idx := strings.Index(str, "*")
	if idx == 0 {
		lenArr, _ = strconv.Atoi(str[idx+1:])
	}
	lentext := 0

	for range lenArr {
		str, _ = r.ReadString('\n')
		str = strings.TrimSuffix(str, "\r\n")
		idx = strings.Index(str, "$")
		if idx == 0 {
			lentext, _ = strconv.Atoi(str[idx+1:])
		}
		buf := make([]byte, lentext)
		io.ReadFull(r, buf)
		r.ReadString('\n')
		res = append(res, string(buf))

	}
	return res
}
