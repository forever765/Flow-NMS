package utils

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTrafficData
//@description: 获取最近上下行流量
//@return: json data, err error
func GetTraffic() string {
	TopCardData, err := global.GVA_REDIS.Get(context.Background(), "IndexTraffic").Result()
	if err != nil{
		global.GVA_LOG.Error("从Redis获取TopCard数据失败：", zap.Error(err))
	}
	return TopCardData
}
