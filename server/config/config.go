package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// Clickhouse
	Clickhouse Clickhouse `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
	// Clickhouse_Sinker
	Clickhouse_SinkerNali ClickhouseSinkerNali `mapstructure:"clickhouse_sinker_nali" json:"clickhouse_sinker_nali" yaml:"clickhouse_sinker_nali"`
	// DownloadPath
	DownloadPath DownloadPath `mapstructure:"download_path" json:"download_path" yaml:"download_path"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	Excel      Excel      `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
}
