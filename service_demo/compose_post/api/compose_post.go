package api

import (
	"fmt"
	"github.com/WeixinX/graduation-project-service/service_demo/compose_post/call"
	"math/rand"
	"net/http"
	"time"

	"github.com/WeixinX/graduation-project-service/service_demo/compose_post/model"
	"github.com/gin-gonic/gin"
)

const (
	// ComposePost 最长睡眠时间 50 ms
	CPMaxSleepMs = 50

	// ComposePost 最短睡眠时间 20 ms
	CPMinSleepMs = 20
)

func ComposePost(ctx *gin.Context) {
	text := model.Text{}
	if err := ctx.ShouldBindJSON(&text); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		//fmt.Printf("text info: %+v\n",text)

		go call.WriteTimeline(ctx, text)

		// 模拟 ComposePost 处理过程
		seed := rand.NewSource(time.Now().Unix())
		random := rand.New(seed)
		sleepTime := random.Intn(CPMaxSleepMs-CPMinSleepMs) + CPMinSleepMs
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		fmt.Printf("compose post exec time: %v\n", time.Millisecond*time.Duration(sleepTime))

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
