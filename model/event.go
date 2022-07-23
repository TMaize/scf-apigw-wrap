package model

import (
	"encoding/base64"
	"strings"
)

// APIGatewayRequest http event request
// https://github.com/tencentyun/scf-go-lib/blob/master/events/apigw.go
type APIGatewayRequest struct {
	Body            string                   `json:"body"`
	Headers         map[string]interface{}   `json:"headers"`
	HTTPMethod      string                   `json:"httpMethod"`
	IsBase64Encoded bool                     `json:"isBase64Encoded"`
	Path            string                   `json:"path"`
	QueryString     map[string][]interface{} `json:"queryString"`
	// RequestContext APIGatewayRequestContext `json:"requestContext"`
}

// CloudBaseGatewayRequest http event request
type CloudBaseGatewayRequest struct {
	Body                  string              `json:"body"`
	Headers               map[string]string   `json:"headers"`
	HTTPMethod            string              `json:"httpMethod"`
	IsBase64Encoded       bool                `json:"isBase64Encoded"`
	MultiValueHeaders     map[string][]string `json:"multiValueHeaders"`
	Path                  string              `json:"path"`
	QueryStringParameters map[string][]string `json:"queryStringParameters"`
	// RequestContext        CloudBaseGatewayRequestContext `json:"requestContext"`
}

type Request struct {
	Body        []byte              `json:"body"`
	Headers     map[string][]string `json:"headers"`
	HTTPMethod  string              `json:"httpMethod"`
	Path        string              `json:"path"`
	QueryString map[string][]string `json:"queryString"`
}

type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

func GetRequest(event map[string]interface{}) (Request, error) {
	request := Request{
		HTTPMethod: strings.ToUpper(event["httpMethod"].(string)),
		Path:       event["path"].(string),
		Body:       []byte{},
	}

	// body 是否为 base64 编码
	isBase64Encoded := false
	if _, ok := event["isBase64Encoded"]; ok {
		isBase64Encoded = event["isBase64Encoded"].(bool)
	}

	// 设置 body
	if _, ok := event["body"]; ok {
		if isBase64Encoded {
			body, err := base64.StdEncoding.DecodeString(event["body"].(string))
			if err != nil {
				return request, err
			}
			request.Body = body
		} else {
			request.Body = []byte(event["body"].(string))
		}
	}

	// 设置 queryString
	queryString := make(map[string]interface{})
	if _, ok := event["queryStringParameters"]; ok {
		temp := event["queryStringParameters"].(map[string]interface{})
		for k, v := range temp {
			queryString[k] = v
		}
	}
	if _, ok := event["queryString"]; ok {
		temp := event["queryString"].(map[string]interface{})
		for k, v := range temp {
			queryString[k] = v
		}
	}
	request.QueryString = ToStringArray(queryString)

	// 设置 headers
	headers := make(map[string]interface{})
	if _, ok := event["multiValueHeaders"]; ok {
		temp := event["multiValueHeaders"].(map[string]interface{})
		for k, v := range temp {
			headers[k] = v
		}
	}
	if _, ok := event["headers"]; ok {
		temp := event["headers"].(map[string]interface{})
		for k, v := range temp {
			headers[k] = v
		}
	}
	request.Headers = ToStringArray(headers)

	return request, nil
}

func ToStringArray(data map[string]interface{}) map[string][]string {
	result := make(map[string][]string)
	for k, v := range data {
		if arr, ok := v.([]interface{}); ok {
			strArr := make([]string, len(arr))
			for idx, item := range arr {
				strArr[idx] = item.(string)
			}
			result[k] = strArr
		}
		if str, ok := v.(string); ok {
			result[k] = []string{str}
		}
		if exist, ok := v.(bool); ok && exist {
			result[k] = []string{""}
		}
	}
	return result
}
