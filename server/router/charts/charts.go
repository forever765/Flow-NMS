package charts

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChartsRouter struct {
}

func (s *ChartsRouter) InitChartsRouter(Router *gin.RouterGroup) {
	chartsRouter := Router.Group("charts").Use(middleware.OperationRecord())
	var chartsApi = v1.ApiGroupApp.ChartsApiGroup.ChartsApi
	{
		chartsRouter.POST("getTraffic", chartsApi.GetTraffic) // 获取上下行流量
		//sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig)     // 设置配置文件内容
		//sysRouter.POST("getServerInfo", systemApi.GetServerInfo)         // 获取服务器信息
		//sysRouter.POST("getClickhouseInfo", systemApi.GetClickhouseInfo) // 获取Clickhouse信息
		//sysRouter.POST("getChSinkerInfo", systemApi.GetChSinkerInfo)     // 获取Clickhouse_Sinker信息
		//sysRouter.POST("reloadSystem", systemApi.ReloadSystem)           // 重启服务
	}
}
