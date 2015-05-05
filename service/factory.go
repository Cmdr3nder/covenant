package service

import (
	"bytes"

	"github.com/spf13/viper"
)

var server = &gojiServer{}
var config = &viperConfig{v: viper.New()}

// GetServer returns the single instance of Server
func GetServer() Server {
	return server
}

// GetConfig returns the single instance of Config
func GetConfig() Config {
	return config
}

// GetRouteBuilder returns a new RouteBuilder
func GetRouteBuilder() RouteBuilder {
	return &goRouteBuilder{buffer: bytes.Buffer{}}
}
