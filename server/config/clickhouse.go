package config

type Clickhouse struct {
	Cluster	string  `mapstructure:"cluster" json:"cluster" yaml:"cluster"`    // Cluster Name
	DB       string    `mapstructure:"db" json:"db" yaml:"db"`                // Clickhouse DB
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址
	Port     int `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	ReadTimeout int `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"` // 读取超时时间
	WriteTimeout int `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` // 写入超时时间
	PrometheusPort int `mapstructure:"prometheus_port" json:"prometheus_port" yaml:"prometheus_port"` // Prometheus Exporter端口
	PrometheusSuffix string `mapstructure:"prometheus_suffix" json:"prometheus_suffix" yaml:"prometheus_suffix"` // Prometheus Exporter地址后缀
}
