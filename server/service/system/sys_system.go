package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/status"
	"go.uber.org/zap"
	"strconv"
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