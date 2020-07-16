package main

import (
	"Go-100days/codes/sockets"
	"Go-100days/codes/sockets/socket_stick"
)

func TcpClient() {
	sockets.DefaultTcpClient()
}

func UdpClient() {
	sockets.DefaultUdpClient()
}

func TcpClientStick() {
	socket_stick.TcpClientStick()
}

func TcpClientSafe() {
	socket_stick.TcpClientSafe()
}

func main() {
	TcpClientSafe()
}
