package utils

import (
	"encoding/json"
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type InTraffic struct {
	Time          string `gorm:"column:Time;type:datetime;" json:"Time"`
	InTrafficMbps string `json:"in_traffic_mbps"`
	//OutTrafficMbps string `json:"out_traffic_mbps"`
}

//type Os struct {
//	GOOS         string `json:"goos"`
//	NumCPU       int    `json:"numCpu"`
//	Compiler     string `json:"compiler"`
//	GoVersion    string `json:"goVersion"`
//	NumGoroutine int    `json:"numGoroutine"`
//}

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
	// select * from InTraffic where -1h
	now := time.Now()
	h, _ := time.ParseDuration("-1h")
	anHourAgo := (now.Add(h)).Format("2006-01-02 15:04:05")
	var result []InTraffic
	db.Table("InTraffic").Where("Time > ?", anHourAgo).Find(&result)
	result2, _ := json.Marshal(result)
	//json.RawMessage(result2)
	//n := len(result2) //Find the length of the byte array
	//s := string(result2[:n])
	return json.RawMessage(result2)
}
