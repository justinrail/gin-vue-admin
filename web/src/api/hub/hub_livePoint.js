import service from '@/utils/request'


// @Tags LivePoint
// @Summary 获取设备livepoint实时数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取设备实时数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livepoint/getdevicelivepoints [get]
export const getDeviceLivePoints = (params) => {
  return service({
    url: '/livepoint/getdevicelivepoints',
    method: 'get',
    params
  })
}
