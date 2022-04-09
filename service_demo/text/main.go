package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/WeixinX/graduation-project-service/service_demo/text/config"
	"github.com/WeixinX/graduation-project-service/service_demo/text/request"
	"github.com/WeixinX/graduation-project/util/gin_mw"
	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 读取命令行参数
	var configFile string
	flag.StringVar(&configFile, "config_file", "", "text 配置文件路径")
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
	serviceName := config.CONFIG_PARAMS.ServiceName
	_, closer := gin_mw.NewGlobalJaegerTracer(serviceName, config.CONFIG_PARAMS.JaegerAgent)
	defer closer.Close()

	// 初始化指标采集
	metrics := gin_mw.NewPromMetrics()

	// 启动服务
	engine := gin.Default()

	engine.POST("/post_text",
		// TODO: 修改 IP
		gin_mw.PromMiddleWare(metrics, serviceName, "127.0.0.1", config.CONFIG_PARAMS.InstanceID),
		gin_mw.JaegerTracerMiddleWare(serviceName),

		PostText,
	)
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Text service failed to start! err: ", err)
	}
}

type Text struct {
	UserID      string `json:"user_id"`
	TimeStamp   string `json:"time_stamp"`
	TextContent string `json:"text_content"`
}

func PostText(ctx *gin.Context) {
	text := Text{}
	if err := ctx.ShouldBindJSON(&text); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		//ctx.JSON(http.StatusOK, gin.H{"status": "test", "message": text})

		bodyBytes, err := json.Marshal(text)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
		}
		req := &xhttp.ReqParams{
			UrlStr: config.CONFIG_PARAMS.DownstreamCallPair["compose-post"],
			Method: http.MethodPost,
			// map[string][]string{"Content-Type": {"application/json"}}
			Header: ctx.Request.Header,
			Body:   strings.NewReader(string(bodyBytes)),
		}

		_, err = request.XHttp.Do(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		}

	}
}
