package config

type Kafka struct {
	KafkaAddr     string `mapstructure:"kafka_addr" json:"kafka_addr" yaml:"kafka_addr"` // 服务器地址和端口
	ZkAddr     string `mapstructure:"zk_addr" json:"zk_addr" yaml:"zk_addr"` // 服务器地址和端口
}
