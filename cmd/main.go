package cmd

import (
	"HubuLearn/conf"
	"HubuLearn/route"
)

func main() {
	loading()
	router := route.NewRouter()
	router.Run(conf.HttpPort)
}

func loading() {
	// 从配置文件读入配置
	conf.Init()
}
