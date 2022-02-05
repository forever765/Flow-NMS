import service from '@/utils/request'

// @Tags systrm
// @Summary 获取流量数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /charts/getTrafficData [post]
export const getTrafficData = () => {
  return service({
    url: '/charts/getTrafficData',
    method: 'post'
  })
}
