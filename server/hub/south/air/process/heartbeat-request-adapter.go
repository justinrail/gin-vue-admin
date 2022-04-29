package process

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

//HeartbeatRequestAdapter 心跳请求包处理
type HeartbeatRequestAdapter struct {
	flow.Component
	In <-chan *protocol.PipelineObjectPair
}

func (adapter *HeartbeatRequestAdapter) Process() {
	//metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	for pair := range adapter.In {
		request := &protocol.HeartBeatRequest{}
		request.FromByteArray(pair.Data, protocol.HeadLength)
		//todo 心跳业务处理 如gateway状态等

		//发送心跳回复
		replayHeartBeat(pair)
	}
}

func replayHeartBeat(pipelineObjectPair *protocol.PipelineObjectPair) {
	data := createHeartBeatResponsePacket(pipelineObjectPair)
	protocol.SendUDPData(pipelineObjectPair.Url, data)
}

func createHeartResponse(pipelineObjectPair *protocol.PipelineObjectPair) *protocol.HeartBeatResponse {
	heartBeatResponse := &protocol.HeartBeatResponse{}
	heartBeatResponse.MessageType = 211
	heartBeatResponse.HeartbeatTime = protocol.GetNTPTime()
	return heartBeatResponse
}

func createRDSHeartResponse(pipelineObjectPair *protocol.PipelineObjectPair) *protocol.RDSHeartBeatResponse {
	heartBeatResponse := &protocol.RDSHeartBeatResponse{}
	heartBeatResponse.MessageType = 214
	heartBeatResponse.HeartbeatTime = protocol.GetNTPTime()
	heartBeatResponse.RealtimeInterval = 30
	heartBeatResponse.RealtimePort = 9000
	return heartBeatResponse
}

func createHeartBeatResponsePacket(pipelineObjectPair *protocol.PipelineObjectPair) []byte {
	pack := &protocol.AirPacket{}
	pack.DestinationHostID = pipelineObjectPair.DestinationHostId
	pack.SourceHostID = global.GVA_VP.GetInt32("hub.engine-id") //主机ID
	pack.PipelineType = 0
	pack.ProtocolType = 1
	heartBeatResponse := createHeartResponse(pipelineObjectPair)
	pack.RawPacket = heartBeatResponse.ToByteArray()
	pack.MessageType = 211
	pack.RawObject = heartBeatResponse
	if strings.Contains(pipelineObjectPair.Url, "6000") {
		heartBeatResponse := createHeartResponse(pipelineObjectPair)
		pack.RawPacket = heartBeatResponse.ToByteArray()
		pack.MessageType = 211
		pack.RawObject = heartBeatResponse
	} else {
		heartBeatResponse := createRDSHeartResponse(pipelineObjectPair)
		pack.RawPacket = heartBeatResponse.ToByteArray()
		pack.MessageType = 214
		pack.RawObject = heartBeatResponse
	}
	data := pack.Pack()
	return data
}
