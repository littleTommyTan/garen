package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	WoaAppId           string `yaml:"woa_appid"`
	WoaAppSecret       string `yaml:"woa_app_secret"`
	DSN                string `yaml:"dsn"`
	RedisHost          string `yaml:"redis_host"`
	RedisPort          int    `yaml:"redis_port"`
	RedisPwd           string `yaml:"redis_pwd"`
	OssAccessKey       string `yaml:"aliyun_oss_accesskey"`
	OssAccessKeySecret string `yaml:"aliyun_oss_accesskey_secret"`
}

var configuration *Configuration

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

func GetConfiguration() *Configuration {
	return configuration
}
