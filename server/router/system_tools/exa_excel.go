package system_tools

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ExcelRouter struct {
}

func (e *ExcelRouter) InitExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("iphost")
	var systemToolsApi = v1.ApiGroupApp.SystemToolsApiGroup.SystemToolsApi
	{
		excelRouter.POST("importExcel", systemToolsApi.ImportExcel)          // 导入Excel
		excelRouter.GET("loadExcel", systemToolsApi.LoadExcel)               // 加载Excel数据
		excelRouter.POST("exportExcel", systemToolsApi.ExportExcel)          // 导出Excel
		excelRouter.GET("downloadTemplate", systemToolsApi.DownloadTemplate) // 下载模板文件
	}
}
