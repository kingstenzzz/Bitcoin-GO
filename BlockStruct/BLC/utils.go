package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

//Json转切片
// "[\"k\"]" -to "[\"k,a\"]" -amount "[\"100\"]"
func JSONToSlice(jsonString string) []string {
	var strSlice []string
	if err := json.Unmarshal([]byte(jsonString), &strSlice); nil != err {
		log.Panicf("json to string faild %d \v", err)
	}
	return strSlice
}
