package process

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	flow "github.com/trustmaster/goflow"
)

type DeviceStater struct {
	flow.Component
	In <-chan *domain.COD
}

func (stater *DeviceStater) Process() {

	for cod := range stater.In {
		device, existDevice := shadow.GetDeviceByID(cod.DeviceID)

		if existDevice {
			//更新网关通讯状态
			device.UpdateConnectState(cod.Flag)
			device.UpdateTime = time.Now().Unix()
			device.AppendPacketLogs(cod)
		}
	}
}
