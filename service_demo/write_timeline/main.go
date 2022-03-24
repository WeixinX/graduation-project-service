package main

import (
	"flag"
	"fmt"

	"WeixinX/graduation-project/service_demo/write_timeline/api"
	"WeixinX/graduation-project/service_demo/write_timeline/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "write timeline 配置文件路径，默认为空")
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

	engine.POST("/write_timeline",api.WriteTimeline)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("write timeline service failed to start! err: ", err)
	}
}