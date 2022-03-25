package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type ReqParams struct {
	UrlStr      string
	Method      string
	Header      map[string][]string
	QueryParams map[string]string
	Body        io.Reader
}

var CLIENT *http.Client

func HttpClientSetUp() *http.Client {
	// 初始化 client
	return &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 3,
	}
}

// HttpDo LB 的请求多半是透传，所以可能和其他服务的请求不太一样，需要注意
func HttpDo(ctx *gin.Context, params *ReqParams) (string, error) {
	// 构造 url, 设置查询参数
	// FIXME:没有服务涉及查询参数，故先不考虑
	Url, _ := url.Parse(params.UrlStr)
	p := url.Values{}
	for k, v := range params.QueryParams {
		p.Set(k, v)
	}
	Url.RawQuery = p.Encode()
	urlPath := Url.String()

	// 构造请求, 设置请求头
	req, err := http.NewRequest(params.Method, urlPath, params.Body)
	if err != nil {
		return "", err
	}
	req.Header = params.Header

	// 发起请求
	resp, err := CLIENT.Do(req)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	return string(content), err
}
