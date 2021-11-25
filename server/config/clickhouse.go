package config

type Clickhouse struct {
	Cluster	string  `mapstructure:"cluster" json:"cluster" yaml:"cluster"`    //Clickhouse Cluster
	DB       string    `mapstructure:"db" json:"db" yaml:"db"`                   // Clickhouse的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	ReadTimeout string `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"` // 密码
	WriteTimeout string `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` // 密码
}
