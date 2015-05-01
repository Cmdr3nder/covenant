package main

import (
    "fmt"

    "github.com/spf13/viper"

    "github.com/ender4021/covenant/controller"
    "github.com/ender4021/covenant/service"
)

func main() {
    server := service.GetServer()

    readConfigFile()

    //Register Controllers
    controller.RegisterRootController(server)
    controller.RegisterStaticFileController(server)
    controller.RegisterHelloController(server)
    controller.RegisterViewController(server)

    server.Serve()
}

func readConfigFile() {
    viper.SetConfigName("covenant_config")

    viper.AddConfigPath("$HOME/.covenant")

    err := viper.ReadInConfig()

    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}
