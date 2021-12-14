package dashboard

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTopCardData
//@description: 从Redis获取首页顶部卡片组件数据
//@return: err error, conf config.Server

type DashboardService struct {
}

func (DashboardService *DashboardService) GetTopCardData() (err error, result json.RawMessage) {
	result = utils.GetTopCardData()
	return nil, result
}
