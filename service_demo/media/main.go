package main

import (
	"flag"
	"fmt"

	"WeixinX/graduation-project/service_demo/media/api"
	"WeixinX/graduation-project/service_demo/media/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "media 配置文件路径")
	flag.Parse()

	// 读取配置文件
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	if config.CONFIG_PARAMS == nil {
		fmt.Println("Failed to read configuration file initialization!")
		return
	}

	// 启动服务
	engine := gin.Default()
	engine.Use()

	engine.POST("/post_media", api.PostMedia)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Media service failed to start! err: ", err)
	}
}
