package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/WeixinX/graduation-project-service/service_demo/user_tag/config"
	"github.com/WeixinX/graduation-project/util/gin_mw"
	"github.com/gin-gonic/gin"
)

type UserTagConfig struct {
	UserTags []string `json:"user_tags"`
}

func main() {
	// 读取命令行参数
	var userTagList string
	var configFile string
	flag.StringVar(&userTagList, "user_tag_list", "./user_tag.json", "user tag 列表文件路径")
	flag.StringVar(&configFile, "config_file", "", "user tag 配置文件路径")
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

	engine.GET("/get_user_tag", func(ctx *gin.Context) {
		userTag := getUniqueID(userTagList)
		fmt.Println("user tag: ", userTag)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   map[string]interface{}{"user_tag": userTag},
		})
	})

	if err := engine.Run(":" + config.CONFIG_PARAMS.Port); err != nil {
		fmt.Println("User Tag service failed to start! err: ", err)
	}
}

// getUniqueID 通过配置文件获取 uid 列表
func getUniqueID(configFile string) string {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Failed to read json configuration file! err: ", err)
		return ""
	}

	userTagConfig := &UserTagConfig{}
	err = json.Unmarshal(bytes, userTagConfig)
	if err != nil {
		fmt.Println("json configuration file parsing failed! err: ", err)
		return ""
	}

	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	ind := random.Intn(len(userTagConfig.UserTags))

	return userTagConfig.UserTags[ind]
}
