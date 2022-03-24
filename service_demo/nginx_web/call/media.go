package call

import (
	"fmt"

	"WeixinX/graduation-project/service_demo/nginx_web/config"
	"WeixinX/graduation-project/service_demo/nginx_web/model"
	"WeixinX/graduation-project/service_demo/nginx_web/request"

	"github.com/gin-gonic/gin"
)

func PostMedia(ctx *gin.Context, postContent *model.PostContent, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["media"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "media service call url not configured",
		}
	}

	req := &request.RequestParams{
		URLStr: config.CONFIG_PARAMS.DownstreamCallPair["media"],
		Method: "POST",
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: model.Media{
			UserID:       postContent.UserID,
			TimeStamp:    postContent.TimeStamp,
			MediaContent: postContent.MediaContent,
		},
	}

	resp, err := request.HttpDo(ctx, req)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "PostMedia: " + err.Error(),
		}
	} else {
		ch <- model.ChError{
			IsError:  false,
			ErrorMsg: "",
		}
		fmt.Println("PostMedia resp: ", resp)
	}
}
