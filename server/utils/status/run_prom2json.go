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

func RunProm2Json(MetricLink string) string {
	cmd := exec.Command(global.Prom2JsonBin, MetricLink)
	fmt.Println(global.Prom2JsonBin,MetricLink)
	output, err := cmd.CombinedOutput()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	result := ParseResult(string(output))
	return result
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
	//fmt.Println("aaa", result, value.String())
	return value.String()
}

func CheckProm2Json() string {
	Prom2jsonPath := filepath.Join(global.GVA_CONFIG.DownloadPath.Prom2Json)
	var Bin string
	switch runtime.GOOS {
	case "windows":
		Bin = "prom2json.exe"
	default:
		Bin = "prom2json"
	}
	ExePath := filepath.Join(Prom2jsonPath, Bin)
	_, err := os.Stat(ExePath)
	if err != nil {
		global.GVA_LOG.Error("Prom2json不存在，尝试从gitee下载...")
		if DownloadProm2json() {
		}
	} else {
		global.GVA_LOG.Info("Prom2json文件存在")
	}
	return ExePath
}