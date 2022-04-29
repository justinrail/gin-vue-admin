import service from '@/utils/request'

// @Tags Zone
// @Summary 创建Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Zone true "创建Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zone/createZone [post]
export const createZone = (data) => {
  return service({
    url: '/zone/createZone',
    method: 'post',
    data
  })
}

// @Tags Zone
// @Summary 删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Zone true "删除Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zone/deleteZone [delete]
export const deleteZone = (data) => {
  return service({
    url: '/zone/deleteZone',
    method: 'delete',
    data
  })
}

// @Tags Zone
// @Summary 删除Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zone/deleteZone [delete]
export const deleteZoneByIds = (data) => {
  return service({
    url: '/zone/deleteZoneByIds',
    method: 'delete',
    data
  })
}

// @Tags Zone
// @Summary 更新Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Zone true "更新Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zone/updateZone [put]
export const updateZone = (data) => {
  return service({
    url: '/zone/updateZone',
    method: 'put',
    data
  })
}

// @Tags Zone
// @Summary 用id查询Zone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Zone true "用id查询Zone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zone/findZone [get]
export const findZone = (params) => {
  return service({
    url: '/zone/findZone',
    method: 'get',
    params
  })
}

// @Tags Zone
// @Summary 分页获取Zone列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Zone列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zone/getZoneList [get]
export const getZoneList = (params) => {
  return service({
    url: '/zone/getZoneList',
    method: 'get',
    params
  })
}
