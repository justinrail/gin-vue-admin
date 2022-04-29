package protocol

import (
	"fmt"
	"net"
)

//UDPReceiver UDP协议接受器(服务端)
type UDPReceiver struct {
	Port       string
	BufferSize int
	OutChannel chan *AddressDataPair
	Conn       *net.UDPConn
	Address    *net.UDPAddr
}

//NewUDPReceiver 创建新的接收器
func NewUDPReceiver(port string, bufferSize int, outChannel chan *AddressDataPair, listenAddress string) (*UDPReceiver, error) {

	receiver := &UDPReceiver{}
	addr, err := net.ResolveUDPAddr("udp", listenAddress)
	if err != nil {
		return nil, err
	}

	receiver.Port = port
	receiver.OutChannel = outChannel
	receiver.BufferSize = bufferSize
	receiver.Address = addr

	return receiver, nil
}

//Serve 启动服务
func (receiver *UDPReceiver) Serve() {

	conn, err := net.ListenUDP("udp", receiver.Address)
	if err != nil {
		fmt.Println("udp receiver listen error:", err)
		return
	}
	receiver.Conn = conn
	//defer conn.Close()

	buf := make([]byte, receiver.BufferSize)
	for {
		n, addr, _ := receiver.Conn.ReadFromUDP(buf)
		data := make([]byte, n)
		copy(data, buf)
		receiver.OutChannel <- &AddressDataPair{Address: addr, Data: data}
		//fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}
