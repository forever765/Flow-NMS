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
		chartsRouter.POST("getTrafficData", chartsApi.GetTraffic) // 获取上下行流量
	}
}
