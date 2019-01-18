package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	Test()
}

func Test() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	for {
		data, _ := Encode("10086")
		_, err := conn.Write(data)
		fmt.Println(err)
		data2,_ := Encode("10010")
		_, err = conn.Write(data2)
		_, err = conn.Write(data2)
		_, err = conn.Write(data)
		time.Sleep(time.Second * 4)
		fmt.Println(err)
	}
}
func Encode(message string) ([]byte, error) {
	// 读取消息的长度
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}