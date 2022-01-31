package reports

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTopN
//@description: 获取最新TopN排行榜数据
//@return: json data, err error
func GetTopN(ParamsMap map[string]gjson.Result) map[string]interface{} {
	db := global.GORM_CH
	Db := db
	// 如果包含协议版本
	if value, exist := ParamsMap["protocolVersion"]; exist {
		switch value.String() {
		case "仅IPv4":
			Db = Db.Where("etype", "800")
		case "仅IPv6":
			Db = Db.Where("etype", "86dd")
		}
	}
	// 如果包含IP地址
	//if _,exist := ParamsMap["ipAddr"]; exist {
	//	Db = Db.Where("ip_src = ? or ip_dst = ?", ParamsMap["ipAddr"].String(), ParamsMap["ipAddr"].String())
	//}
	// 如果包含时间范围
	if _, exist := ParamsMap["startTime"]; exist {
		Db = Db.Where("timestamp_min >= ? and timestamp_max <= ?", ParamsMap["startTime"].String(), ParamsMap["endTime"].String())
	} else {
		Db = Db.Where("timestamp_min >= NOW()-3600")
	}

	// src和dst各执行一次查询
	var tmpResult []map[string]interface{}
	finalResult := map[string]interface{} {
		"src": nil,
		"dst": nil,
	}
	types := [2]string{"src", "dst"}
	for _, obj := range types {
		if err := Db.
			Table("nms_data.gateway_pmacctd").
			// 此处为固定输入src/dst，无需考虑防注入，所以使用Sprintf
			Select(fmt.Sprintf("%v, any(%v) %v, any(%v) %v, sum(bytes)/1048576 MBytes, sum(packets) Packets", "ip_"+obj, "loc_"+obj, "loc_"+obj, "isp_"+obj, "isp_"+obj)).
			Offset(OffsetCompute(int(ParamsMap["pageNum"].Num), int(ParamsMap["pageSize"].Num))).
			Limit(int(ParamsMap["pageSize"].Num)).
			Group("ip_" + obj).
			Order("MBytes DESC").
			Debug().
			Find(&tmpResult).
			Error; err != nil {
			global.GVA_LOG.Error("获取TopN最新数据失败:", zap.Error(err))
		}
		finalResult[obj] = tmpResult
	}
	return finalResult
}