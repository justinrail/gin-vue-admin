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

type SiteApi struct {
}

var siteService = service.ServiceGroupApp.HubServiceGroup.SiteService


// CreateSite 创建Site
// @Tags Site
// @Summary 创建Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Site true "创建Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /site/createSite [post]
func (siteApi *SiteApi) CreateSite(c *gin.Context) {
	var site hub.Site
	_ = c.ShouldBindJSON(&site)
	if err := siteService.CreateSite(site); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSite 删除Site
// @Tags Site
// @Summary 删除Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Site true "删除Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /site/deleteSite [delete]
func (siteApi *SiteApi) DeleteSite(c *gin.Context) {
	var site hub.Site
	_ = c.ShouldBindJSON(&site)
	if err := siteService.DeleteSite(site); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSiteByIds 批量删除Site
// @Tags Site
// @Summary 批量删除Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /site/deleteSiteByIds [delete]
func (siteApi *SiteApi) DeleteSiteByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := siteService.DeleteSiteByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSite 更新Site
// @Tags Site
// @Summary 更新Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Site true "更新Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /site/updateSite [put]
func (siteApi *SiteApi) UpdateSite(c *gin.Context) {
	var site hub.Site
	_ = c.ShouldBindJSON(&site)
	if err := siteService.UpdateSite(site); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSite 用id查询Site
// @Tags Site
// @Summary 用id查询Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hub.Site true "用id查询Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /site/findSite [get]
func (siteApi *SiteApi) FindSite(c *gin.Context) {
	var site hub.Site
	_ = c.ShouldBindQuery(&site)
	if err, resite := siteService.GetSite(site.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resite": resite}, c)
	}
}

// GetSiteList 分页获取Site列表
// @Tags Site
// @Summary 分页获取Site列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.SiteSearch true "分页获取Site列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /site/getSiteList [get]
func (siteApi *SiteApi) GetSiteList(c *gin.Context) {
	var pageInfo hubReq.SiteSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := siteService.GetSiteInfoList(pageInfo); err != nil {
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
