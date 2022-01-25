package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/reports"
	"github.com/tidwall/gjson"
)

//@author: [forever765](https://github.com/forever765)
//@function: 获取报表最新数据
//@description: 获取报表最新数据
//@return: err error, result string

type ReportsService struct {
}


func (ReportsService *ReportsService) GetNewestData(ParamsMap map[string]gjson.Result) (err error, result []map[string]interface{}) {
	result = reports.GetNewestData(ParamsMap)
	return nil, result
}
