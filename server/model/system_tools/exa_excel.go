package system_tools

type ExcelInfo struct {
	FileName string   `json:"fileName"` // 文件名
	InfoList []IpHost `json:"infoList"`
}
