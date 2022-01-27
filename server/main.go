package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/status"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	global.GORM_CH = initialize.GormCH() // gorm连接Clickhouse数据库
	global.Prom2JsonBin = status.CheckProm2Json()	// 检查Prom2Json是否存在
	initialize.Timer()	// 加载定时任务
	initialize.Redis() // 因IpHostInfo使用需要，提前初始化Redis
	initialize.IpHostInfo() // 从Excel文件读取IP-Host数据到Redis
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		dbCH, _ := global.GORM_CH.DB()
		defer db.Close()
		defer dbCH.Close()
	}
	core.RunWindowsServer()
}
