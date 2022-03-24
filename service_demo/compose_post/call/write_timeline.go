package call

import (
	"fmt"

	"WeixinX/graduation-project/service_demo/compose_post/config"
	"WeixinX/graduation-project/service_demo/compose_post/request"

	"github.com/gin-gonic/gin"
)

func WriteTimeline(ctx *gin.Context,body interface{}){
	req := &request.RequestParams{
		URLStr:  config.CONFIG_PARAMS.DownstreamCallPair["write_timeline"],
		Method: "POST",
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body:    body,
	}

	_,err := request.HttpDo(ctx,req)
	if err != nil{
		fmt.Println("Request failed, err: ",err.Error())
	}
}