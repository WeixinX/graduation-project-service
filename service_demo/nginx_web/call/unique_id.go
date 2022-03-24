package call

import (
	"fmt"

	"WeixinX/graduation-project/service_demo/nginx_web/config"
	"WeixinX/graduation-project/service_demo/nginx_web/model"
	"WeixinX/graduation-project/service_demo/nginx_web/request"

	"github.com/gin-gonic/gin"
)

func GetUniqueID(ctx *gin.Context, userID string, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["unique_id"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "unique_id service call url not configured",
		}
	}

	req := &request.RequestParams{
		URLStr: config.CONFIG_PARAMS.DownstreamCallPair["unique_id"],
		Method: "GET",
	}

	resp, err := request.HttpDo(ctx, req)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "GetUniqueID: " + err.Error(),
		}
	} else {
		ch <- model.ChError{
			IsError:  false,
			ErrorMsg: "",
		}
		fmt.Println("user id: ", userID)
		fmt.Println("GetUniqueID resp :" + resp)
	}
}
