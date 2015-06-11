package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/master"
	"github.com/spf13/viper"
)

var server = &gojiServer{}
var config = viper.New()
var layoutMap = make(map[string]layout.Layout)

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

// GetLayout returns a new layout for the given path or the same instance if previously called
func GetLayout(configPath string) layout.Layout {
	if layoutMap[configPath] == nil {
		layoutPath := config.GetString(configPath)

		t, err := template.ParseFiles(layoutPath)

		if err != nil {
			panic(fmt.Errorf("Could not read template: %s \n", err))
		}

		layoutMap[configPath] = master.New(t)
	}

	return layoutMap[configPath]
}
