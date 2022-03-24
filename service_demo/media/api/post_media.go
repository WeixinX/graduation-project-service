package api

import (
	"net/http"

	"github.com/WeixinX/graduation-project-service/service_demo/media/db"
	"github.com/gin-gonic/gin"
)

func PostMedia(ctx *gin.Context) {
	media := db.Media{}
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})

		go db.MongoDBPost(media)
		go db.RedisPost(media)
	}
}
