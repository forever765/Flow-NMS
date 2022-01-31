package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/reports"
	"github.com/tidwall/gjson"
)

//@author: [forever765](https://github.com/forever765)
//@function: 获取最新TopN数据
//@description: 获取最新TopN数据
//@return: err error, result string
func (ReportsService *ReportsService) GetTopN(ParamsMap map[string]gjson.Result) (err error, result map[string]interface{}) {
	result = reports.GetTopN(ParamsMap)
	return nil, result
}
