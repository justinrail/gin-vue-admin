import service from '@/utils/request'


// @Tags Gateway Device
// @Summary 获取gatewaydevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取gatewaydevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gatewaydevicetree/getgatewaydevicetreebody [get]
export const getGatewayDeviceTree = (params) => {
  return service({
    url: '/gatewaydevicetree/getgatewaydevicetreebody',
    method: 'get',
    params
  })
}
