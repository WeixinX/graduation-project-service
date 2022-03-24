package call

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/config"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/model"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/request"
	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
)

func PostText(ctx *gin.Context, postContent *model.PostContent, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["text"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "text service call url not configured",
		}
	}

	bodyBytes,err := json.Marshal(model.Text{
		UserID:      postContent.UserID,
		TimeStamp:   postContent.TimeStamp,
		TextContent: postContent.MediaContent,
	})
	if err != nil{
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "PostText: "+err.Error(),
		}
	}
	req := &xhttp.ReqParams{
		UrlStr:      config.CONFIG_PARAMS.DownstreamCallPair["text"],
		Method:      http.MethodPost,
		Header:      map[string][]string{"Content-Type": {"application/json"}},
		Body:        strings.NewReader(string(bodyBytes)),
	}
	resp, err := request.XHttp.Do(ctx,req)
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
