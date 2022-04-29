package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

//Unpacker Air协议解包
type Unpacker struct {
	flow.Component
	In  <-chan *protocol.AddressDataPair
	Out chan<- *protocol.AddressDataPair

	routerPacketCount int // processor stateful variable
}

//OnIn Request handler
func (unpacker *Unpacker) Process() {
	unpacker.routerPacketCount++
	for pair := range unpacker.In {
		if protocol.IsAirPacket(pair.Data) {
			unpacker.Out <- pair
		}
	}

	//metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	// if protocol.IsAirPacket(pair.Data) {
	// 	unpacker.Out <- pair
	// }
}
