import service from '@/utils/request'

// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
// export const getApiList = (data) => {
//   return service({
//     url: '/api/getApiList',
//     method: 'post',
//     data
//   })
// }

// @Tags Report
// @Summary 分页获取最新数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/getNewestData [post]
// {
//  page     int
//	pageSize int
// }
export const getNewestData = (data) => {
  return service({
    url: '/reports/getNewestData',
    method: 'post',
    data
  })
}

// @Tags Report
// @Summary 分页获取src使用流量TopN
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/getSrcTopN [post]
// {
//  page     int
//	pageSize int
// }
export const getTopN = (data) => {
  return service({
    url: '/reports/getTopN',
    method: 'post',
    data
  })
}
