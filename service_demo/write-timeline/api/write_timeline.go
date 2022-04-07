package api

import (
	"net/http"
	"sync"

	"github.com/WeixinX/graduation-project-service/service_demo/write_timeline/db"

	"github.com/gin-gonic/gin"
)

func WriteTimeline(ctx *gin.Context) {
	text := db.Text{}
	if err := ctx.ShouldBindJSON(&text); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			db.MongoDBPost(ctx, text)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			db.RedisPost(ctx, text)
		}()

		wg.Wait()
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
