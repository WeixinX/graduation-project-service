package call

import (
	"fmt"
	"net/http"

	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/config"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/model"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/request"
	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
)

func GetUniqueID(ctx *gin.Context, userID string, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["unique_id"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "unique_id service call url not configured",
		}
	}

	req := &xhttp.ReqParams{
		UrlStr:      config.CONFIG_PARAMS.DownstreamCallPair["unique_id"],
		Method:      http.MethodGet,
	}
	resp, err := request.XHttp.Do(ctx,req)
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
