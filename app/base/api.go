package base

import (
	"bytes"
	"net/http"
	"time"

	messageEnum "github.com/gin-qt-business/app/errors"
)

const (
	FAIL_CODE = iota
	SUCCESS_CODE
)

type Response struct {
	Code      int         `json:"code"`           // 1：成功 | 0：失败
	Message   string      `json:"message"`        // 成功：success | 失败：报错信息
	Timestamp int         `json:"timestamp"`      // 当前时间
	Data      interface{} `json:"data,omitempty"` // 返回参数内容主体
}

func NewResponseSuccess(data interface{}) *Response {
	return newResponse(SUCCESS_CODE, messageEnum.SUCCESS_MESSAGE, data)
}

func NewResponseError(message string) *Response {
	return newResponse(FAIL_CODE, message, nil)
}

func newResponse(code int, message string, data interface{}) *Response {
	res := new(Response)
	res.Code = code
	res.Message = message
	res.Timestamp = int(time.Now().Unix())
	res.Data = data
	return res
}

// HTTPClient 是自定义的HTTP客户端
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient 用于创建新的HTTP客户端
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get 用于发送GET请求
func (c *HTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// Post 用于发送POST请求
func (c *HTTPClient) Post(url string, body []byte, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.client.Do(req)
}
