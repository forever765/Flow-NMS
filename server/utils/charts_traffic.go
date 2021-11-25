package utils

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"time"
)

type Traffic struct {
	Time           string  `gorm:"column:Time;type:datetime;" json:"Time"`
	InTrafficMbps  float32 `json:"in_traffic_mbps"`
	OutTrafficMbps float32 `json:"out_traffic_mbps"`
}

//@author: [forever765](https://github.com/forever765)
//@function: GetTrafficData
//@description: 获取最近上下行流量
//@return: json data, err error
func GetTraffic() json.RawMessage {
	ClickhouseCfg := global.GVA_CONFIG.Clickhouse
	dsn := fmt.Sprintf("tcp://%v?dateabase=%v&usernam=%v&password=%v&read_timeout=%v&write_timeout=%v", ClickhouseCfg.Addr,ClickhouseCfg.DB,ClickhouseCfg.Username,ClickhouseCfg.Password,ClickhouseCfg.ReadTimeout,ClickhouseCfg.WriteTimeout)
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		global.GVA_LOG.Error("Failed to connect Clickhouse!", zap.Error(err))
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(100)
	now := time.Now()
	h, _ := time.ParseDuration("-1h")
	anHourAgo := (now.Add(h)).Format("2006-01-02 15:04:05")
	var result []Traffic
	// RAW SQL: SELECT In.Time,In.InTrafficMbps,Out.OutTrafficMbps FROM (SELECT Time,InTrafficMbps FROM `InTraffic` WHERE Time > '2021-10-06 16:05:31') as In, (SELECT Time,OutTrafficMbps FROM `OutTraffic` WHERE Time > '2021-10-06 16:05:31') as Out WHERE In.Time = Out.Time
	//subQuery1 := db.Table("InTraffic").Select("Time,InTrafficMbps").Where("Time > ?", anHourAgo)
	//subQuery2 := db.Table("OutTraffic").Select("Time,OutTrafficMbps").Where("Time > ?", anHourAgo)
	//db.Table("(?) as In, (?) as Out", subQuery1, subQuery2).Select("Time,FLOOR(InTrafficMbps,2) as in_traffic_mbps,FLOOR(OutTrafficMbps,2) as out_traffic_mbps").Where("In.Time = Out.Time ORDER BY Time").Find(&result)

	//RAW SQL: select In.Time,In.InTrafficMbps,Out.OutTrafficMbps FROM(
	//select toStartOfInterval(timestamp_max, INTERVAL 10 second ) as Time, sum(bytes)/1048576 as InMBytes, sum(packets) as InPackets, InMBytes*0.8 as InTrafficMbps
	//from(select * from nms_data.pmacctd_data prewhere timestamp_max >= NOW() - 3600) where loc_dst = '局域网'
	//group by Time
	//order by Time) as In,(select toStartOfInterval(timestamp_max, INTERVAL 10 second ) as Time, sum(bytes)/1048576 as OutMBytes, sum(packets) as OutPackets, OutMBytes*0.8 as OutTrafficMbps
	//from(select * from nms_data.pmacctd_data prewhere timestamp_max >= NOW() - 3600) where loc_src = '局域网'
	//group by Time
	//order by Time) as Out WHERE In.Time=Out.Time;
	subQuery1 := db.Raw("select toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 InMBytes, sum(packets) InPackets, InMBytes*0.8 InTrafficMbps from(select * from nms_data.pmacctd_data prewhere timestamp_max >= ?) where loc_dst = '局域网' group by Time", anHourAgo)
	subQuery2 := db.Raw("select toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 OutMBytes, sum(packets) OutPackets, OutMBytes*0.8 OutTrafficMbps from(select * from nms_data.pmacctd_data prewhere timestamp_max >= ?) where loc_src = '局域网' group by Time", anHourAgo)
	db.Table("(?) as In, (?) as Out", subQuery1, subQuery2).Select("In.Time,FLOOR(In.InTrafficMbps,2) in_traffic_mbps,FLOOR(Out.OutTrafficMbps,2) out_traffic_mbps").Where("In.Time = Out.Time ORDER BY Time").Find(&result)
	result2, _ := json.Marshal(result)
	return json.RawMessage(result2)
}
