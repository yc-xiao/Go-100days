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

func TcpClient() {
	fmt.Println("启动 TcpClient")
	sockets.DefaultTcpClient()
}

func UdpClient() {
	fmt.Println("启动 UdpClient")
	sockets.DefaultUdpClient()
}

func TcpClientStick() {
	fmt.Println("启动 TcpClientStick")
	socket_stick.TcpClientStick()
}

func TcpClientSafe() {
	fmt.Println("启动 TcpClientSafe")
	socket_stick.TcpClientSafe()
}

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Printf("请选择启动服务(default 1)\n" +
		"1.TcpClient \n" +
		"2.UdpClient \n" +
		"3.TcpClientStick \n" +
		"4.TcpClientSafe \n" +
		"5.Q(退出) \n")
	str, err := input.ReadString('\n')
	if err != nil {
		log.Fatal("socket_client err -> ", err)
	}
	str = strings.Trim(str, "\r\n")
	switch str {
	case "1":
		TcpClient()
	case "2":
		UdpClient()
	case "3":
		TcpClientStick()
	case "4":
		TcpClientSafe()
	case "Q", "q", "5":
		break
	default:
		TcpClient()
	}
	fmt.Println("socket_client close")
}
