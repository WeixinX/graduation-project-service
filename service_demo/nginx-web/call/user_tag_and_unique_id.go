package call

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/config"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/model"
	"github.com/WeixinX/graduation-project-service/service_demo/nginx_web/request"
	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
)

func GetUserTagAndUniqueID(ctx *gin.Context, userID string, ch chan model.ChError) {
	err := getUserTag(ctx, userID)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: err.Error(),
		}
	}

	err = getUniqueID(ctx, userID)
	if err != nil {
		ch <- model.ChError{
			IsError:  true,
			ErrorMsg: err.Error(),
		}
	}

	ch <- model.ChError{
		IsError:  false,
		ErrorMsg: "",
	}
}

func getUserTag(ctx *gin.Context, userID string) error {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["user-tag"]; !ok {
		return errors.New("user_tag service call url not configured")
	}

	req := &xhttp.ReqParams{
		UrlStr: config.CONFIG_PARAMS.DownstreamCallPair["user-tag"],
		Method: http.MethodGet,
	}
	resp, err := request.XHttp.Do(ctx, req)
	if err != nil {
		return errors.New("GetUserTag: " + err.Error())

	} else {
		fmt.Println("user id: ", userID)
		fmt.Println("GetUserTag resp: " + resp)

		return nil
	}
}

func getUniqueID(ctx *gin.Context, userID string) error {
	if _, ok := config.CONFIG_PARAMS.DownstreamCallPair["unique-id"]; !ok {
		return errors.New("unique_id service call url not configured")
	}

	req := &xhttp.ReqParams{
		UrlStr: config.CONFIG_PARAMS.DownstreamCallPair["unique-id"],
		Method: http.MethodGet,
	}
	resp, err := request.XHttp.Do(ctx, req)
	if err != nil {
		return errors.New("GetUniqueID: " + err.Error())

	} else {
		fmt.Println("user id: ", userID)
		fmt.Println("GetUniqueID resp :" + resp)

		return nil
	}
}
