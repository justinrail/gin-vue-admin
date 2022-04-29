package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
)

type GatewayDeviceVM struct {
	GatewayID      int    `json:"gatewayID"`
	GatewayName    string `json:"gatewayName"`
	GatewayAddress string `json:"gatewayAddress"`
	Devices        []*DeviceVM
}

type DeviceVM struct {
	DeviceID   int    `json:"deviceID"`
	DeviceName string `json:"deviceName"`
}

type GatewayDeviceTreeService struct {
}

// GetAllGatewayWithDevices 获取GatewayWithDevices记录
func (srv *GatewayDeviceTreeService) GetAllGatewayWithDevices() []*GatewayDeviceVM {
	gateways := shadow.GetGateways()
	tree := make([]*GatewayDeviceVM, 0)

	for index := range gateways {
		gw := gateways[index]

		gatewayVM := &GatewayDeviceVM{
			GatewayName:    gw.Name,
			GatewayID:      gw.GatewayID,
			GatewayAddress: gw.IP,
		}
		gatewayVM.Devices = make([]*DeviceVM, 0)

		for devIndex := range gw.Devices {
			dev := gw.Devices[devIndex]

			deviceVM := &DeviceVM{
				DeviceID:   dev.DeviceID,
				DeviceName: dev.DeviceName,
			}

			gatewayVM.Devices = append(gatewayVM.Devices, deviceVM)
		}

		tree = append(tree, gatewayVM)
	}
	return tree
}
