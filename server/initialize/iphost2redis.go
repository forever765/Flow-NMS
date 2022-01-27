package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

func IpHostInfo() {
	var systemToolsService = service.ServiceGroupApp.SystemToolsServiceGroup.SystemToolsService
	err := systemToolsService.ParseExcel2Redis()
	if err != nil {
		global.GVA_LOG.Error("从Excel读取数据到Redis失败!", zap.Any("err", err))
	}
}
