package charts

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ChartsApi
}

var chartsService = service.ServiceGroupApp.ChartsServiceGroup
