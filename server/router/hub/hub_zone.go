package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZoneRouter struct {
}

// InitZoneRouter 初始化 Zone 路由信息
func (s *ZoneRouter) InitZoneRouter(Router *gin.RouterGroup) {
	zoneRouter := Router.Group("zone").Use(middleware.OperationRecord())
	zoneRouterWithoutRecord := Router.Group("zone")
	var zoneApi = v1.ApiGroupApp.HubApiGroup.ZoneApi
	{
		zoneRouter.POST("createZone", zoneApi.CreateZone)   // 新建Zone
		zoneRouter.DELETE("deleteZone", zoneApi.DeleteZone) // 删除Zone
		zoneRouter.DELETE("deleteZoneByIds", zoneApi.DeleteZoneByIds) // 批量删除Zone
		zoneRouter.PUT("updateZone", zoneApi.UpdateZone)    // 更新Zone
	}
	{
		zoneRouterWithoutRecord.GET("findZone", zoneApi.FindZone)        // 根据ID获取Zone
		zoneRouterWithoutRecord.GET("getZoneList", zoneApi.GetZoneList)  // 获取Zone列表
	}
}
