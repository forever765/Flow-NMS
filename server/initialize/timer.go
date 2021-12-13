package initialize

import (
	"context"
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

	//	测试：读取clickhouse数据写入redis
	if global.GVA_CONFIG.Timer.Start {
		t := global.GVA_Timer
		// 首页Dashboard顶栏数据缓存
		_, err := t.AddTaskByFunc("IndexDashboardTopCard", "*/1 * * * *", func() {
			var (
				todayTraffic string
				MonthTraffic string
				ChTotalCount string
			)
			db := global.GORM_CH
			if err := db.Table("nms_data.gateway_pmacctd").Select("SUM(bytes)/1073741824 as GBytes").Where("timestamp_min > toStartOfDay(NOW())").Find(&todayTraffic).Error; err != nil {
				global.GVA_LOG.Error("当天流量查询失败:", zap.Error(err))
			}
			if err := db.Table("nms_data.gateway_pmacctd").Select("SUM(bytes)/1073741824 as GBytes").Where("timestamp_min > toStartOfMonth(NOW())").Find(&MonthTraffic).Error; err != nil {
				global.GVA_LOG.Error("当月流量查询失败:", zap.Error(err))
			}
			if err := db.Table("nms_data.gateway_pmacctd").Select("COUNT(timestamp_min)").Find(&ChTotalCount).Error; err != nil {
				global.GVA_LOG.Error("CH记录总数查询失败:", zap.Error(err))
			}
			// 写入 Redis
			timer := time.Duration(3600) * time.Second
			err := global.GVA_REDIS.Set(context.Background(), "TopCard", todayTraffic[:5]+","+MonthTraffic[:5]+","+ChTotalCount, timer).Err()
			if err != nil {
				global.GVA_LOG.Error("写入Redis失败:", zap.Error(err))
			}
		})
		if err != nil {
			global.GVA_LOG.Error("定时任务添加失败：", zap.Error(err))
		}
	}
}
