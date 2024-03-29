package scf_apigw_wrap

import (
	"bytes"
	"encoding/base64"
	"github.com/TMaize/scf-apigw-wrap/model"
	"net/http"
	"net/http/httptest"
	"strings"
)

var textTypes = []string{"text/", "javascript", "json", "/xml"}

func isNoContentTypeCode(code int) bool {
	return code == 204 || code == 304 || code == 401
}

func WrapRequest(gwRequest model.Request, requestPath string, h http.Handler) model.Response {
	request := httptest.NewRequest(gwRequest.HTTPMethod, requestPath, bytes.NewReader(gwRequest.Body))

	requestQuery := request.URL.Query()
	for k, arr := range gwRequest.QueryString {
		for _, v := range arr {
			requestQuery.Add(k, v)
		}
	}
	request.URL.RawQuery = requestQuery.Encode()

	for k, arr := range gwRequest.Headers {
		for _, v := range arr {
			request.Header.Add(k, v)
		}
	}

	response := httptest.NewRecorder()

	// 模拟请求
	h.ServeHTTP(response, request)

	gwResp := model.Response{
		StatusCode: response.Code,
		Headers:    map[string]string{},
	}

	for k := range response.Header() {
		gwResp.Headers[k] = response.Header().Get(k)
	}

	// 对于没有Content-Type的，全部按照二进制处理
	_, ok := gwResp.Headers["Content-Type"]
	if !ok && !isNoContentTypeCode(gwResp.StatusCode) {
		gwResp.Headers["Content-Type"] = "application/octet-stream"
	}

	// 判断是否需要对body编码
	bodyBase64 := true
	contentType := gwResp.Headers["Content-Type"]
	if contentType == "contentType" {
		bodyBase64 = false
	} else {
		for _, v := range textTypes {
			if strings.Contains(contentType, v) {
				bodyBase64 = false
				break
			}
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

// Wrap 执行HTTP触发器请求体
// https://cloud.tencent.com/document/product/583/12513
func Wrap(event map[string]interface{}, requestPath string, handler http.Handler) model.Response {
	request, err := model.GetRequest(event)

	if err != nil {
		return model.Response{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            "parse event fail",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}
	}
	return WrapRequest(request, requestPath, handler)
}
