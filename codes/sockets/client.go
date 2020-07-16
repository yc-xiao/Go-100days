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

func UdpClient(addr string) {
	IP, Port, err := strToip(addr)
	if err != nil {
		log.Fatal("strToip error -> ", err)
	}
	client, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: IP, Port: Port})
	if err != nil {
		log.Fatal("UdpClient error -> ", err)
	}
	fmt.Println("UpdClient -> ", IP, Port)
	defer client.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for {
		data, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("inputReader error -> ", err)
			break
		}
		data = strings.Trim(data, "\r\n")
		if strings.ToUpper(data) == "Q" {
			fmt.Println("client Q break ")
			break
		}

		_, err = client.Write([]byte(data))
		if err != nil {
			fmt.Println("client error -> ", err)
			break
		}
		fmt.Println("client write -> ", data)
	}
}

func DefaultUdpClient() {
	UdpClient(ADDR)
}
