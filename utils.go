package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

/*工具方法*/

//将int转换成二进制，再转为【】byte
func IntToHex (num int64) []byte {
	buff := new(bytes.Buffer)//创建buffer缓冲器，可存取东西
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}