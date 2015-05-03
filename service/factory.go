package service

import (
    "github.com/spf13/viper"
)

var server = &gojiServer{}
var config = &viperConfig{v: viper.New()}

func GetServer() Server {
    return server
}

func GetConfig() Config {
    return config
}
