package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	URLStr  string
	Method  string
	Headers map[string][]string
	Params  map[string]string
	Body    interface{}
}

var CLIENT *http.Client

func HttpClientSetUp() *http.Client {
	// 初始化 client
	return &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 5,
	}
}

func HttpDo(ctx *gin.Context, requestParams *RequestParams) (string, error) {
	// 初始化 client
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{
	//	Timeout:   time.Second * 5, //默认5秒超时时间
	//	Transport: tr,
	//}

	// 构造 url, 设置查询参数
	Url, _ := url.Parse(requestParams.URLStr)
	p := url.Values{}
	for k, v := range requestParams.Params {
		p.Set(k, v)
	}
	Url.RawQuery = p.Encode()
	urlPath := Url.String()

	// JSON 序列化请求体
	jsonBytes, err := json.Marshal(requestParams.Body)
	if err != nil {
		return "", err
	}

	// 构造请求, 设置请求头
	req, err := http.NewRequest(requestParams.Method, urlPath, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header = requestParams.Headers

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
	return string(content), nil
}
