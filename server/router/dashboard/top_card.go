package dashboard

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DashboardRouter struct {
}

func (s *DashboardRouter) InitDashboardRouter(Router *gin.RouterGroup) {
	dashboardRouter := Router.Group("dashboard").Use(middleware.OperationRecord())
	var dashboardApi = v1.ApiGroupApp.DashboardApiGroup
	{
		dashboardRouter.POST("getTopCardData", dashboardApi.GetTopCardData) // 获取上下行流量
	}
}
