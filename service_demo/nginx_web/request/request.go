package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/WeixinX/graduation-project/util/xhttp"
	"github.com/gin-gonic/gin"
)

var XHttp *xhttp.Req

func XHttpInit() {
	// 初始化 http client
	XHttp := xhttp.NewReq()
}
