package config

type ClickhouseSinkerNali struct {

	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"` // 服务器地址
	Port	 string `mapstructure:"port" json:"port" yaml:"port"` // 端口
	Prom2json_path	 string `mapstructure:"prom2json_path" json:"prom2json_path" yaml:"prom2json_path"` // Prom2Json文件保存路径
}
