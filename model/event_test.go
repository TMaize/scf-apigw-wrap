package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBase64Encoded(t *testing.T) {
	data := `
		{
		  "body": "Nzg5",
		  "headers": {},
		  "httpMethod": "GET",
		  "isBase64Encoded": true,
		  "multiValueHeaders": {},
		  "path": "/test",
		  "queryStringParameters": {},
		  "requestContext": {}
		}
`

	event := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}

	request, err := GetRequest(event)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(request.Body))
}

func TestQuery(t *testing.T) {
	data := `
		{
		  "headerParameters": {},
		  "headers": {},
		  "httpMethod": "GET",
		  "isBase64Encoded": false,
		  "path": "test",
		  "pathParameters": {},
		  "queryString": { "content1": ["1", "2"], "content2": "1", "content3": true },
		  "queryStringParameters": { "content4": ["1", "2"] },
		  "requestContext": {}
		}
`

	event := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}

	request, err := GetRequest(event)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(request.QueryString)
}

func TestHeaders(t *testing.T) {
	data := `
		{
		  "headerParameters": {},
		  "headers": {"content1": ["1", "2"], "content2": "1"},
          "multiValueHeaders": {"content3": ["1", "2"], "content4": "1"},
		  "httpMethod": "GET",
		  "isBase64Encoded": false,
		  "path": "test",
		  "pathParameters": {},
		  "requestContext": {}
		}
`

	event := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}

	request, err := GetRequest(event)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(request.Headers)
}
