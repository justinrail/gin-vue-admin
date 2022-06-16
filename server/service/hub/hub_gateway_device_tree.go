package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	hubRes "github.com/flipped-aurora/gin-vue-admin/server/model/hub/response"
)

type GatewayDeviceTreeService struct {
}

// GetAllGatewayWithDevices 获取GatewayWithDevices记录
func (srv *GatewayDeviceTreeService) GetAllGatewayWithDevices() []*hubRes.GatewayDevice {
	gateways := shadow.GetGateways()
	tree := make([]*hubRes.GatewayDevice, 0)

	for index := range gateways {
		gw := gateways[index]

		gatewayVM := &hubRes.GatewayDevice{
			GatewayName:    gw.Name,
			GatewayID:      gw.GatewayID,
			GatewayAddress: gw.IP,
		}
		gatewayVM.Devices = make([]*hubRes.LiteDevice, 0)

		for devIndex := range gw.Devices {
			dev := gw.Devices[devIndex]

			deviceVM := &hubRes.LiteDevice{
				DeviceID:   dev.DeviceID,
				DeviceName: dev.DeviceName,
			}

			gatewayVM.Devices = append(gatewayVM.Devices, deviceVM)
		}

		tree = append(tree, gatewayVM)
	}
	return tree
}
