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

type ZoneGroupApi struct {
}

var zoneGroupService = service.ServiceGroupApp.HubServiceGroup.ZoneGroupService


// CreateZoneGroup 创建ZoneGroup
// @Tags ZoneGroup
// @Summary 创建ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.ZoneGroup true "创建ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zoneGroup/createZoneGroup [post]
func (zoneGroupApi *ZoneGroupApi) CreateZoneGroup(c *gin.Context) {
	var zoneGroup hub.ZoneGroup
	_ = c.ShouldBindJSON(&zoneGroup)
	if err := zoneGroupService.CreateZoneGroup(zoneGroup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZoneGroup 删除ZoneGroup
// @Tags ZoneGroup
// @Summary 删除ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.ZoneGroup true "删除ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zoneGroup/deleteZoneGroup [delete]
func (zoneGroupApi *ZoneGroupApi) DeleteZoneGroup(c *gin.Context) {
	var zoneGroup hub.ZoneGroup
	_ = c.ShouldBindJSON(&zoneGroup)
	if err := zoneGroupService.DeleteZoneGroup(zoneGroup); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZoneGroupByIds 批量删除ZoneGroup
// @Tags ZoneGroup
// @Summary 批量删除ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zoneGroup/deleteZoneGroupByIds [delete]
func (zoneGroupApi *ZoneGroupApi) DeleteZoneGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := zoneGroupService.DeleteZoneGroupByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZoneGroup 更新ZoneGroup
// @Tags ZoneGroup
// @Summary 更新ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.ZoneGroup true "更新ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zoneGroup/updateZoneGroup [put]
func (zoneGroupApi *ZoneGroupApi) UpdateZoneGroup(c *gin.Context) {
	var zoneGroup hub.ZoneGroup
	_ = c.ShouldBindJSON(&zoneGroup)
	if err := zoneGroupService.UpdateZoneGroup(zoneGroup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZoneGroup 用id查询ZoneGroup
// @Tags ZoneGroup
// @Summary 用id查询ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hub.ZoneGroup true "用id查询ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zoneGroup/findZoneGroup [get]
func (zoneGroupApi *ZoneGroupApi) FindZoneGroup(c *gin.Context) {
	var zoneGroup hub.ZoneGroup
	_ = c.ShouldBindQuery(&zoneGroup)
	if err, rezoneGroup := zoneGroupService.GetZoneGroup(zoneGroup.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezoneGroup": rezoneGroup}, c)
	}
}

// GetZoneGroupList 分页获取ZoneGroup列表
// @Tags ZoneGroup
// @Summary 分页获取ZoneGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.ZoneGroupSearch true "分页获取ZoneGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zoneGroup/getZoneGroupList [get]
func (zoneGroupApi *ZoneGroupApi) GetZoneGroupList(c *gin.Context) {
	var pageInfo hubReq.ZoneGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := zoneGroupService.GetZoneGroupInfoList(pageInfo); err != nil {
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
