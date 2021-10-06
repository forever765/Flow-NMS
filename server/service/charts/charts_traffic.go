package charts

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTraffic
//@description: 读取最近上下行流量
//@return: err error, conf config.Server

type ChartsService struct {
}

func (ChartsService *ChartsService) GetTraffic() (err error, result json.RawMessage) {
	result = utils.GetTraffic()
	return nil, result
}
