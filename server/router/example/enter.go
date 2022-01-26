package example

import "github.com/flipped-aurora/gin-vue-admin/server/router/system_tools"

type RouterGroup struct {
	CustomerRouter
	system_tools.ExcelRouter
	FileUploadAndDownloadRouter
}
