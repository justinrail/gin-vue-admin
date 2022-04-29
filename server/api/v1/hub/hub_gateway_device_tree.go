package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type GatewayDeviceTreeApi struct {
}

var gatewayDeviceTreeService = service.ServiceGroupApp.HubServiceGroup.GatewayDeviceTreeService

// GetGatewayDeviceTreeBody 获取网关及设备的树内容
// @Tags Tree
// @Summary 获取网关及设备的树内容
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query GatewayDeviceVM true "获取网关及设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gatewaydevicetree/getGatewayDeviceTreeBody [get]
func (gatewayDeviceTreeApi *GatewayDeviceTreeApi) GetGatewayDeviceTreeBody(c *gin.Context) {
	gateways := gatewayDeviceTreeService.GetAllGatewayWithDevices()
	response.OkWithDetailed(gateways, "获取网关及设备成功", c)
}
