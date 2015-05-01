package service

var server = &gojiServer{}

func GetServer() Server {
    return server
}
