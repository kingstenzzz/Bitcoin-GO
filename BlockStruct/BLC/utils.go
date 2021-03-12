package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(data int64) []byte {
	//	string := strconv.FormatInt(data, 10)//库的方法
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if err != nil {
		log.Panic("int2[]byte failed!")
	}
	return buffer.Bytes()
}
