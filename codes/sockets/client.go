package sockets

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func TcpClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	fmt.Printf("客户端连接 %s TCP服务\n", addr)
	if err != nil {
		log.Fatal("client err -> ", err)
	}
	defer conn.Close()
	fmt.Println("localaddr -> ", conn.LocalAddr())
	fmt.Println("remoteaddr -> ", conn.RemoteAddr())
	inputReader := bufio.NewReader(os.Stdin)
	for {
		data, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("inputRead error ->", err)
			break
		}
		data = strings.Trim(data, "\r\n")
		if strings.ToUpper(data) == "Q" {
			fmt.Println("inputRead Q to break")
			break
		}
		fmt.Println("tcpClient write -> ", data)
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Println("conn.Write error ->", err)
			break
		}
	}
}

func DefaultTcpClient() {
	TcpClient(ADDR)
}
