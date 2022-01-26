package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/charts"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/reports"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system_tools"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	SystemToolsServiceGroup   system_tools.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	ChartsServiceGroup   charts.ServiceGroup
	ReportsServiceGroup   reports.ServiceGroup
	DashboardServiceGroup   dashboard.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
