package request

import (
	"github.com/WeixinX/graduation-project/util/xhttp"
)

var XHttp *xhttp.Req

func NewXHttpReq() *xhttp.Req {
	// 初始化 http client
	return xhttp.NewReq()
}