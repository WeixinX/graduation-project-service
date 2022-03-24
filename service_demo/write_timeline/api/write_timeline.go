package api

import (
	"net/http"

	"WeixinX/graduation-project/service_demo/write_timeline/db"

	"github.com/gin-gonic/gin"
)

func WriteTimeline(ctx *gin.Context) {
	text := db.Text{}
	if err := ctx.ShouldBindJSON(&text); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})

		go db.MongoDBPost(text)
		go db.RedisPost(text)
	}
}
