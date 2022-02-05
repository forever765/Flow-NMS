package system_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemToolsApi struct {
}

// /iphost/importExcel 接口，与upload接口作用类似，只是把文件存到resource/excel目录下，用于导入Excel时存放Excel文件(ExcelImport.xlsx)
// /iphost/loadExcel接口，用于读取resource/excel目录下的文件((ExcelImport.xlsx)并加载为[]model.SysBaseMenu类型的示例数据
// /iphost/exportExcel 接口，用于读取前端传来的tableData，生成Excel文件并返回
// /iphost/downloadTemplate 接口，用于下载resource/excel目录下的 ExcelTemplate.xlsx 文件，作为导入的模板

// @Tags iphost
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body example.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /iphost/exportExcel [post]
func (e *SystemToolsApi) ExportExcel(c *gin.Context) {
	var excelInfo system_tools.ExcelInfo
	_ = c.ShouldBindJSON(&excelInfo)
	filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
	err := systemToolsService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Any("err", err))
		response.FailWithMessage("转换Excel失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

// @Tags iphost
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /iphost/importExcel [post]
func (e *SystemToolsApi) ImportExcel(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	_ = c.SaveUploadedFile(header, global.GVA_CONFIG.Excel.Dir+"ExcelImport.xlsx")
	if err := systemToolsService.ParseExcel2Redis(); err != nil {
		response.FailWithMessage("导入失败："+err.Error(), c)
	} else {
		response.OkWithMessage("导入成功，已加载到缓存", c)
	}
}

// @Tags iphost
// @Summary 从Redis加载Excel数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"加载数据成功"}"
// @Router /iphost/loadExcel [get]
func (e *SystemToolsApi) LoadExcel(c *gin.Context) {
	menus, err := systemToolsService.ParseInfoList4Redis()
	if err != nil {
		global.GVA_LOG.Error("加载数据失败!", zap.Any("err", err))
		response.FailWithMessage("加载数据失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menus,
		Total:    int64(len(menus)),
		Page:     1,
		PageSize: 999,
	}, "加载数据成功", c)
}

// @Tags iphost
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "模板名称"
// @Success 200
// @Router /iphost/downloadTemplate [get]
func (e *SystemToolsApi) DownloadTemplate(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := global.GVA_CONFIG.Excel.Dir + fileName
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		global.GVA_LOG.Error("文件不存在!", zap.Any("err", err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}