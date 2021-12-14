import service from '@/utils/request'

// @Tags dashboard
// @Summary 获取首页顶部卡片栏位的数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /dashboard/getTopCardData [post]
export const getTopCardData = () => {
  return service({
    url: '/dashboard/getTopCardData',
    method: 'post'
  })
}
