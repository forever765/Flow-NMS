package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/charts"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	ChartsServiceGroup   charts.ServiceGroup
	DashboardServiceGroup   dashboard.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
