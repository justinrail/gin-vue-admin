package protocol

import "net"

//AddressDataPair UDP连接数据对象
type AddressDataPair struct {
	Address *net.UDPAddr
	Data    []byte
}
