package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type LivePointApi struct {
}

var livePointService = service.ServiceGroupApp.HubServiceGroup.LivePointService

// GetLivePointsByDevice 获取设备实时数据列表
// @Tags LivePoint
// @Summary 获取设备实时数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.GatewayDeviceParam true "获取设备实时数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livepoint/getdevicelivepoints [get]
func (livePointApi *LivePointApi) GetLivePointsByDevice(c *gin.Context) {
	livePoints := livePointService.GetLivePointsByDevice(queryInt(c, "deviceID"))
	response.OkWithDetailed(livePoints, "获取设备实时数据成功", c)
}
