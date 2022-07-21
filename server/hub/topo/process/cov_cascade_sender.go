package process

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/topo/net"
	flow "github.com/trustmaster/goflow"
)

type COVCascadeSender struct {
	flow.Component
	In chan []*domain.COV
}

func (sender *COVCascadeSender) Process() {

	if !net.EnableMQTT {
		return
	}

	// for covs := range sender.In {
	// 	net.MQTTClient.Publish("ark/cov", 0, true, "covs")
	// }

	for covs := range sender.In {
		if len(covs) > 0 {
			marshal, err := json.Marshal(&covs)
			if err != nil {
				fmt.Println("covs json.Marshal err:", err)
			}
			net.MQTTClient.Publish("ark/cov", 0, true, marshal)
		}

	}
}
