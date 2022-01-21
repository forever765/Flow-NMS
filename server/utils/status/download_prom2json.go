package status

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/walle/targz"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func DownloadProm2json() (Status bool){
	// Todo：对接配置文件
	Prom2jsonPath := ""
	if Prom2jsonPath == "" {
		Prom2jsonPath = filepath.Join("D:/script/prom2json-1.3.0")
	}
	if _, err := os.Stat(Prom2jsonPath); os.IsNotExist(err) {
		if err := os.MkdirAll(Prom2jsonPath, 0777); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("can't create %v", Prom2jsonPath))
		}
	}

	// 下载 prom2json v1.3.0
	var (
		DownloadUrl string
		FileName string
		Bin	string
	)
	// = "https://github.com/prometheus/prom2json/releases/download/v1.3.0/prom2json-1.3.0.linux-amd64.tar.gz"
	switch runtime.GOOS {
	case "windows":
		FileName = "prom2json-1.3.0-win.tar.gz"
		DownloadUrl = "https://gitee.com/forever0765/prom2json-v1.3.0/attach_files/949395/download/"+FileName
		Bin = "prom2json.exe"
	default:
		global.GVA_LOG.Info("Get runtime.GOOS failed, use linux url")
		FileName = "prom2json-1.3.0-linux.tar.gz"
		DownloadUrl = "https://gitee.com/forever0765/prom2json-v1.3.0/attach_files/949396/download/"+FileName
		Bin = "prom2json"
	}
	global.GVA_LOG.Info(fmt.Sprintf("Starting download %v", DownloadUrl))
	resp, err := http.Get(DownloadUrl)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	FullPath := filepath.Join(Prom2jsonPath, FileName)

	if err = ioutil.WriteFile(FullPath, body, 0644); err == nil {
		global.GVA_LOG.Info("Prom2json 压缩包下载完毕，", zap.String("本地路径: ", FullPath))
	} else {
		global.GVA_LOG.Error(err.Error())
	}
	defer resp.Body.Close()

	// 解压文件
	ExePath := filepath.Join(Prom2jsonPath, Bin)
	ExtraErr := targz.Extract(FullPath, Prom2jsonPath)
	if ExtraErr != nil {
		global.GVA_LOG.Error(ExtraErr.Error())
		return false
	} else {
		global.GVA_LOG.Info(fmt.Sprintf("文件解压完成，路径：%v", ExePath))
		return true
	}
}
