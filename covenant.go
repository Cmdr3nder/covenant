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
    SetupViewConfigDefaults(config)

    //Register Controllers
    controller.RegisterRootController(server, config)
    controller.RegisterBlogController(server)
    controller.RegisterResumeController(server)
    controller.RegisterStaticFileController(server)

    server.Serve()
}

func SetupViewConfigDefaults(config service.Config) {
    config.SetDefault("views_root", "./view")
    config.SetDefault("views_shared", config.GetString("views_root") + "/shared")
    config.SetDefault("views_index", config.GetString("views_shared") + "/index.html")
    fmt.Printf("Views Index: %s\n", config.GetString("views_index"))
}
