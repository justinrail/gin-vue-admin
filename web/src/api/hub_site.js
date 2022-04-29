import service from '@/utils/request'

// @Tags Site
// @Summary 创建Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Site true "创建Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /site/createSite [post]
export const createSite = (data) => {
  return service({
    url: '/site/createSite',
    method: 'post',
    data
  })
}

// @Tags Site
// @Summary 删除Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Site true "删除Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /site/deleteSite [delete]
export const deleteSite = (data) => {
  return service({
    url: '/site/deleteSite',
    method: 'delete',
    data
  })
}

// @Tags Site
// @Summary 删除Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /site/deleteSite [delete]
export const deleteSiteByIds = (data) => {
  return service({
    url: '/site/deleteSiteByIds',
    method: 'delete',
    data
  })
}

// @Tags Site
// @Summary 更新Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Site true "更新Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /site/updateSite [put]
export const updateSite = (data) => {
  return service({
    url: '/site/updateSite',
    method: 'put',
    data
  })
}

// @Tags Site
// @Summary 用id查询Site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Site true "用id查询Site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /site/findSite [get]
export const findSite = (params) => {
  return service({
    url: '/site/findSite',
    method: 'get',
    params
  })
}

// @Tags Site
// @Summary 分页获取Site列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Site列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /site/getSiteList [get]
export const getSiteList = (params) => {
  return service({
    url: '/site/getSiteList',
    method: 'get',
    params
  })
}
