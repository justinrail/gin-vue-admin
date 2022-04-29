package process

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	flow "github.com/trustmaster/goflow"
)

type GatewayStater struct {
	flow.Component
	In <-chan *domain.COG
}

func (stater *GatewayStater) Process() {

	for cog := range stater.In {
		gateway, existGateway := shadow.GetGatewayByID(cog.GatewayID)

		if existGateway {
			//更新网关通讯状态
			gateway.UpdateConnectState(cog.Flag)

			gateway.UpdateTime = time.Now().Unix()
			gateway.AppendPacketLogs(cog)
		}
	}
}
