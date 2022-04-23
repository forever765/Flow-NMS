package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/status"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type SystemConfigService struct {
}

func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}


//@author: [forever765](https://github.com/forever765)
//@function: GetClickhouseInfo
//@description: 获取Clickhouse信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetClickhouseInfo() string {
	clickhouseCfg := global.GVA_CONFIG.Clickhouse
	MetricLink := fmt.Sprintf("http://%v:%v/%v", clickhouseCfg.Addr, strconv.Itoa(clickhouseCfg.PrometheusPort), clickhouseCfg.PrometheusSuffix)
	return status.RunProm2Json(MetricLink)
}


//@author: [forever765](https://github.com/forever765)
//@function: GetChSinkerNaliInfo
//@description: 获取ClickhouseSinkerNali信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetChSinkerNaliInfo() string {
	chSinkerNaliCfg := global.GVA_CONFIG.Clickhouse_SinkerNali
	MetricLink := fmt.Sprintf("%v:%v/metrics", chSinkerNaliCfg.Addr, strconv.Itoa(chSinkerNaliCfg.Port))
	return status.RunProm2Json(MetricLink)
}


//@author: [forever765](https://github.com/forever765)
//@function: GetChSinkerNaliInfo
//@description: 获取ClickhouseSinkerNali信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetKafkaAndZkInfo() map[string]map[string]string {
	kafkaAndZkCfg := global.GVA_CONFIG.KafkaAndZk
	kafkaNodeList := strings.Split(kafkaAndZkCfg.KafkaAddr, ",")
	zkNodeList := strings.Split(kafkaAndZkCfg.ZkAddr, ",")
	result := make(map[string]map[string]string)
	kafkaResult := make(map[string]string)
	zkResult := make(map[string]string)
	for _, knode := range kafkaNodeList {
		kafkaClusterSync := getRequest(fmt.Sprintf("http://%v:8000/", knode), "kafka")
		if kafkaClusterSync == "ok" {
			kafkaResult[knode] = fmt.Sprintf("%v", kafkaClusterSync)
		} else {
			kafkaResult[knode] = fmt.Sprintf("%v", kafkaClusterSync)
		}
	}
	result["kafka"] = kafkaResult
	for _, zknode := range zkNodeList {
		zkResp := getRequest("http://"+zknode +":12181", "zk")
		if zkResp == "ok" {
			zkResult[zknode] = fmt.Sprintf("%v", zkResp)
		} else {
			zkResult[zknode] = fmt.Sprintf("%v", zkResp)
		}
	}
	result["zookeeper"] = zkResult
	return result
}

func getRequest(url string, svc_type string) string{
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var value string
	if svc_type == "kafka" {
		if gjson.GetBytes(result, "status").String() == "sync"{
			value = "success"
		} else {
			value = "error"
		}
	} else {
		if gjson.GetBytes(result, "healthy").String() == "true" {
			value = "success"
		} else {
			value = "error"
		}
	}
	return value
}