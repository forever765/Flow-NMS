import service from '@/utils/request'

// @Tags systrm
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /charts/getTraffic [post]
export const getTraffic = () => {
  return service({
    url: '/charts/getTraffic',
    method: 'post'
  })
}