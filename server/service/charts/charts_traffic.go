package charts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/tidwall/gjson"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTraffic
//@description: 读取最近上下行流量
//@return: err error, conf config.Server

type ChartsService struct {
}

func (ChartsService *ChartsService) GetTraffic(ParamsMap map[string]gjson.Result) (err error, result string) {
	result = utils.GetTraffic(ParamsMap)
	return nil, result
}
