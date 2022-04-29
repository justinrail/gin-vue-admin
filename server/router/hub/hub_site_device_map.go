package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SiteDeviceMapRouter struct {
}

// InitSiteDeviceMapRouter 初始化 SiteDeviceMap 路由信息
func (s *SiteDeviceMapRouter) InitSiteDeviceMapRouter(Router *gin.RouterGroup) {
	siteDeviceMapRouter := Router.Group("siteDeviceMap").Use(middleware.OperationRecord())
	siteDeviceMapRouterWithoutRecord := Router.Group("siteDeviceMap")
	var siteDeviceMapApi = v1.ApiGroupApp.HubApiGroup.SiteDeviceMapApi
	{
		siteDeviceMapRouter.POST("createSiteDeviceMap", siteDeviceMapApi.CreateSiteDeviceMap)   // 新建SiteDeviceMap
		siteDeviceMapRouter.DELETE("deleteSiteDeviceMap", siteDeviceMapApi.DeleteSiteDeviceMap) // 删除SiteDeviceMap
		siteDeviceMapRouter.DELETE("deleteSiteDeviceMapByIds", siteDeviceMapApi.DeleteSiteDeviceMapByIds) // 批量删除SiteDeviceMap
		siteDeviceMapRouter.PUT("updateSiteDeviceMap", siteDeviceMapApi.UpdateSiteDeviceMap)    // 更新SiteDeviceMap
	}
	{
		siteDeviceMapRouterWithoutRecord.GET("findSiteDeviceMap", siteDeviceMapApi.FindSiteDeviceMap)        // 根据ID获取SiteDeviceMap
		siteDeviceMapRouterWithoutRecord.GET("getSiteDeviceMapList", siteDeviceMapApi.GetSiteDeviceMapList)  // 获取SiteDeviceMap列表
	}
}
