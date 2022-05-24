package model

import (
	"encoding/base64"
	"strings"
)

// APIGatewayRequest http event request
// https://github.com/tencentyun/scf-go-lib/blob/master/events/apigw.go
type APIGatewayRequest struct {
	Body        string              `json:"body"`
	Headers     map[string]string   `json:"headers"`
	HTTPMethod  string              `json:"httpMethod"`
	Path        string              `json:"path"`
	QueryString map[string][]string `json:"queryString"`
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
	Headers     map[string]string   `json:"headers"`
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
	}

	var err error
	var isBase64Encoded bool
	var body string
	var headers map[string]interface{}
	var queryString map[string]interface{}

	if _, ok := event["body"]; ok {
		body = event["body"].(string)
	}

	if _, ok := event["isBase64Encoded"]; ok {
		isBase64Encoded = event["isBase64Encoded"].(bool)
	}

	if _, ok := event["headers"]; ok {
		headers = event["headers"].(map[string]interface{})
	} else {
		headers = map[string]interface{}{}
	}

	if _, ok := event["queryStringParameters"]; ok {
		queryString = event["queryStringParameters"].(map[string]interface{})
	} else if _, ok := event["queryString"]; ok {
		queryString = event["queryString"].(map[string]interface{})
	} else {
		queryString = map[string]interface{}{}
	}

	if isBase64Encoded {
		request.Body, err = base64.StdEncoding.DecodeString(body)
	} else {
		request.Body = []byte(body)
	}

	if err != nil {
		return request, err
	}

	request.Headers = make(map[string]string)
	for k, v := range headers {
		request.Headers[k] = v.(string)
	}

	request.QueryString = make(map[string][]string)
	for k, v := range queryString {
		arr, ok := v.([]interface{})
		if ok {
			strArr := make([]string, len(arr))
			for idx, item := range arr {
				strArr[idx] = item.(string)
			}
			request.QueryString[k] = strArr
		} else {
			request.QueryString[k] = []string{v.(string)}
		}

	}

	return request, err
}
