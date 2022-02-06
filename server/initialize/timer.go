package initialize

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"time"
)

func Timer() {
	//if global.GVA_CONFIG.Timer.Start {
	//	for i := range global.GVA_CONFIG.Timer.Detail {
	//		go func(detail config.Detail) {
	//			global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
	//				err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
	//				if err != nil {
	//					fmt.Println("timer error:", err)
	//				}
	//			})
	//		}(global.GVA_CONFIG.Timer.Detail[i])
	//	}

	if global.GVA_CONFIG.Timer.Start {
		t := global.GVA_Timer

		// 任务名称：IndexDashboardTopCard
		// 任务用途：首页 Dashboard 顶栏数据缓存，to Redis
		// 执行间隔：Every minute
		_, err := t.AddTaskByFunc("IndexDashboardTopCard", "*/1 * * * *", func() {
			var (
				todayTraffic string
				MonthTraffic string
				HistoryTraffic	string
				ChTotalCount string
			)
			db := global.GORM_CH
			// golang做四舍五入很麻烦，故而让clickhouse代为操作
			if err := db.Table("nms_data.gateway_pmacctd").Select("FLOOR(SUM(bytes)/1073741824,2) as GBytes").Where("timestamp_min > toStartOfDay(NOW())").Find(&todayTraffic).Error; err != nil {
				global.GVA_LOG.Error("当天流量查询失败:", zap.Error(err))
			}
			if err := db.Table("nms_data.gateway_pmacctd").Select("FLOOR(SUM(bytes)/1073741824,2) as GBytes").Where("timestamp_min > toStartOfMonth(NOW())").Find(&MonthTraffic).Error; err != nil {
				global.GVA_LOG.Error("当月流量查询失败:", zap.Error(err))
			}
			if err := db.Table("nms_data.gateway_pmacctd").Select("FLOOR(SUM(bytes)/1073741824,2) as GBytes").Find(&HistoryTraffic).Error; err != nil {
				global.GVA_LOG.Error("历史流量查询失败:", zap.Error(err))
			}
			if err := db.Table("nms_data.gateway_pmacctd").Select("COUNT(timestamp_min)").Find(&ChTotalCount).Error; err != nil {
				global.GVA_LOG.Error("CH记录总数查询失败:", zap.Error(err))
			}
			// 写入 Redis
			timer := time.Duration(3600) * time.Second
			err := global.GVA_REDIS.Set(context.Background(), "TopCard", todayTraffic+","+MonthTraffic+","+HistoryTraffic+","+ChTotalCount, timer).Err()
			if err != nil {
				global.GVA_LOG.Error("写入Redis失败:", zap.Error(err))
			} else {
				//global.GVA_LOG.Info("定时任务 IndexDashboardTopCard 执行完成")
			}
		})
		if err != nil {
			//global.GVA_LOG.Error("定时任务 IndexDashboardTopCard添加失败：", zap.Error(err))
		}



		// 任务名称：IndexDashboardTraffic
		// 任务用途：首页 Dashboard 流量图数据缓存，to Redis
		// 执行间隔：Every 5s
		type Traffic struct {
			Time           string  `gorm:"column:Time;type:datetime;" json:"Time"`
			InTrafficMbps  float32 `json:"in_traffic_mbps"`
			OutTrafficMbps float32 `json:"out_traffic_mbps"`
		}
		_, err2 := t.AddSecondTaskByFunc("IndexDashboardTraffic", "@every 5s", func() {
			now := time.Now()
			h, _ := time.ParseDuration("-1h")
			anHourAgo := (now.Add(h)).Format("2006-01-02 15:04:05")
			var result []Traffic
			Db := global.GORM_CH
			subQuery1 := Db.Raw("SELECT toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 InMBytes, sum(packets) InPackets, InMBytes*0.8 InTrafficMbps FROM(SELECT * FROM nms_data.gateway_pmacctd PREWHERE timestamp_max >= ?) WHERE loc_dst = '局域网' GROUP BY Time", anHourAgo)
			subQuery2 := Db.Raw("SELECT toStartOfInterval(timestamp_max, INTERVAL 10 second ) Time, sum(bytes)/1048576 OutMBytes, sum(packets) OutPackets, OutMBytes*0.8 OutTrafficMbps FROM(SELECT * FROM nms_data.gateway_pmacctd PREWHERE timestamp_max >= ?) WHERE loc_src = '局域网' GROUP BY Time", anHourAgo)
			global.GORM_CH.
				Table("(?) as In, (?) as Out", subQuery1, subQuery2).
				Select("In.Time,FLOOR(In.InTrafficMbps,2) in_traffic_mbps,FLOOR(Out.OutTrafficMbps,2) out_traffic_mbps").
				Where("In.Time = Out.Time").
				Order("Time").
				//Debug().
				Find(&result)
			result2, _ := json.Marshal(result)

			// 写入 Redis
			timer := time.Duration(3600) * time.Second
			err := global.GVA_REDIS.Set(context.Background(), "1hTrafficData", string(result2), timer).Err()
			if err != nil {
				global.GVA_LOG.Error("写入Redis失败:", zap.Error(err))
			}else {
				//global.GVA_LOG.Info("定时任务 IndexDashboardTraffic 执行完成")
			}
		})
		if err2 != nil {
			global.GVA_LOG.Error("定时任务 IndexDashboardTraffic 添加失败：", zap.Error(err))
		}
	}
}
