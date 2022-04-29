package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CenterRouter struct {
}

// InitCenterRouter 初始化 Center 路由信息
func (s *CenterRouter) InitCenterRouter(Router *gin.RouterGroup) {
	centerRouter := Router.Group("center").Use(middleware.OperationRecord())
	centerRouterWithoutRecord := Router.Group("center")
	var centerApi = v1.ApiGroupApp.HubApiGroup.CenterApi
	{
		centerRouter.POST("createCenter", centerApi.CreateCenter)   // 新建Center
		centerRouter.DELETE("deleteCenter", centerApi.DeleteCenter) // 删除Center
		centerRouter.DELETE("deleteCenterByIds", centerApi.DeleteCenterByIds) // 批量删除Center
		centerRouter.PUT("updateCenter", centerApi.UpdateCenter)    // 更新Center
	}
	{
		centerRouterWithoutRecord.GET("findCenter", centerApi.FindCenter)        // 根据ID获取Center
		centerRouterWithoutRecord.GET("getCenterList", centerApi.GetCenterList)  // 获取Center列表
	}
}
