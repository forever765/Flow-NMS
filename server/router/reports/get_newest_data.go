package reports

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReportsRouter struct {
}

func (s *ReportsRouter) InitReportsRouter(Router *gin.RouterGroup) {
	reportsRouter := Router.Group("reports").Use(middleware.OperationRecord())
	var reportsApi = v1.ApiGroupApp.ReportsApiGroup.ReportsApi
	{
		reportsRouter.POST("getNewestData", reportsApi.GetNewestData) // 获取报表最新数据
	}
}
