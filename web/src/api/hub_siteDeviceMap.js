import service from '@/utils/request'

// @Tags SiteDeviceMap
// @Summary 创建SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SiteDeviceMap true "创建SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteDeviceMap/createSiteDeviceMap [post]
export const createSiteDeviceMap = (data) => {
  return service({
    url: '/siteDeviceMap/createSiteDeviceMap',
    method: 'post',
    data
  })
}

// @Tags SiteDeviceMap
// @Summary 删除SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SiteDeviceMap true "删除SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /siteDeviceMap/deleteSiteDeviceMap [delete]
export const deleteSiteDeviceMap = (data) => {
  return service({
    url: '/siteDeviceMap/deleteSiteDeviceMap',
    method: 'delete',
    data
  })
}

// @Tags SiteDeviceMap
// @Summary 删除SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /siteDeviceMap/deleteSiteDeviceMap [delete]
export const deleteSiteDeviceMapByIds = (data) => {
  return service({
    url: '/siteDeviceMap/deleteSiteDeviceMapByIds',
    method: 'delete',
    data
  })
}

// @Tags SiteDeviceMap
// @Summary 更新SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SiteDeviceMap true "更新SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /siteDeviceMap/updateSiteDeviceMap [put]
export const updateSiteDeviceMap = (data) => {
  return service({
    url: '/siteDeviceMap/updateSiteDeviceMap',
    method: 'put',
    data
  })
}

// @Tags SiteDeviceMap
// @Summary 用id查询SiteDeviceMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SiteDeviceMap true "用id查询SiteDeviceMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /siteDeviceMap/findSiteDeviceMap [get]
export const findSiteDeviceMap = (params) => {
  return service({
    url: '/siteDeviceMap/findSiteDeviceMap',
    method: 'get',
    params
  })
}

// @Tags SiteDeviceMap
// @Summary 分页获取SiteDeviceMap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SiteDeviceMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /siteDeviceMap/getSiteDeviceMapList [get]
export const getSiteDeviceMapList = (params) => {
  return service({
    url: '/siteDeviceMap/getSiteDeviceMapList',
    method: 'get',
    params
  })
}
