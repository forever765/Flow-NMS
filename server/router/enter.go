package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/router/charts"
	"github.com/flipped-aurora/gin-vue-admin/server/router/dashboard"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/reports"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system_tools"
)

type RouterGroup struct {
	System   system.RouterGroup
	SystemTools   system_tools.RouterGroup
	Charts   charts.RouterGroup
	Reports   reports.RouterGroup
	Dashboard   dashboard.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
