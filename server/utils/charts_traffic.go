package utils

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTrafficData
//@description: 获取最近上下行流量
//@return: json data, err error
func GetTraffic(ParamsMap map[string]gjson.Result) string {
	var (
		TopCardData string
		err error
	)
	// 不带参数的请求直接查Redis，拿1h的数据
	if ParamsMap["startTime"].String() != ""{
		if TopCardData = QueryDB(ParamsMap["startTime"].String(), ParamsMap["endTime"].String()); TopCardData == "" {
			global.GVA_LOG.Error("从Redis获取最近1h上下行流量数据失败，查询返回空结果")
		}
	} else {
		TopCardData, err = global.GVA_REDIS.Get(context.Background(), "1hTrafficData").Result()
		if err != nil{
			global.GVA_LOG.Error("从Redis获取最近1h上下行流量数据失败：", zap.Error(err))
		}
	}
	return TopCardData
}

func QueryDB(startTime string, endTime string) (string) {
	type Traffic struct {
		Time           string  `gorm:"column:Time;type:datetime;" json:"Time"`
		InTrafficMbps  float32 `json:"in_traffic_mbps"`
		OutTrafficMbps float32 `json:"out_traffic_mbps"`
	}
	var result []Traffic
	subQuery1 := global.GORM_CH.Raw("select toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 InMBytes, sum(packets) InPackets, InMBytes*0.8 InTrafficMbps from(select * from nms_data.gateway_pmacctd prewhere timestamp_min >= ? AND timestamp_max <= ?) where loc_dst = '局域网' group by Time", startTime, endTime)
	subQuery2 := global.GORM_CH.Raw("select toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 OutMBytes, sum(packets) OutPackets, OutMBytes*0.8 OutTrafficMbps from(select * from nms_data.gateway_pmacctd prewhere timestamp_min >= ? AND timestamp_max <= ?) where loc_src = '局域网' group by Time", startTime, endTime)
	global.GORM_CH.Table("(?) as In, (?) as Out", subQuery1, subQuery2).Select("In.Time,FLOOR(In.InTrafficMbps,2) in_traffic_mbps,FLOOR(Out.OutTrafficMbps,2) out_traffic_mbps").Where("In.Time = Out.Time ORDER BY Time").Find(&result)
	result2, _ := json.Marshal(result)
	return string(result2)
}