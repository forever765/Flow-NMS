package status

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tidwall/gjson"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

func RunProm2Json() bool {
		chSinkerNaliCfg := global.GVA_CONFIG.Clickhouse_SinkerNali
		cmd := exec.Command(ExePath, chSinkerNaliCfg.Addr+":"+strconv.Itoa(chSinkerNaliCfg.Port)+"/metrics")
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
	//fmt.Println("aaa", value2)
	return value.String()
}

func CheckProm2Json() {
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
		DownloadProm2json()
	} else {
		global.GVA_LOG.Info("Prom2json文件存在")
	}
}