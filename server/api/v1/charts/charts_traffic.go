package charts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io/ioutil"
)

type ChartsApi struct {
}

// @Tags Charts
// @Summary 获取最近上下行流量
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /charts/getTrafficData [post]
func (s *ChartsApi) GetTraffic(c *gin.Context) {
	raw, _ := ioutil.ReadAll(c.Request.Body)
	ParamsMap := gjson.GetBytes(raw, "@this").Map()
	if err, result := chartsService.GetTraffic(ParamsMap); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}
