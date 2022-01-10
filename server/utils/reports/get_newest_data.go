package reports

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"strings"
)

type dbTable struct {
	Timestamp_min string
	Timestamp_max string
	Ip_src string
	Port_src uint16
	Ip_dst string
	Port_dst uint16
	Bytes int64
	Class string
	Ip_proto string
	Packets	int64
	Loc_src string
	Loc_dst string
	Isp_src string
	Isp_dst string
	Etype string
}
var NewestData []dbTable
//@author: [forever765](https://github.com/forever765)
//@function: GetNewestData
//@description: 获取最新报表数据
//@return: json data, err error
func GetNewestData() string {
//func GetNewestData() map[string]interface{} {
	db := global.GORM_CH
	if err := db.Table("nms_data.gateway_pmacctd").Select("*").Where("timestamp_min >= NOW()-300 ORDER BY timestamp_min DESC LIMIT 200").Find(&NewestData).Error; err != nil {
		global.GVA_LOG.Error("获取报表最新数据失败:", zap.Error(err))
	}
	return strings.Replace(strings.Trim(fmt.Sprint(NewestData), "[]"), " ", ",", -1)
}
