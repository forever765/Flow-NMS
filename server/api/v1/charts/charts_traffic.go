package charts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChartsApi struct {
}

// @Tags Charts
// @Summary 获取最近上下行流量
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /charts/getTraffic [post]
func (s *ChartsApi) GetTraffic(c *gin.Context) {
	if err, result := chartsService.GetTraffic(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}
