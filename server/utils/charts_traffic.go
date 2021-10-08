package utils

import (
	"encoding/json"
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
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
	dsn := "tcp://192.168.123.221:9000?database=nms_data&username=&password=&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(100)
	// RAW SQL: SELECT In.Time,In.InTrafficMbps,Out.OutTrafficMbps FROM (SELECT Time,InTrafficMbps FROM `InTraffic` WHERE Time > '2021-10-06 16:05:31') as In, (SELECT Time,OutTrafficMbps FROM `OutTraffic` WHERE Time > '2021-10-06 16:05:31') as Out WHERE In.Time = Out.Time
	now := time.Now()
	h, _ := time.ParseDuration("-1h")
	anHourAgo := (now.Add(h)).Format("2006-01-02 15:04:05")
	var result []Traffic
	subQuery1 := db.Table("InTraffic").Select("Time,InTrafficMbps").Where("Time > ?", anHourAgo)
	subQuery2 := db.Table("OutTraffic").Select("Time,OutTrafficMbps").Where("Time > ?", anHourAgo)
	db.Table("(?) as In, (?) as Out", subQuery1, subQuery2).Select("Time,FLOOR(InTrafficMbps,2) as in_traffic_mbps,FLOOR(OutTrafficMbps,2) as out_traffic_mbps").Where("In.Time = Out.Time ORDER BY Time").Find(&result)
	result2, _ := json.Marshal(result)
	return json.RawMessage(result2)
}
