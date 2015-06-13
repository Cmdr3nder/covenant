package main

import (
	"fmt"

	"github.com/ender4021/covenant/controller"
	"github.com/ender4021/covenant/service"
)

func main() {
	server := service.GetServer()
	config := service.GetConfig()

	config.SetConfigName("covenant_config")
	config.AddConfigPath("$HOME/.covenant")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	setupViewConfigDefaults(config)

	//Register Controllers
	controller.RegisterRootController(server, config)
	controller.RegisterBlogController(server)
	controller.RegisterResumeController(server)
	controller.RegisterStaticFileController(server, config)

	server.Serve()
}

func setupViewConfigDefaults(config service.Config) {
	config.SetDefault("views", "./view")

	setupSharedViewsConfig(config)
	setupRootViewsConfig(config)

	config.SetDefault("media", "./media")
	config.SetDefault("css", config.GetString("media")+"/css")
	config.SetDefault("js", config.GetString("media")+"/js")
	config.SetDefault("img", config.GetString("media")+"/img")
	config.SetDefault("favicon", config.GetString("img")+"/favicon.ico")
}

func setupSharedViewsConfig(config service.Config) {
	config.SetDefault("views_shared", config.GetString("views")+"/shared")

	config.SetDefault("views_shared_layout", config.GetString("views_shared")+"/layout.html")
}

func setupRootViewsConfig(config service.Config) {
	config.SetDefault("views_root", config.GetString("views")+"/root")

	config.SetDefault("views_root_index", config.GetString("views_root")+"/index.html")
}
