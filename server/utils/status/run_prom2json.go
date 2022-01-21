package status

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunProm2Json() string {
	// Todo：对接配置页面
	var Bin string
	Prom2jsonPath := filepath.Join("D:/script/prom2json-1.3.0")
	switch runtime.GOOS {
	case "windows":
		Bin = "prom2json.exe"
	default:
		Bin = "prom2json"
	}
	ExePath := filepath.Join(Prom2jsonPath, Bin)
	var DownResult bool = true
	_, err := os.Stat(ExePath)
	if err != nil {
		if os.IsNotExist(err) {
			global.GVA_LOG.Error("Prom2json不存在，尝试从gitee下载...")
				DownResult = DownloadProm2json()
		}
		fmt.Println("Prom2Json 文件已存在，继续")
	}
	// Todo: 对接配置页面
	if DownResult == true {
		cmd := exec.Command(ExePath, "http://nms01.home.com:21888/metrics")
		output, err := cmd.CombinedOutput()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		result := ParseResult(string(output))
		return result
	}
	return ""
}

func ParseResult(result string) string {
	if !gjson.Valid(result) {
		global.GVA_LOG.Error("无效json，请检查 metrics源")
	}
	value := gjson.Get(result, "@this")
	//value2 := gjson.GetMany(result, "@this.@flatten.#.name", "@this.@flatten.#.metrics.@flatten.#.labels", "@this.@flatten.#.metrics.@flatten.#.value")
	//value := gjson.GetMany(result, "@this.@flatten.#.name", "@this.@flatten.#.metrics.@flatten.#.labels", "@this.@flatten.#.metrics.@flatten.#.value")
	//name := "@this.@flatten.#.name"
	//labels := "@this.@flatten.#.metrics.@flatten.#.labels"
	//value := "@this.@flatten.#.metrics.@flatten.#.value"
	//value2 := GetManyToMap(result, name, labels, value)
	//fmt.Println("aaa", value2)
	return value.String()
}
