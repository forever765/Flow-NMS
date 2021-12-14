package dashboard

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DashboardApi
}

var dashboardService = service.ServiceGroupApp.DashboardServiceGroup
