import service from '@/utils/request'

// @Tags Center
// @Summary 创建Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Center true "创建Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /center/createCenter [post]
export const createCenter = (data) => {
  return service({
    url: '/center/createCenter',
    method: 'post',
    data
  })
}

// @Tags Center
// @Summary 删除Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Center true "删除Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /center/deleteCenter [delete]
export const deleteCenter = (data) => {
  return service({
    url: '/center/deleteCenter',
    method: 'delete',
    data
  })
}

// @Tags Center
// @Summary 删除Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /center/deleteCenter [delete]
export const deleteCenterByIds = (data) => {
  return service({
    url: '/center/deleteCenterByIds',
    method: 'delete',
    data
  })
}

// @Tags Center
// @Summary 更新Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Center true "更新Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /center/updateCenter [put]
export const updateCenter = (data) => {
  return service({
    url: '/center/updateCenter',
    method: 'put',
    data
  })
}

// @Tags Center
// @Summary 用id查询Center
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Center true "用id查询Center"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /center/findCenter [get]
export const findCenter = (params) => {
  return service({
    url: '/center/findCenter',
    method: 'get',
    params
  })
}

// @Tags Center
// @Summary 分页获取Center列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Center列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /center/getCenterList [get]
export const getCenterList = (params) => {
  return service({
    url: '/center/getCenterList',
    method: 'get',
    params
  })
}
