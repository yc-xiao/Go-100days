package main

import (
	"Go-100days/codes/sockets"
	"Go-100days/codes/sockets/socket_stick"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func TcpServer() {
	fmt.Println("启动 TcpServer")
	sockets.DefaultTcpServer()
}

func UdpServer() {
	fmt.Println("启动 UdpServer")
	sockets.DefaultUdpServer()
}

func TcpServerStick() {
	fmt.Println("启动 TcpServerStick")
	socket_stick.TcpServerStick()
}

func TcpServerSafe() {
	fmt.Println("启动 TcpServerSafe")
	socket_stick.TcpServerSafe()
}

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Printf("请选择启动服务(default 1)\n" +
		"1.TCPServer \n" +
		"2.UDPServer \n" +
		"3.TcpServerStick \n" +
		"4.TCPServerSafe \n" +
		"5.Q(退出) \n")
	str, err := input.ReadString('\n')
	if err != nil {
		log.Fatal("socket_server err -> ", err)
	}
	str = strings.Trim(str, "\r\n")
	switch str {
	case "1":
		TcpServer()
	case "2":
		UdpServer()
	case "3":
		TcpServerStick()
	case "4":
		TcpServerSafe()
	case "Q", "q", "5":
		break
	default:
		TcpServer()
	}
	fmt.Println("socket_server close")
}
