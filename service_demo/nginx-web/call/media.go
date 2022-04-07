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

func PostMedia(ctx *gin.Context, postContent *model.PostContent, ch chan model.ChError) {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["media"]; !ok {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "media service call url not configured",
		}
	}

	bodyBytes, err := json.Marshal(model.Media{
		UserID:       postContent.UserID,
		TimeStamp:    postContent.TimeStamp,
		MediaContent: postContent.MediaContent,
	})
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: "PostMedia: " + err.Error(),
		}
	}
	req := &xhttp.ReqParams{
		UrlStr: config.CONFIG_PARAMS.DownstreamCallPair["media"],
		Method: http.MethodPost,
		// map[string][]string{"Content-Type": {"application/json"}}
		Header: ctx.Request.Header,
		Body:   strings.NewReader(string(bodyBytes)),
	}

	resp, err := request.XHttp.Do(ctx, req)
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
