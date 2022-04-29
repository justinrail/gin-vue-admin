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

type CenterApi struct {
}

var centerService = service.ServiceGroupApp.HubServiceGroup.CenterService


// CreateCenter 创建Center
// @Tags Center
// @Summary 创建Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Center true "创建Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /center/createCenter [post]
func (centerApi *CenterApi) CreateCenter(c *gin.Context) {
	var center hub.Center
	_ = c.ShouldBindJSON(&center)
	if err := centerService.CreateCenter(center); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCenter 删除Center
// @Tags Center
// @Summary 删除Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Center true "删除Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /center/deleteCenter [delete]
func (centerApi *CenterApi) DeleteCenter(c *gin.Context) {
	var center hub.Center
	_ = c.ShouldBindJSON(&center)
	if err := centerService.DeleteCenter(center); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCenterByIds 批量删除Center
// @Tags Center
// @Summary 批量删除Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /center/deleteCenterByIds [delete]
func (centerApi *CenterApi) DeleteCenterByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := centerService.DeleteCenterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCenter 更新Center
// @Tags Center
// @Summary 更新Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hub.Center true "更新Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /center/updateCenter [put]
func (centerApi *CenterApi) UpdateCenter(c *gin.Context) {
	var center hub.Center
	_ = c.ShouldBindJSON(&center)
	if err := centerService.UpdateCenter(center); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCenter 用id查询Center
// @Tags Center
// @Summary 用id查询Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hub.Center true "用id查询Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /center/findCenter [get]
func (centerApi *CenterApi) FindCenter(c *gin.Context) {
	var center hub.Center
	_ = c.ShouldBindQuery(&center)
	if err, recenter := centerService.GetCenter(center.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recenter": recenter}, c)
	}
}

// GetCenterList 分页获取Center列表
// @Tags Center
// @Summary 分页获取Center列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hubReq.CenterSearch true "分页获取Center列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /center/getCenterList [get]
func (centerApi *CenterApi) GetCenterList(c *gin.Context) {
	var pageInfo hubReq.CenterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := centerService.GetCenterInfoList(pageInfo); err != nil {
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
