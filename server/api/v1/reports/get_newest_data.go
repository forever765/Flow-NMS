package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

type ReportsApi struct {
}

// @Tags Reports
// @Summary 获取最新报表数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reports/getNewestData [post]
func (s *ReportsApi) GetNewestData(c *gin.Context) {
	raw, _ := ioutil.ReadAll(c.Request.Body)
	ParamsMap := gjson.GetBytes(raw, "@this").Map()
	//fmt.Printf("解码：%v", ParamsMap)
	if err, result := reportsService.GetNewestData(ParamsMap) ; err != nil {
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