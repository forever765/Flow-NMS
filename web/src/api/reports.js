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
// @Router /report/getNewestData [get]
// {
//  page     int
//	pageSize int
// }
export const getNewestData = (data) => {
  return service({
    url: '/reports/getNewestData?pageNum=' + data.pageNum + '&pageSize=' + data.pageSize,
    method: 'get',
  })
}

// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
// export const deleteApisByIds = (data) => {
//   return service({
//     url: '/api/deleteApisByIds',
//     method: 'delete',
//     data
//   })
// }
