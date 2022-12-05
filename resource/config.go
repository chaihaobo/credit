// Package resource
// @author： Boice
// @createTime：2022/11/28 11:35
package resource

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Addr  string `yaml:"addr"`
		Model string `yaml:"model"`
	} `yaml:"server"`
	Log struct {
		File  string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"log"`
	DB struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Database    string `yaml:"database"`
		MaxOpenConn int    `yaml:"maxOpenConn"`
		MaxIdleConn int    `yaml:"maxIdleConn"`
	} `yaml:"db"`
}

func newConfig(configPath string) Config {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var config Config
	viper.Unmarshal(&config)
	return config
}
