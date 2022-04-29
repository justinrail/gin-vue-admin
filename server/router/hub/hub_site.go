package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SiteRouter struct {
}

// InitSiteRouter 初始化 Site 路由信息
func (s *SiteRouter) InitSiteRouter(Router *gin.RouterGroup) {
	siteRouter := Router.Group("site").Use(middleware.OperationRecord())
	siteRouterWithoutRecord := Router.Group("site")
	var siteApi = v1.ApiGroupApp.HubApiGroup.SiteApi
	{
		siteRouter.POST("createSite", siteApi.CreateSite)   // 新建Site
		siteRouter.DELETE("deleteSite", siteApi.DeleteSite) // 删除Site
		siteRouter.DELETE("deleteSiteByIds", siteApi.DeleteSiteByIds) // 批量删除Site
		siteRouter.PUT("updateSite", siteApi.UpdateSite)    // 更新Site
	}
	{
		siteRouterWithoutRecord.GET("findSite", siteApi.FindSite)        // 根据ID获取Site
		siteRouterWithoutRecord.GET("getSiteList", siteApi.GetSiteList)  // 获取Site列表
	}
}
