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

type SiteDeviceMapApi struct {
}

var siteDeviceMapService = service.ServiceGroupApp.HubServiceGroup.SiteDeviceMapService


// CreateSiteDeviceMap 创建SiteDeviceMap
// @Tags SiteDeviceMap
// @Summary 创建SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.SiteDeviceMap true "创建SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteDeviceMap/createSiteDeviceMap [post]
func (siteDeviceMapApi *SiteDeviceMapApi) CreateSiteDeviceMap(c *gin.Context) {
	var siteDeviceMap hub.SiteDeviceMap
	_ = c.ShouldBindJSON(&siteDeviceMap)
	if err := siteDeviceMapService.CreateSiteDeviceMap(siteDeviceMap); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSiteDeviceMap 删除SiteDeviceMap
// @Tags SiteDeviceMap
// @Summary 删除SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.SiteDeviceMap true "删除SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /siteDeviceMap/deleteSiteDeviceMap [delete]
func (siteDeviceMapApi *SiteDeviceMapApi) DeleteSiteDeviceMap(c *gin.Context) {
	var siteDeviceMap hub.SiteDeviceMap
	_ = c.ShouldBindJSON(&siteDeviceMap)
	if err := siteDeviceMapService.DeleteSiteDeviceMap(siteDeviceMap); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSiteDeviceMapByIds 批量删除SiteDeviceMap
// @Tags SiteDeviceMap
// @Summary 批量删除SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /siteDeviceMap/deleteSiteDeviceMapByIds [delete]
func (siteDeviceMapApi *SiteDeviceMapApi) DeleteSiteDeviceMapByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := siteDeviceMapService.DeleteSiteDeviceMapByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSiteDeviceMap 更新SiteDeviceMap
// @Tags SiteDeviceMap
// @Summary 更新SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.SiteDeviceMap true "更新SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /siteDeviceMap/updateSiteDeviceMap [put]
func (siteDeviceMapApi *SiteDeviceMapApi) UpdateSiteDeviceMap(c *gin.Context) {
	var siteDeviceMap hub.SiteDeviceMap
	_ = c.ShouldBindJSON(&siteDeviceMap)
	if err := siteDeviceMapService.UpdateSiteDeviceMap(siteDeviceMap); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSiteDeviceMap 用id查询SiteDeviceMap
// @Tags SiteDeviceMap
// @Summary 用id查询SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hub.SiteDeviceMap true "用id查询SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /siteDeviceMap/findSiteDeviceMap [get]
func (siteDeviceMapApi *SiteDeviceMapApi) FindSiteDeviceMap(c *gin.Context) {
	var siteDeviceMap hub.SiteDeviceMap
	_ = c.ShouldBindQuery(&siteDeviceMap)
	if err, resiteDeviceMap := siteDeviceMapService.GetSiteDeviceMap(siteDeviceMap.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resiteDeviceMap": resiteDeviceMap}, c)
	}
}

// GetSiteDeviceMapList 分页获取SiteDeviceMap列表
// @Tags SiteDeviceMap
// @Summary 分页获取SiteDeviceMap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.SiteDeviceMapSearch true "分页获取SiteDeviceMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteDeviceMap/getSiteDeviceMapList [get]
func (siteDeviceMapApi *SiteDeviceMapApi) GetSiteDeviceMapList(c *gin.Context) {
	var pageInfo hubReq.SiteDeviceMapSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := siteDeviceMapService.GetSiteDeviceMapInfoList(pageInfo); err != nil {
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
