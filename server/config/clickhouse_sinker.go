package config

type ClickhouseSinker struct {

	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址
	Port	 string `mapstructure:"port" json:"port" yaml:"port"` // 端口
}
