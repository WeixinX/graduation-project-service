package main

import (
	"flag"
	"fmt"

	"github.com/WeixinX/graduation-project-service/service_demo/media/api"
	"github.com/WeixinX/graduation-project-service/service_demo/media/config"
	"github.com/WeixinX/graduation-project/util/gin_mw"
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
	engine.Use(gin_mw.JaegerTracerInit(config.CONFIG_PARAMS.ServiceName))

	engine.POST("/post_media", api.PostMedia)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Media service failed to start! err: ", err)
	}
}
