package socket_stick

import (
	"Go-100days/codes/sockets"
	"Go-100days/codes/sockets/proto"
	"fmt"
	"log"
	"net"
)

func TcpClientSafe()  {
	label := "TcpClientSafe -> "
	conn, err := net.Dial("tcp", sockets.ADDR)
	if err != nil {
		log.Fatal(label, err)
	}
	fmt.Println(label, sockets.ADDR)
	defer conn.Close()
	strs := []string{
		"风急天高猿啸哀，",
		"渚清沙白鸟飞回。",
		"无边落木萧萧下，",
		"不尽长江滚滚来。",
		"万里悲秋常作客，",
		"百年多病独登台。",
		"艰难苦恨繁霜鬓，",
		"潦倒新停浊酒杯。",
	}
	for i:=1; i<=10; i++ {
		for _, s := range strs {
			data := []byte(s)
			data = proto.Encode(data)
			conn.Write(data)
			fmt.Println(label, s)
		}
	}
}