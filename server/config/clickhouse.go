package config

type Clickhouse struct {
	Cluster	string  `mapstructure:"cluster" json:"cluster" yaml:"cluster"`    //Clickhouse Cluster
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
