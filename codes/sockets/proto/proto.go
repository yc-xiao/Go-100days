package proto

import (
	"bytes"
	"encoding/binary"
)

func Encode(data []byte) []byte {
	lenght := int32(len(data))
	pkg := bytes.NewBuffer([]byte{})
	binary.Write(pkg, binary.LittleEndian, lenght)
	binary.Write(pkg, binary.LittleEndian, data)
	return pkg.Bytes()
}

func Decode(data []byte) [][]byte {
	var datas [][]byte
	var lenght int32
	index := 0
	for i := 0; index+1 < len(data); i++ {
		lenghtByte := data[index : index+4]
		index += 4
		bytesBuffer := bytes.NewBuffer(lenghtByte)
		binary.Read(bytesBuffer, binary.LittleEndian, &lenght)
		next := int(lenght)
		datas = append(datas, data[index:index+next])
		index += next
	}
	return datas
}
