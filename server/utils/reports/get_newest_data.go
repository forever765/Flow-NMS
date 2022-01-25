package reports

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
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

var NewestData []map[string]interface{}
//@author: [forever765](https://github.com/forever765)
//@function: GetNewestData
//@description: 获取最新报表数据
//@return: json data, err error
func GetNewestData(ParamsMap map[string]gjson.Result) []map[string]interface{} {
	db := global.GORM_CH
	Db := db
	// 如果包含class类型
	if _,exist := ParamsMap["class"]; exist {
		Db = Db.Where("class", ParamsMap["class"].String())
	}
	// 如果包含IP地址
	if _,exist := ParamsMap["ipAddr"]; exist {
		Db = Db.Where("ip_src = ? or ip_dst = ?", ParamsMap["ipAddr"].String(), ParamsMap["ipAddr"].String())
	}
	if err := Db.
		Table("nms_data.gateway_pmacctd").
		Select("timestamp_min, timestamp_max, ip_src, port_src, isp_src, loc_src, ip_dst, isp_dst, loc_dst, port_dst, bytes, class, ip_proto, packets, etype").
		Where("timestamp_min >= NOW()-600").
		Offset(OffsetCompute(int(ParamsMap["pageNum"].Num), int(ParamsMap["pageSize"].Num))).
		Limit(int(ParamsMap["pageSize"].Num)).
		Order("timestamp_min DESC").
		Debug().
		Find(&NewestData).
		Error; err != nil {
		global.GVA_LOG.Error("获取报表最新数据失败:", zap.Error(err))
	}
	// 遍历合并字段，简化前端操作
	for i:= 0; i<len(NewestData); i++{
		if strings.Contains(fmt.Sprintf("%v",NewestData[i]), "nil") {
			fmt.Println("found nil")
			break
		}
		NewestData[i]["src_ip_port"] = fmt.Sprintf("%v:%v", NewestData[i]["ip_src"], NewestData[i]["port_src"])
		NewestData[i]["dst_ip_port"] = fmt.Sprintf("%v:%v", NewestData[i]["ip_dst"], NewestData[i]["port_dst"])
		// 优化局域网显示，局域网:局域网 => 局域网
		if NewestData[i]["loc_src"] == "局域网" {
			NewestData[i]["src_loc_isp"] = "局域网"
		} else {
			NewestData[i]["src_loc_isp"] = fmt.Sprintf("%v%v", NewestData[i]["loc_src"], NewestData[i]["isp_src"])
		}
		if NewestData[i]["loc_dst"] == "局域网" {
			NewestData[i]["dst_loc_isp"] = "局域网"
		} else {
			NewestData[i]["dst_loc_isp"] = fmt.Sprintf("%v%v", NewestData[i]["loc_dst"], NewestData[i]["isp_dst"])
		}
		// 优化ipv4/6显示，和协议合并显示
		if NewestData[i]["etype"] == "800" {
			NewestData[i]["protocol_etype"] = fmt.Sprintf("%v/v4", NewestData[i]["ip_proto"])
		} else {
			NewestData[i]["protocol_etype"] = fmt.Sprintf("%v/v6", NewestData[i]["ip_proto"])
		}
		// 处理时间
		MinTimeStep1 := strings.Replace(fmt.Sprintf("%v", NewestData[i]["timestamp_min"]) , "CST", "", -1)
		MinTimeStep2 := strings.Replace(MinTimeStep1 , "T", " ", -1)
		NewestData[i]["timestamp_min"] = strings.Replace(MinTimeStep2 , " +0800 ", "", -1)
		MaxTimeStep1 := strings.Replace(fmt.Sprintf("%v", NewestData[i]["timestamp_max"]) , "CST", "", -1)
		MaxTimeStep2 := strings.Replace(MaxTimeStep1 , "T", " ", -1)
		NewestData[i]["timestamp_max"] = strings.Replace(MaxTimeStep2 , " +0800 ", "", -1)

		delete(NewestData[i], "ip_src")
		delete(NewestData[i], "port_src")
		delete(NewestData[i], "ip_dst")
		delete(NewestData[i], "port_dst")
		delete(NewestData[i], "loc_src")
		delete(NewestData[i], "isp_src")
		delete(NewestData[i], "loc_dst")
		delete(NewestData[i], "isp_dst")
		delete(NewestData[i], "ip_proto")
		delete(NewestData[i], "etype")
	}
	defer func() {
		NewestData = nil
	}()
	return NewestData
}

// offset计算
func OffsetCompute(pageNum int, pageSize int) (offset int) {
	if pageNum == 1 {
		offset = -1
	} else {
		offset = (pageNum-1) * pageSize
	}
	return offset
}