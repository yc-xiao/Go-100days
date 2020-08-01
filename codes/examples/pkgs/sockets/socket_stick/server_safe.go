package socket_stick

import (
	"Go-100days/codes/sockets"
	"Go-100days/codes/sockets/proto"
	"fmt"
	"log"
	"net"
	"strings"
)

func processSafe(conn net.Conn) {
	defer conn.Close()
	fmt.Println("localaddr -> ", conn.LocalAddr())
	fmt.Println("remoteaddr -> ", conn.RemoteAddr())

	acceptData := make([]byte, 1024, 1024)
	for {
		// 可以使用缓存接口　reader := bufio.NewReader(conn)
		n, err := conn.Read(acceptData)
		if err != nil {
			fmt.Println("remoteaddr -> ", conn.RemoteAddr(), "-> err ->", err)
			break
		}
		datas := proto.Decode(acceptData[:n])
		for _, dataByte := range datas {
			data := string(dataByte)
			fmt.Println("tcpServer read -> ", data)
			if strings.HasSuffix(data, "end") {
				fmt.Println("remoteaddr -> ", conn.RemoteAddr(), "-> closed")
				break
			}
		}
	}

}

func TcpServerSafe() {
	fmt.Println("启动TCP服务监听端口->", sockets.ADDR)
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
		go processSafe(conn)
	}
}
