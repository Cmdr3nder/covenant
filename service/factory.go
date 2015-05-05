package service

import (
    "bytes"

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

func GetRouteBuilder() RouteBuilder {
    return &goRouteBuilder{buffer: bytes.Buffer{}}
}
