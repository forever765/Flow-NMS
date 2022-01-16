package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type ReportsApi struct {
}

// @Tags Reports
// @Summary 获取最新报表数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reports/getNewestData [get]
func (s *ReportsApi) GetNewestData(c *gin.Context) {
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	pageNumInt := ConvNum(pageNum)
	pageSizeInt := ConvNum(pageSize)
	if err, result := reportsService.GetNewestData(pageNumInt, pageSizeInt); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}

func ConvNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		global.GVA_LOG.Error("请求参数无效：", zap.Error(err))
	}
	return i
}