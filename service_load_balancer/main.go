package main

import (
	"flag"
	"fmt"

	"github.com/WeixinX/graduation-project-service/service_load_balancer/config"
	"github.com/WeixinX/graduation-project-service/service_load_balancer/load_balancer"
	"github.com/WeixinX/graduation-project-service/service_load_balancer/request"
	"github.com/WeixinX/graduation-project-service/service_load_balancer/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "lb 配置文件路径，默认为空")
	flag.Parse()

	// 读取配置文件
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	if config.CONFIG_PARAMS == nil {
		fmt.Println("Failed to read configuration file initialization!")
		return
	}
	load_balancer.INSTANCE_LIST = config.CONFIG_PARAMS.InstanceList

	// 初始化 Http client
	request.CLIENT = request.HttpClientSetUp()

	// 启动服务
	engine := gin.Default()
	// 不使用中间件采集 lb 的 traces 和 metrics
	engine.Use()

	router.RouterSetUp(engine)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Load Balancer service failed to start! err: ", err)
	}
}
