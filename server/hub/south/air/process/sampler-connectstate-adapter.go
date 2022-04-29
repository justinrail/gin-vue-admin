package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

type SamplerConnectStateAdapter struct {
	flow.Component
	In <-chan *protocol.PipelineObjectPair
}

func (adapter *SamplerConnectStateAdapter) Process() {
	for pair := range adapter.In {
		response := &protocol.SamplerConnectStateResponse{}
		response.FromByteArray(pair.Data, protocol.HeadLength)

		//response级别的采集单元状态信息（需要测试下，这里不清楚协议实现）
		cos := &domain.COS{
			GatewayID:     int(response.MonitorUnitId),
			SamplerUnitID: int(response.SamplerUnitId),
			State:         int(response.ConnectState),
			UpdateTime:    int64(response.SampleTime),
		}

		bus.COSBus <- cos

		//处理SamplerConnectStateItems级别监控单元通讯状态
		//TODO 这里的监控单元的SamplerConnectStateResponseItem结构很奇怪，像是信号？
		//由于不是很理解具体的结构体含义，这块代码先注释在调试中确认
		// for _, stateItem := range response.SamplerConnectStateItems {
		// 	for _, stateResponseItem := range stateItem.SamplerConnectStateResponseItem {
		// 		cosInner := &domain.COS{
		// 			DeviceID:  stateResponseItem.EquipmentId,
		// 			GatewayID: int(response.MonitorUnitId),
		// 		}
		// 	}
		// }

		if pair.RequireAcknowledge {
			ack := protocol.CreateAckPacket(pair.SequenceNumber, 0, int(HubEngineID))
			protocol.SendUDPData(pair.Url, ack)
		}
	}

}
