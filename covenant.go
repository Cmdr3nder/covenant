package main

import (
    "fmt"

    "github.com/zenazn/goji"
    "github.com/spf13/viper"

    "github.com/ender4021/covenant/controller"
)

func main() {
    readConfigFile()
    registerControllers()
    startServer()
}

func readConfigFile() {
    viper.SetConfigName("covenant_config")

    viper.AddConfigPath("$HOME/.covenant")

    err := viper.ReadInConfig()

    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}

func registerControllers() {
    controller.RegisterRootController()
    controller.RegisterStaticFileController()
    controller.RegisterHelloController()
    controller.RegisterViewController()
}

func startServer() {
    goji.Serve()
}
