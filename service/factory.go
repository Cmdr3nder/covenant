package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/master"
	"github.com/ender4021/covenant/service/util"
	"github.com/spf13/viper"
)

var server = &gojiServer{}
var vConfig = viper.New()
var layoutMap = make(map[string]layout.Layout)

// GetServer returns the single instance of Server
func GetServer() Server {
	return server
}

// GetConfig returns the single instance of Config
func GetConfig() config.Config {
	return vConfig
}

// GetRouteBuilder returns a new RouteBuilder
func GetRouteBuilder() RouteBuilder {
	return &goRouteBuilder{buffer: bytes.Buffer{}}
}

// GetLayout returns a new layout for the given path or the same instance if previously called
func GetLayout(configPath string) layout.Layout {
	if layoutMap[configPath] == nil || vConfig.GetBool("debug") {
		layoutPath := vConfig.GetString(configPath)

		t, err := template.ParseFiles(layoutPath)

		if err != nil {
			panic(fmt.Errorf("Could not read template: %s \n", err))
		}

		layoutMap[configPath] = master.New(t, readStyleSheets(layoutPath), readScripts(layoutPath))
	}

	return layoutMap[configPath]
}

// GetRootLayout returns a layout object for the primary layout file
func GetRootLayout() layout.Layout {
	return GetLayout("views_shared_layout")
}

func readScripts(layoutPath string) []template.HTMLAttr {
	return readExtras(layoutPath + ".scripts")
}

func readStyleSheets(layoutPath string) []template.HTMLAttr {
	return readExtras(layoutPath + ".styles")
}

func readExtras(extrasPath string) []template.HTMLAttr {
	lines, err := util.ReadFileAsLines(extrasPath)
	var attrs []template.HTMLAttr

	if err == nil {
		for _, line := range lines {
			attrs = append(attrs, template.HTMLAttr(line))
		}
	}

	return attrs
}
