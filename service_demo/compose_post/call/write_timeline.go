package call

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/WeixinX/graduation-project-service/service_demo/compose_post/config"
	"github.com/WeixinX/graduation-project-service/service_demo/compose_post/model"
	"github.com/WeixinX/graduation-project-service/service_demo/compose_post/request"
	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
)

func WriteTimeline(ctx *gin.Context, text model.Text) {
	bodyBytes, err := json.Marshal(text)
	if err != nil {
		fmt.Printf("Json marshal failed: %s\n", err)
	}
	req := &xhttp.ReqParams{
		UrlStr: config.CONFIG_PARAMS.DownstreamCallPair["write_timeline"],
		Method: http.MethodPost,
		Header: map[string][]string{"Content-Type": {"application/json"}},
		Body:   strings.NewReader(string(bodyBytes)),
	}
	_, err = request.XHttp.Do(ctx, req)
	if err != nil {
		fmt.Println("Request failed, err: ", err.Error())
	}
}
