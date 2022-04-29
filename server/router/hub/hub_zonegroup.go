package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZoneGroupRouter struct {
}

// InitZoneGroupRouter 初始化 ZoneGroup 路由信息
func (s *ZoneGroupRouter) InitZoneGroupRouter(Router *gin.RouterGroup) {
	zoneGroupRouter := Router.Group("zoneGroup").Use(middleware.OperationRecord())
	zoneGroupRouterWithoutRecord := Router.Group("zoneGroup")
	var zoneGroupApi = v1.ApiGroupApp.HubApiGroup.ZoneGroupApi
	{
		zoneGroupRouter.POST("createZoneGroup", zoneGroupApi.CreateZoneGroup)   // 新建ZoneGroup
		zoneGroupRouter.DELETE("deleteZoneGroup", zoneGroupApi.DeleteZoneGroup) // 删除ZoneGroup
		zoneGroupRouter.DELETE("deleteZoneGroupByIds", zoneGroupApi.DeleteZoneGroupByIds) // 批量删除ZoneGroup
		zoneGroupRouter.PUT("updateZoneGroup", zoneGroupApi.UpdateZoneGroup)    // 更新ZoneGroup
	}
	{
		zoneGroupRouterWithoutRecord.GET("findZoneGroup", zoneGroupApi.FindZoneGroup)        // 根据ID获取ZoneGroup
		zoneGroupRouterWithoutRecord.GET("getZoneGroupList", zoneGroupApi.GetZoneGroupList)  // 获取ZoneGroup列表
	}
}
