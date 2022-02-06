package utils

import (
	"context"
	"encoding/json"
	"fmt"
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
	if ParamsMap["startTime"].String() != "" || ParamsMap["protocolVersion"].String() != "" {
		if TopCardData = QueryDB(ParamsMap); TopCardData == "" {
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

func QueryDB(ParamsMap map[string]gjson.Result) (string) {
	type Traffic struct {
		Time           string  `gorm:"column:Time;type:datetime;" json:"Time"`
		InTrafficMbps  float32 `json:"in_traffic_mbps"`
		OutTrafficMbps float32 `json:"out_traffic_mbps"`
	}
	var result []Traffic
	Db := global.GORM_CH
	var (
		startTime = "NOW() - 3600"
		endTime = "NOW()"
	)
	// 如果有指定时间范围就把查询条件改为该时间，否则保持默认值1h
	if startTimeTemp := ParamsMap["startTime"].String(); startTimeTemp!=""{
		startTime = startTimeTemp
	}
	if endTimeTemp := ParamsMap["endTime"].String(); endTimeTemp!=""{
		endTime = endTimeTemp
	}
	// 如果要区别协议版本，直接加在后面作为条件
	if value, exist := ParamsMap["protocolVersion"]; exist {
		switch value.String() {
		case "仅IPv4":
			endTime = endTime + " AND etype = '800'"
		case "仅IPv6":
			endTime = endTime + " AND etype = '86dd'"
		}
	}
	subQuery1 := Db.Raw(fmt.Sprintf("SELECT toStartOfInterval(timestamp_max, INTERVAL 10 second) Time, sum(bytes)/1048576 InMBytes, sum(packets) InPackets, InMBytes*0.8 InTrafficMbps FROM (SELECT * FROM nms_data.gateway_pmacctd PREWHERE timestamp_min >= %v AND timestamp_max <= %v) WHERE loc_dst = '局域网' GROUP BY Time", startTime, endTime))
	subQuery2 := Db.Raw(fmt.Sprintf("SELECT toStartOfInterval(timestamp_max, INTERVAL 10 second) Time, sum(bytes)/1048576 OutMBytes, sum(packets) OutPackets, OutMBytes*0.8 OutTrafficMbps FROM (SELECT * FROM nms_data.gateway_pmacctd PREWHERE timestamp_min >= %v AND timestamp_max <= %v) WHERE loc_src = '局域网' GROUP BY Time", startTime, endTime))
	global.GORM_CH.
		Table("(?) as In, (?) as Out", subQuery1, subQuery2).
		Select("In.Time,FLOOR(In.InTrafficMbps,2) in_traffic_mbps,FLOOR(Out.OutTrafficMbps,2) out_traffic_mbps").
		Where("In.Time = Out.Time").
		Order("Time").
		Debug().
		Find(&result)
	result2, _ := json.Marshal(result)
	return string(result2)
}