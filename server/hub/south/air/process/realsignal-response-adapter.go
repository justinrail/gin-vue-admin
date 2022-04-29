package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

type RealSignalResponseAdapter struct {
	flow.Component
	In <-chan *protocol.PipelineObjectPair
}

//OnIn Request handler
func (adapter *RealSignalResponseAdapter) Process() {

	for pair := range adapter.In {
		response := &protocol.RealSignalResponse{}
		response.FromByteArray(pair.Data)
		bus.COVBus <- response.GetCOVs()
	}

}
