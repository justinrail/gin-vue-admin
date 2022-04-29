package domain

import (
	"container/list"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/base"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/entity"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/util"
)

const (
	//GateWayRequestQueueSize gateway相关控制命令下发dto收集队列 默认长度为4
	GateWayRequestQueueSize = 4
	//GateWayLogQueueSize gateway级别日志队列
	GateWayLogQueueSize = 10
)

//Gateway Gateway Domain Object
type Gateway struct {
	GatewayID  int
	UUID       string
	Name       string
	SouthType  string
	IP         string
	ConState   int
	SynState   int
	DebugState bool
	Devices    map[int]*Device
	UpdateTime int64

	//RequestQueue Gateway，CoreSource，CorePoint的所有下发请求队列
	RequestQueue *util.EsQueue
	//Gateway下发队列日志
	RequestLogs *util.EsQueue
	//ResponseLogs Gateway反馈队列日志，反馈队列由chan进行收集处理，这里只记录最终日志
	ResponseLogs *util.EsQueue
	//PacketLogs 网关相关数据包的调试队列（什么时候Register，什么时候offline等切换）
	PacketLogs *list.List
}

//NewGateway create domain gateway
func NewGateway() *Gateway {
	gateway := Gateway{}
	gateway.ConState = base.GatewayConStateUnkown
	gateway.SynState = base.GatewaySynStateUnkown
	gateway.DebugState = false
	gateway.Devices = map[int]*Device{}
	gateway.RequestQueue = util.NewQueue(GateWayRequestQueueSize)
	gateway.RequestLogs = util.NewQueue(GateWayLogQueueSize)
	gateway.ResponseLogs = util.NewQueue(GateWayLogQueueSize)
	gateway.PacketLogs = list.New()
	return &gateway
}

//AppendPacketLogs update state
func (gateway *Gateway) AppendPacketLogs(cog *COG) {
	gateway.PacketLogs.PushBack(cog) // Enqueue

	if gateway.PacketLogs.Len() > GateWayLogQueueSize {
		e := gateway.PacketLogs.Front()

		if e != nil {
			gateway.PacketLogs.Remove(e)
		}
	}
}

func (gateway *Gateway) From(mu *entity.MonitorUnit) {
	gateway.GatewayID = mu.MonitorUnitId
	gateway.IP = mu.IpAddress
	gateway.ConState = base.GatewayConStateOffline //先默认设置为false，未来可能会取历史状态
	gateway.DebugState = false                     //先默认设置为false，未来可能会取历史状态
	gateway.Name = mu.MonitorUnitName
	gateway.UUID = mu.MonitorUnitCode
}

func (gateway *Gateway) UpdateConnectState(flag int) {
	switch flag {
	case base.GatewayConStateOnline:
		gateway.ConState = base.GatewayConStateOnline
	case base.GatewayConStateOffline:
		gateway.ConState = base.GatewayConStateOffline
	case base.GatewayConStateUnkown:
		gateway.ConState = base.GatewayConStateUnkown
	}
}
