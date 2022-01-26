package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/charts"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/reports"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system_tools"
)

type ApiGroup struct {
	ExampleApiGroup  example.ApiGroup
	SystemApiGroup   system.ApiGroup
	SystemToolsApiGroup   system_tools.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	ChartsApiGroup   charts.ApiGroup
	ReportsApiGroup   reports.ApiGroup
	DashboardApiGroup   dashboard.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
