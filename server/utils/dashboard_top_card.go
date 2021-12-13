package utils

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTopCardData
//@description: 从Redis获取首页顶部卡片组件数据
//@return: json data, err error
func GetTopCardData() json.RawMessage {
	TopCardData, err := global.GVA_REDIS.Get(context.Background(), "TopCard").Result()
	if err != nil{
		global.GVA_LOG.Error("从Redis获取TopCard数据失败：", zap.Error(err))
	}
	TopCardDataJson, _ := json.Marshal(TopCardData)
	return json.RawMessage(TopCardDataJson)
}
