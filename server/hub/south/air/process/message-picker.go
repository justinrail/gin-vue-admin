package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

//MessagePicker Air协议确认及转换
type MessagePicker struct {
	flow.Component
	In  <-chan *protocol.AddressDataPair
	Out chan<- *protocol.PipelineObjectPair
}

//OnIn Request handler
func (picker *MessagePicker) Process() {
	//metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	for pair := range picker.In {

		msg := picker.pickMessage(pair)

		picker.Out <- msg
	}

	//fmt.Println(msg)

	/* if packet.RequireAcknowledge {
		picker.sendPackACK(packet)
	}*/
}

func (picker *MessagePicker) sendPackACK(packet *protocol.AirPacket) {

}

func (picker *MessagePicker) pickMessage(pair *protocol.AddressDataPair) *protocol.PipelineObjectPair {
	packet := protocol.NewAirPacket(pair.Data)
	msg := &protocol.PipelineObjectPair{
		PipelineType:       packet.PipelineType,
		RequireAcknowledge: packet.RequireAcknowledge,
		Data:               packet.RawPacket,
		Ttl:                packet.TTL,
		Url:                pair.Address.String(),
		SequenceNumber:     packet.SequenceNumber,
		SourceHostId:       packet.SourceHostID,
		DestinationHostId:  packet.DestinationHostID,
		MessageType:        packet.MessageType,
	}

	return msg
}
