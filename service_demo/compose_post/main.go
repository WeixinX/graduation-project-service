package main

import (
	"flag"
	"fmt"

	"WeixinX/graduation-project/service_demo/compose_post/api"
	"WeixinX/graduation-project/service_demo/compose_post/config"
	"WeixinX/graduation-project/service_demo/compose_post/request"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "compose post 配置文件路径，默认为空")
	flag.Parse()

	// 读取配置文件
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	if config.CONFIG_PARAMS == nil {
		fmt.Println("Failed to read configuration file initialization!")
		return
	}

	// 初始化 Http client
	request.CLIENT = request.HttpClientSetUp()

	// 启动服务
	engine := gin.Default()
	engine.Use()

	engine.POST("/compose_post", api.ComposePost)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Compose Post service failed to start! err: ", err)
	}
}
