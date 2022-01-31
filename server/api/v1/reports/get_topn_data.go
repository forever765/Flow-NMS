package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io/ioutil"
)

// @Tags Reports
// @Summary 获取最新TopN数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reports/getTopN [post]
func (s *ReportsApi) GetTopN(c *gin.Context) {
	raw, _ := ioutil.ReadAll(c.Request.Body)
	ParamsMap := gjson.GetBytes(raw, "@this").Map()
	if err, result := reportsService.GetTopN(ParamsMap) ; err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}
