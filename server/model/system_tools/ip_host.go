package system_tools

type IpHost struct {
	//ID        	  uint								 `json:"id"`     // ID
	Area          string                             `json:"area"`     // 地区
	Type          string                             `json:"type"`     // 类型
	HostName      string                             `json:"hostname"` // 主机名
	IpAddr        string                             `json:"ipaddr"`   // IP地址
}