package main

import (
	"flag"
	"fmt"

	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/api"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/config"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/request"
	"github.com/WeixinX/graduation-project/util/gin_mw"
	"github.com/gin-gonic/gin"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "nginx web 配置文件路径，默认为空")
	flag.Parse()

	// 读取配置文件
	config.CONFIG_PARAMS = config.ConfigSetUp(configFile)
	if config.CONFIG_PARAMS == nil {
		fmt.Println("Failed to read configuration file initialization!")
		return
	}

	// 初始化 Http client
	request.XHttp = request.NewXHttpReq()

	// 初始化全局 Tracer
	_, closer := gin_mw.NewGlobalJaegerTracer(config.CONFIG_PARAMS.ServiceName)
	defer closer.Close()

	// 启动服务
	engine := gin.Default()
	engine.Use(gin_mw.JaegerTracerInit(config.CONFIG_PARAMS.ServiceName))

	engine.POST("/do_post", api.DoPost)

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Nginx Web service failed to start! err: ", err)
	}
}
