package main

import (
	"Go-100days/codes/sockets"
	"Go-100days/codes/sockets/socket_stick"
)

func TcpServer() {
	sockets.DefaultTcpServer()
}

func UdpServer() {
	sockets.DefaultUdpServer()
}

func TcpServerStick() {
	socket_stick.TcpServerStick()
}

func TcpServerSafe() {
	socket_stick.TcpServerSafe()
}

func main() {
	TcpServerSafe()
}
