package protocol

import (
	"fmt"
	"net"
)

//SendUDPData 发送数据包到udp server
func SendUDPData(address string, data []byte) {
	conn, err := net.Dial("udp", address)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	conn.Write(data)

	//fmt.Println("send msg")

	//var msg [20]byte
	//conn.Read(msg[0:])

	//fmt.Println("msg is", string(msg[0:10]))
}
