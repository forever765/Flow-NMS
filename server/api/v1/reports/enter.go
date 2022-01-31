package reports

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ReportsApi
}

type ReportsApi struct {
}

var reportsService = service.ServiceGroupApp.ReportsServiceGroup
