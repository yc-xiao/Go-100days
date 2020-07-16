package sockets

import (
	"net"
	"strconv"
	"strings"
)

// 公开变量
const ADDR = "localhost:9000"

// upd addr -> ip, port
func strToip(addr string) (IP net.IP, Port int, err error) {
	strs := strings.Split(addr, ":")
	Port, err = strconv.Atoi(strs[1])
	if err != nil {
		return
	}
	if strs[0] == "localhost" {
		strs[0] = "127.0.0.1"
	}
	strs = strings.Split(strs[0], ".")
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, 0, err
		}
		IP = append(IP, uint8(n))
	}
	return
}
