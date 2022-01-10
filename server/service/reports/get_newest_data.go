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

func (ReportsService *ReportsService) GetNewestData() (err error, result string) {
	result = reports.GetNewestData()
	return nil, result
}
