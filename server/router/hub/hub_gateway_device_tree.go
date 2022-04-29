package hub

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type GatwayDeviceTreeRouter struct {
}

// InitGatwayDeviceTreeRouter 初始化 GatwayDeviceTree 路由信息
func (s *CenterRouter) InitGatwayDeviceTreeRouter(Router *gin.RouterGroup) {
	treeRouterWithoutRecord := Router.Group("gatewaydevicetree")
	var treeApi = v1.ApiGroupApp.HubApiGroup.GatewayDeviceTreeApi
	{
		treeRouterWithoutRecord.GET("getgatewaydevicetreebody", treeApi.GetGatewayDeviceTreeBody) // 获取Tree内容
	}
}
