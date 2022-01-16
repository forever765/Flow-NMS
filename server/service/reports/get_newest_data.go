package reports

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/reports"
)

//@author: [forever765](https://github.com/forever765)
//@function: 获取报表最新数据
//@description: 获取报表最新数据
//@return: err error, result string

type ReportsService struct {
}


func (ReportsService *ReportsService) GetNewestData(pageNum int, pageSize int) (err error, result []map[string]interface{}) {
	result = reports.GetNewestData(pageNum, pageSize)
	return nil, result
}
