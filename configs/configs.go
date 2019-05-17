package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	//PGHost    string `yaml:"pg_host"`
	//PGPort    int    `yaml:"pg_port"`
	//PGDb      string `yaml:"pg_db"`
	//PGUser    string `yaml:"pg_user"`
	//PGPwd     string `yaml:"pg_pwd"`
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
