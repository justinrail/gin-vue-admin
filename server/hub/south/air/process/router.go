package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

//Router SiteWeb协议DTO分包
type Router struct {
	flow.Component
	In                       <-chan *protocol.PipelineObjectPair
	OutHeartbeatRequest      chan<- *protocol.PipelineObjectPair
	OutRegisterRequest       chan<- *protocol.PipelineObjectPair
	OutRealSignalResponse    chan<- *protocol.PipelineObjectPair
	OutEventResponse         chan<- *protocol.PipelineObjectPair
	OutHistorySignalResponse chan<- *protocol.PipelineObjectPair
	OutRelatedSignalResponse chan<- *protocol.PipelineObjectPair
	OutConfigFileResponse    chan<- *protocol.PipelineObjectPair
	OutSamplerConnectState   chan<- *protocol.PipelineObjectPair
}

//OnIn Request handler
func (router *Router) Process() {
	//metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	for msg := range router.In {
		//msg.Offset += protocol.SWDPOffset
		switch msg.MessageType {
		case 210:
			router.OutHeartbeatRequest <- msg
			break
		case 11:
			router.OutRealSignalResponse <- msg
			break
		case 70:
			router.OutRegisterRequest <- msg
			break
		case 40:
			router.OutEventResponse <- msg
			break
		case 30:
			router.OutHistorySignalResponse <- msg
			break
		case 33:
			router.OutRelatedSignalResponse <- msg
			break
		case 80:
			router.OutConfigFileResponse <- msg
		case 43:
			router.OutSamplerConnectState <- msg
		default:
			break
		}
	}
}
