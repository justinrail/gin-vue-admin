package hub

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LivePointRouter struct {
}

// LivePointRouter 初始化 LivePointRouter 路由信息
func (s *LivePointRouter) InitLivePointRouter(Router *gin.RouterGroup) {
	livePointRouter := Router.Group("livepoint")
	var treeApi = v1.ApiGroupApp.HubApiGroup.LivePointApi
	{
		livePointRouter.GET("getdevicelivepoints", treeApi.GetLivePointsByDevice) // 获取内容
	}
}
