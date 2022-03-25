package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/WeixinX/graduation-project-service/service_demo/media/db"
	"github.com/gin-gonic/gin"
)

const (
	// ComposePost 最长睡眠时间 30 ms
	CPMaxSleepMs = 30

	// ComposePost 最短睡眠时间 10 ms
	CPMinSleepMs = 10
)

func PostMedia(ctx *gin.Context) {
	media := db.Media{}
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})

		go db.MongoDBPost(ctx, media)
		go db.RedisPost(ctx, media)

		// 模拟 media 处理过程
		seed := rand.NewSource(time.Now().Unix())
		random := rand.New(seed)
		sleepTime := random.Intn(CPMaxSleepMs-CPMinSleepMs) + CPMinSleepMs
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		fmt.Printf("media exec time: %v\n", time.Millisecond*time.Duration(sleepTime))

		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
