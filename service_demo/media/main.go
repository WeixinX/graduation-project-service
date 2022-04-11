package main

import (
	"flag"
	"fmt"

	"github.com/WeixinX/graduation-project-service/service_demo/media/api"
	"github.com/WeixinX/graduation-project-service/service_demo/media/config"
	"github.com/WeixinX/graduation-project/util/gin_mw"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// 初始化全局 Tracer
	serviceName := config.CONFIG_PARAMS.ServiceName
	_, closer := gin_mw.NewGlobalJaegerTracer(serviceName, config.CONFIG_PARAMS.JaegerAgent)
	defer closer.Close()

	// 初始化指标采集
	metrics := gin_mw.NewPromMetrics()

	// 启动服务
	engine := gin.Default()

	engine.POST("/post_media",
		// TODO: 修改 IP
		gin_mw.PromMiddleWare(metrics, serviceName, "127.0.0.1", config.CONFIG_PARAMS.InstanceID),
		gin_mw.JaegerTracerMiddleWare(serviceName),

		api.PostMedia,
	)
	engine.POST("/anomaly_inject", anomalyInjectHandler)
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Media service failed to start! err: ", err)
	}
}
