package api

import (
	"net/http"
	"strings"

	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/call"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/model"

	"github.com/gin-gonic/gin"
)

func DoPost(ctx *gin.Context) {
	postContent := model.PostContent{}
	err := ctx.ShouldBindJSON(&postContent)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		// 使用 ch 收集 goroutine 中的错误
		chs := make([]chan model.ChError, 3)
		for i := 0; i < 3; i++ {
			chs[i] = make(chan model.ChError)
		}

		// GetUserTag 与 GetUniqueID 串行, 其余并行
		go call.GetUserTagAndUniqueID(ctx, postContent.UserID, chs[0])
		go call.PostMedia(ctx, &postContent, chs[1])
		go call.PostText(ctx, &postContent, chs[2])

		errorMsgList := make([]string, 0, 3)
		for _, ch := range chs {
			v := <-ch
			if v.IsError {
				errorMsgList = append(errorMsgList, v.ErrorMsg)
			}
		}

		if len(errorMsgList) == 0 {
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "error", "message": strings.Join(errorMsgList, ";")})
		}
	}
}
