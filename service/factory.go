package service

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/master"
	"github.com/ender4021/covenant/service/route"
	"github.com/ender4021/covenant/service/server"
	"github.com/ender4021/covenant/service/util"
	"github.com/spf13/viper"
)

var gServer = server.NewGoji()
var vConfig = viper.New()
var layoutMap = make(map[string]layout.Layout)

// GetServer returns the single instance of Server
func GetServer() server.Server {
	return gServer
}

// GetConfig returns the single instance of Config
func GetConfig() config.Config {
	return vConfig
}

// GetRouteBuilder returns a new route.Builder
func GetRouteBuilder() route.Builder {
	return route.NewBuilder()
}

var funcMap = template.FuncMap{
	"ToLower":    strings.ToLower,
	"MonthAsInt": util.MonthAsInt,
}

// GetLayout returns a new layout for the given path or the same instance if previously called
func GetLayout(configPath string) (layout.Layout, error) {
	if layoutMap[configPath] == nil || vConfig.GetBool("debug") {
		layoutPath := vConfig.GetString(configPath)

		fmt.Printf("%v\n with %+v\n at %v\n", configPath, funcMap, layoutPath)
		t, err := template.New(filepath.Base(layoutPath)).Funcs(funcMap).ParseFiles(layoutPath)

		if err != nil {
			return layoutMap[configPath], err
		}

		layoutMap[configPath] = master.New(t, readStyleSheets(layoutPath), readScripts(layoutPath))
	}

	return layoutMap[configPath], nil
}

// GetRootLayout returns a layout object for the primary layout file
func GetRootLayout() (layout.Layout, error) {
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
