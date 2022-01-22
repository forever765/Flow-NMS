package config

type DownloadPath struct {
	Prom2Json	string `mapstructure:"prom2json" json:"prom2json" yaml:"prom2json"` // prom2json保存目录
}
