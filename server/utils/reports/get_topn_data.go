package reports

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net"
	"strings"
)

//@author: [forever765](https://github.com/forever765)
//@function: GetTopN
//@description: 获取最新TopN排行榜数据
//@return: json data, err error
func GetTopN(ParamsMap map[string]gjson.Result) map[string]interface{} {
	// src和dst各执行一次查询，放到一个map内
	var tmpResult []map[string]interface{}
	finalResult := map[string]interface{} {
		"src": nil,
		"dst": nil,
	}
	types := [2]string{"src", "dst"}
	for _, obj := range types {
		Db := defineDB(ParamsMap)
		if err := Db.
			Table("nms_data.gateway_pmacctd").
			// 此处为固定输入src/dst，无需考虑防注入，所以使用Sprintf
			Select(fmt.Sprintf("%v ipaddr, any(%v) location, any(%v) isp, sum(bytes) mbytes, sum(packets) packets", "ip_"+obj, "loc_"+obj, "isp_"+obj)).
			Offset(OffsetCompute(int(ParamsMap["pageNum"].Num), int(ParamsMap["pageSize"].Num))).
			Limit(int(ParamsMap["pageSize"].Num)).
			Group("ip_" + obj).
			Order("mbytes DESC").
			Debug().
			Find(&tmpResult).
			Error; err != nil {
			global.GVA_LOG.Error("获取TopN最新数据失败:", zap.Error(err))
		}
		// 遍历合并字段，简化前端操作
		for i:= 0; i<len(tmpResult); i++ {
			if strings.Contains(fmt.Sprintf("%v", tmpResult[i]), "nil") {
				fmt.Println("found nil")
				break
			}
			// 局域网精简显示
			if tmpResult[i]["isp"] == "局域网" {
				tmpResult[i]["isp"] = "局域网"
			} else {
				// 拼接 loc-isp
				tmpResult[i]["isp"] = fmt.Sprintf("%v%v", tmpResult[i]["location"], tmpResult[i]["isp"])
			}
			delete(tmpResult[i], "location")
		}
		// 在循环外执行，只需要读两次redis
		tmpResult = searchHostnameAndSubnet(tmpResult)
		finalResult[obj] = tmpResult
		tmpResult = nil
	}
	return finalResult
}

func searchHostnameAndSubnet(tmpResult []map[string]interface{}) []map[string]interface{} {
	redisResult, err := global.GVA_REDIS.Get(context.Background(), "IpHostList").Result()
	if err != nil {
		global.GVA_LOG.Error("从Redis获取IpHostList失败：", zap.Error(err))
	}

	// 遍历处理所有ip，提取要匹配的ip
	for i:= 0; i<len(tmpResult); i++ {
		// 默认值
		tmpResult[i]["hostname"] = tmpResult[i]["ipaddr"]
		ipaddr := fmt.Sprintf("%v", tmpResult[i]["ipaddr"])
		// redisResultValue example: [{"id":1,"type":"内网","area":"顺德","hostname":"主路由","ipaddr":"192.168.123.1"},{"id":2,"area":"顺德","type":"内网","hostname":"核心交换机","ipaddr":"192.168.123.2"}]
		// 先用ip-host表匹配，无结果的话再去dns反查
		if redisResultValue := gjson.Get(redisResult, "@values"); redisResultValue.String() != "" {
			for _, record := range redisResultValue.Array() {
				obj := record.Map()
				if (obj["ipaddr"]).String() == ipaddr {
					tmpResult[i]["hostname"] = obj["hostname"].String()
					tmpResult[i]["isp"] = obj["area"].String()+"-"+obj["type"].String()
				}
			}
		} else {
			iprecords, _ := net.LookupIP(ipaddr)
			tmpResult[i]["hostname"] = iprecords[0]
		}
	}
	return tmpResult
}

func defineDB(ParamsMap map[string]gjson.Result) *gorm.DB {
	Db := global.GORM_CH
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
	return Db
}