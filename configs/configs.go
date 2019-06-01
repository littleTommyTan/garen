package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration 配置结构体
type Configuration struct {
	WoaAppID           string `yaml:"woa_appid"`
	WoaAppSecret       string `yaml:"woa_app_secret"`
	DSN                string `yaml:"dsn"`
	RedisHost          string `yaml:"redis_host"`
	RedisPort          int    `yaml:"redis_port"`
	RedisPwd           string `yaml:"redis_pwd"`
	OssAccessKey       string `yaml:"aliyun_oss_accesskey"`
	OssAccessKeySecret string `yaml:"aliyun_oss_accesskey_secret"`
	LogEndpoint        string `yaml:"aliyun_log_endpoint"`
	LogAccessKey       string `yaml:"aliyun_log_accesskey"`
	LogAccessKeySecret string `yaml:"aliyun_log_accesskey_secret"`
}

var configuration *Configuration

// LoadConfiguration 加载配置
func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

// GetConfiguration 暴露配置
func GetConfiguration() *Configuration {
	return configuration
}
