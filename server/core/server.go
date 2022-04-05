package core

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//if global.GVA_CONFIG.System.UseMultipoint {
	// 修改：因为NMS多处需要使用Redis缓存，所以无论是否启用【多点登录拦截】，都初始化redis服务
	// 因IpHost 需要使用，调整到main.go提前初始化Redis
	//initialize.Redis()

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf("欢迎使用Flow-NMS！当前监听地址：%v", address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
