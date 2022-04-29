package air

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	"github.com/trustmaster/goflow"
)

var Vars AirVars

//系统标准B接口的全局对象体
type AirVars struct {
	Graph           *goflow.Graph
	RawPackets      chan *protocol.AddressDataPair
	DefaultReceiver *protocol.UDPReceiver
}

//RDS的实时数据不经过正常流程，直接通过协议即可拿到
//如果做测试，也可以单独先调RDS对应的代码，这样能更快的调通协议格式和包解析
//下面的部分代码就是当时调RDS的代码，因标准协议已经部分开发了，所以RDS的暂时注释了
//var rdsReceiver     *protocol.UDPReceiver

const (
	//UDPBufferSize UDP 协议缓冲
	UDPBufferSize = 2048
	//PacketCacheLength 包缓冲区长度
	PacketCacheLength = 10240
)

func InitVars() {
	Vars = AirVars{}
	Vars.Graph = goflow.NewGraph()
	Vars.RawPackets = make(chan *protocol.AddressDataPair, PacketCacheLength)
}
