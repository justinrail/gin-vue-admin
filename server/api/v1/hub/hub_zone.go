package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hub"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ZoneApi struct {
}

var zoneService = service.ServiceGroupApp.HubServiceGroup.ZoneService


// CreateZone 创建Zone
// @Tags Zone
// @Summary 创建Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Zone true "创建Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zone/createZone [post]
func (zoneApi *ZoneApi) CreateZone(c *gin.Context) {
	var zone hub.Zone
	_ = c.ShouldBindJSON(&zone)
	if err := zoneService.CreateZone(zone); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZone 删除Zone
// @Tags Zone
// @Summary 删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Zone true "删除Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zone/deleteZone [delete]
func (zoneApi *ZoneApi) DeleteZone(c *gin.Context) {
	var zone hub.Zone
	_ = c.ShouldBindJSON(&zone)
	if err := zoneService.DeleteZone(zone); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZoneByIds 批量删除Zone
// @Tags Zone
// @Summary 批量删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zone/deleteZoneByIds [delete]
func (zoneApi *ZoneApi) DeleteZoneByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := zoneService.DeleteZoneByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZone 更新Zone
// @Tags Zone
// @Summary 更新Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Zone true "更新Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zone/updateZone [put]
func (zoneApi *ZoneApi) UpdateZone(c *gin.Context) {
	var zone hub.Zone
	_ = c.ShouldBindJSON(&zone)
	if err := zoneService.UpdateZone(zone); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZone 用id查询Zone
// @Tags Zone
// @Summary 用id查询Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hub.Zone true "用id查询Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zone/findZone [get]
func (zoneApi *ZoneApi) FindZone(c *gin.Context) {
	var zone hub.Zone
	_ = c.ShouldBindQuery(&zone)
	if err, rezone := zoneService.GetZone(zone.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezone": rezone}, c)
	}
}

// GetZoneList 分页获取Zone列表
// @Tags Zone
// @Summary 分页获取Zone列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.ZoneSearch true "分页获取Zone列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zone/getZoneList [get]
func (zoneApi *ZoneApi) GetZoneList(c *gin.Context) {
	var pageInfo hubReq.ZoneSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := zoneService.GetZoneInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
