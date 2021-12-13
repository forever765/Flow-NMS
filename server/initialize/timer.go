package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
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
		_, err := t.AddTaskByFunc("testFunc", "*/1 * * * *", func() {
			global.GVA_LOG.Info("nowTime write to Redis")
			//timer := time.Duration(3600) * time.Second
			//err := global.GVA_REDIS.Set(context.Background(), "nowTime", time.Now().Format("2006-01-02 15:04:05"), timer).Err()
			//if err != nil {
			//	global.GVA_LOG.Error("timer error:", zap.Error(err))
			//}
		})
		if err != nil {
			global.GVA_LOG.Error("No!!!", zap.Error(err))
		}
	}
}
