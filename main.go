package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()
	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	strSize ,_ := strconv.ParseInt(string(size), 10, 64)
	reader.ReadByte()~
	reader.ReadByte()
	fmt.Println("helloworld")

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)

		if err != io.EOF {
			break
		}
		fmt.Println(err)
		os.Exit(1)
	}

	conn.Write([]byte("+OK\r\n"))

}
