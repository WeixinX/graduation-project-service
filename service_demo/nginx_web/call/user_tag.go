package call

import (
	"fmt"

	"WeixinX/graduation-project/service_demo/nginx_web/config"
	"WeixinX/graduation-project/service_demo/nginx_web/model"
	"WeixinX/graduation-project/service_demo/nginx_web/request"

	"github.com/gin-gonic/gin"
)

func GetUserTag(ctx *gin.Context, userID string, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["user_tag"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "user_tag service call url not configured",
		}
	}

	req := &request.RequestParams{
		URLStr: config.CONFIG_PARAMS.DownstreamCallPair["user_tag"],
		Method: "GET",
	}

	resp, err := request.HttpDo(ctx, req)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "GetUserTag: " + err.Error(),
		}
	} else {
		ch <- model.ChError{
			IsError:  false,
			ErrorMsg: "",
		}
		fmt.Println("user id: ", userID)
		fmt.Println("GetUserTag resp: " + resp)
	}
}
