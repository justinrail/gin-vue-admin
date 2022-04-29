import service from '@/utils/request'

// @Tags ZoneGroup
// @Summary 创建ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZoneGroup true "创建ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zoneGroup/createZoneGroup [post]
export const createZoneGroup = (data) => {
  return service({
    url: '/zoneGroup/createZoneGroup',
    method: 'post',
    data
  })
}

// @Tags ZoneGroup
// @Summary 删除ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZoneGroup true "删除ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zoneGroup/deleteZoneGroup [delete]
export const deleteZoneGroup = (data) => {
  return service({
    url: '/zoneGroup/deleteZoneGroup',
    method: 'delete',
    data
  })
}

// @Tags ZoneGroup
// @Summary 删除ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zoneGroup/deleteZoneGroup [delete]
export const deleteZoneGroupByIds = (data) => {
  return service({
    url: '/zoneGroup/deleteZoneGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags ZoneGroup
// @Summary 更新ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZoneGroup true "更新ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zoneGroup/updateZoneGroup [put]
export const updateZoneGroup = (data) => {
  return service({
    url: '/zoneGroup/updateZoneGroup',
    method: 'put',
    data
  })
}

// @Tags ZoneGroup
// @Summary 用id查询ZoneGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZoneGroup true "用id查询ZoneGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zoneGroup/findZoneGroup [get]
export const findZoneGroup = (params) => {
  return service({
    url: '/zoneGroup/findZoneGroup',
    method: 'get',
    params
  })
}

// @Tags ZoneGroup
// @Summary 分页获取ZoneGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ZoneGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zoneGroup/getZoneGroupList [get]
export const getZoneGroupList = (params) => {
  return service({
    url: '/zoneGroup/getZoneGroupList',
    method: 'get',
    params
  })
}
