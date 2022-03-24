package call

import (
	"fmt"

	"WeixinX/graduation-project/service_demo/nginx_web/config"
	"WeixinX/graduation-project/service_demo/nginx_web/model"
	"WeixinX/graduation-project/service_demo/nginx_web/request"

	"github.com/gin-gonic/gin"
)

func PostText(ctx *gin.Context, postContent *model.PostContent, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["text"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "text service call url not configured",
		}
	}

	req := &request.RequestParams{
		URLStr: config.CONFIG_PARAMS.DownstreamCallPair["text"],
		Method: "POST",
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: model.Text{
			UserID:      postContent.UserID,
			TimeStamp:   postContent.TimeStamp,
			TextContent: postContent.MediaContent,
		},
	}

	resp, err := request.HttpDo(ctx, req)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "PostText: " + err.Error(),
		}
	} else {
		ch <- model.ChError{
			IsError:  false,
			ErrorMsg: "",
		}
		fmt.Println("PostText resp: ", resp)
	}
}
