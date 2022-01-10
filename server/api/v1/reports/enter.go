package reports

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ReportsApi
}

var reportsService = service.ServiceGroupApp.ReportsServiceGroup
