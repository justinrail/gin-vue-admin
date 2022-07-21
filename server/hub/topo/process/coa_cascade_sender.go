package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/topo/net"
	flow "github.com/trustmaster/goflow"
)

type COACascadeSender struct {
	flow.Component
	In chan *domain.COA
}

func (sender *COACascadeSender) Process() {

	if !net.EnableMQTT {
		return
	}

	for coa := range sender.In {
		net.MQTTClient.Publish("ark/coa", 1, true, coa)
	}
}
