package system_tools

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SystemToolsApi
}

var systemToolsService = service.ServiceGroupApp.SystemToolsServiceGroup.SystemToolsService
