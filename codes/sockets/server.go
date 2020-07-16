package sockets

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("localaddr -> ", conn.LocalAddr())
	fmt.Println("remoteaddr -> ", conn.RemoteAddr())

	var data string
	acceptData := make([]byte, 1024, 1024)
	for {
		// 可以使用缓存接口　reader := bufio.NewReader(conn)
		n, err := conn.Read(acceptData)
		if err != nil {
			fmt.Println("remoteaddr -> ", conn.RemoteAddr(), "-> err ->", err)
			break
		}
		data = string(acceptData[:n])
		fmt.Println("tcpServer read -> ", data)
		if strings.HasSuffix(data, "end") {
			fmt.Println("remoteaddr -> ", conn.RemoteAddr(), "-> closed")
			break
		}
	}

}

func TcpServer(address string) {
	fmt.Println("启动TCP服务监听端口->", address)
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error -> ", err)
			continue
		}
		go process(conn)
	}
}

func DefaultTcpServer() {
	TcpServer(ADDR)
}
