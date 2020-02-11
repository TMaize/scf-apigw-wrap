package scf_apigw_wrap

import (
	"github.com/tencentyun/scf-go-lib/events"
	"net/http"
	"net/http/httptest"
	"strings"
)

// Wrap 请求包装+委托模拟http请求
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
		IsBase64Encoded: false,
		Body:            response.Body.String(),
		StatusCode:      response.Code,
		Headers:         map[string]string{},
	}

	for k := range response.Header() {
		gwResp.Headers[k] = response.Header().Get(k)
	}

	_, ok := gwResp.Headers["Content-Type"]
	if !ok && gwResp.StatusCode == 200 {
		gwResp.Headers["Content-Type"] = "application/octet-stream"
	}
	if !ok && gwResp.StatusCode == 404 {
		gwResp.Headers["Content-Type"] = "text/plain"
	}

	return gwResp
}
