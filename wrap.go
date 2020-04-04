package scf_apigw_wrap

import (
	"encoding/base64"
	"github.com/tencentyun/scf-go-lib/events"
	"net/http"
	"net/http/httptest"
	"strings"
)

var textTypes = []string{"text/", "javascript", "json"}

// Wrap 请求包装+委托模拟http请求
// https://cloud.tencent.com/document/product/583/12513
// pathname 为请求路径，由于环境的原因，
func Wrap(event events.APIGatewayRequest, pathname string, h http.Handler) events.APIGatewayResponse {

	requestBody := strings.NewReader(event.Body)
	request := httptest.NewRequest(event.Method, pathname, requestBody)

	requestQuery := request.URL.Query()
	for k, arr := range event.QueryString {
		for _, v := range arr {
			requestQuery.Add(k, v)
		}
	}
	request.URL.RawQuery = requestQuery.Encode()

	for k, v := range event.Headers {
		request.Header.Set(k, v)
	}

	response := httptest.NewRecorder()

	// 模拟请求
	h.ServeHTTP(response, request)

	// 转换为APIGatewayResponse
	gwResp := events.APIGatewayResponse{
		StatusCode: response.Code,
		Headers:    map[string]string{},
	}

	for k := range response.Header() {
		gwResp.Headers[k] = response.Header().Get(k)
	}

	// 判断是否需要对body编码
	bodyBase64 := true
	contentType, ok := gwResp.Headers["Content-Type"]
	if !ok {
		contentType = "application/octet-stream"
		gwResp.Headers["Content-Type"] = contentType
	}
	for _, v := range textTypes {
		if strings.Contains(contentType, v) {
			bodyBase64 = false
			break
		}
	}
	gwResp.IsBase64Encoded = bodyBase64

	if bodyBase64 {
		gwResp.Body = base64.StdEncoding.EncodeToString(response.Body.Bytes())
	} else {
		gwResp.Body = response.Body.String()
	}

	return gwResp
}
