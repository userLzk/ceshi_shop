package config

import (
	"gopkg.in/yaml.v3"
	"fmt"
	"io/ioutil"
)

var (
	cfgFile = "./conf/conf.yaml"
)
//定义结构
type MysqlConfig struct {
	Host     string
	Port     string
	User string
	PassWord string
	NetWork string
	Databases string
}
type RedisConfig []struct {
	Host string
	PassWord string
	Port     int
	Db int
}

//监听客户端信息
type ClientConfig struct {

}

type ConfigData struct {
	MysqlConfig
	RedisConfig
}
func GetConfigDesc() *ConfigData {
	data, err := ioutil.ReadFile(cfgFile)
	cfg := &ConfigData{}
	if err != nil {
		panic(err)
	}
	//序列化 数据结构
	err = yaml.Unmarshal(data,cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}