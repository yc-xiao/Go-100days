package main

import (
	"Go-100days/codes/sockets/proto"
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//sockets.DefaultUdpClient()
	b := []byte{0,0,0,5}
	nb := proto.Encode(b)
	//bbs := proto.Decode(nb)
	buffer := bytes.NewBuffer(nb)
	binary.Write(buffer, binary.LittleEndian, proto.Encode(b))
	bbf := buffer.Bytes()
	bbs := proto.Decode(bbf)
	fmt.Println(bbs)
}

