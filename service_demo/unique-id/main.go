package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/WeixinX/graduation-project-service/service_demo/unique_id/config"
	"github.com/WeixinX/graduation-project/util/gin_mw"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type UniqueIDConfig struct {
	UniqueIDs []string `json:"unique_ids"`
}

func main() {
	// 读取命令行参数
	var uniqueIDList string
	var configFile string
	flag.StringVar(&uniqueIDList, "unique_id_list", "./unique_id.json", "unique id 列表文件路径")
	flag.StringVar(&configFile, "config_file", "", "unique id 配置文件路径")
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

	engine.GET("/get_unique_id",
		// TODO: 修改 IP
		gin_mw.PromMiddleWare(metrics, serviceName, "127.0.0.1", config.CONFIG_PARAMS.InstanceID),
		gin_mw.JaegerTracerMiddleWare(serviceName),

		func(ctx *gin.Context) {
			uid := getUniqueID(uniqueIDList)
			fmt.Println("uid: ", uid)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   map[string]interface{}{"uid": uid},
			})
		})

	engine.POST("/anomaly_inject", anomalyInjectHandler)
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("Unique ID service failed to start! err: ", err)
	}
}

// getUniqueID 通过配置文件获取 uid 列表
func getUniqueID(configFile string) string {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Failed to read json configuration file! err: ", err)
		return ""
	}

	uniqueIDConfig := &UniqueIDConfig{}
	err = json.Unmarshal(bytes, uniqueIDConfig)
	if err != nil {
		fmt.Println("json configuration file parsing failed! err: ", err)
		return ""
	}

	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	ind := random.Intn(len(uniqueIDConfig.UniqueIDs))

	return uniqueIDConfig.UniqueIDs[ind]
}
