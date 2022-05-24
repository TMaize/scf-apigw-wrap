package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetRequest(t *testing.T) {
	data1 := `{"body":"Nzg5","headers":{"accept":"*/*","accept-encoding":"gzip, deflate, br","connection":"keep-alive","content-length":"3","content-type":"text/plain","host":"cloud.xxxxx.cn","referer":"http://cloud.xxxxx.cn/test/notice/text?content=78","user-agent":"PostmanRuntime/7.29.0","x-client-proto":"https","x-client-proto-ver":"HTTP/1.1","x-daa-tunnel":"hop_count=1","x-forwarded-for":"180.162.40.148, 122.228.230.43","x-forwarded-proto":"https","x-nws-log-uuid":"44f25ec2-d1bb-4d67-8417-88d5cd159262","x-real-ip":"122.228.230.43","x-stgw-time":"1653359076.243","x-tencent-ua":"Qcloud"},"httpMethod":"GET","isBase64Encoded":true,"multiValueHeaders":{"accept":["*/*"],"accept-encoding":["gzip, deflate, br"],"connection":["keep-alive"],"content-length":["3"],"content-type":["text/plain"],"host":["cloud.xxxxx.cn"],"referer":["http://cloud.xxxxx.cn/test/notice/text?content=78"],"user-agent":["PostmanRuntime/7.29.0"],"x-client-proto":["https"],"x-client-proto-ver":["HTTP/1.1"],"x-daa-tunnel":["hop_count=1"],"x-forwarded-for":["180.162.40.148, 122.228.230.43"],"x-forwarded-proto":["https"],"x-nws-log-uuid":["44f25ec2-d1bb-4d67-8417-88d5cd159262"],"x-real-ip":["122.228.230.43"],"x-stgw-time":["1653359076.243"],"x-tencent-ua":["Qcloud"]},"path":"/text","queryStringParameters":{"content":"78"},"requestContext":{"appId":"xxxxxx","envId":"cloud-xxxxxx","requestId":"3de8b674caf1f4311ebd09ffafdc0807","uin":"xxxxxx"}}`

	data2 := `{"body":"Nzg5","headers":{"accept":"*/*","accept-encoding":"gzip, deflate, br","connection":"keep-alive","content-length":"3","content-type":"text/plain","host":"cloud.xxxxx.cn","referer":"http://cloud.xxxxx.cn/test/notice/text?content=78\u0026content=788","user-agent":"PostmanRuntime/7.29.0","x-client-proto":"https","x-client-proto-ver":"HTTP/1.1","x-daa-tunnel":"hop_count=1","x-forwarded-for":"180.162.40.148, 122.228.230.43","x-forwarded-proto":"https","x-nws-log-uuid":"f4f020e5-8526-4eb1-98bb-a72068e05ac1","x-real-ip":"122.228.230.43","x-stgw-time":"1653359262.824","x-tencent-ua":"Qcloud"},"httpMethod":"GET","isBase64Encoded":true,"multiValueHeaders":{"accept":["*/*"],"accept-encoding":["gzip, deflate, br"],"connection":["keep-alive"],"content-length":["3"],"content-type":["text/plain"],"host":["cloud.xxxxx.cn"],"referer":["http://cloud.xxxxx.cn/test/notice/text?content=78\u0026content=788"],"user-agent":["PostmanRuntime/7.29.0"],"x-client-proto":["https"],"x-client-proto-ver":["HTTP/1.1"],"x-daa-tunnel":["hop_count=1"],"x-forwarded-for":["180.162.40.148, 122.228.230.43"],"x-forwarded-proto":["https"],"x-nws-log-uuid":["f4f020e5-8526-4eb1-98bb-a72068e05ac1"],"x-real-ip":["122.228.230.43"],"x-stgw-time":["1653359262.824"],"x-tencent-ua":["Qcloud"]},"path":"/text","queryStringParameters":{"content":["78","788"]},"requestContext":{"appId":"xxxxxx","envId":"cloud-xxxxxx","requestId":"22b839526aafa4ee22d9215dc09b15d8","uin":"xxxxxx"}}`

	dataSet := []string{data1, data2}

	for _, data := range dataSet {

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
		fmt.Println(request.QueryString)
		fmt.Println("-----")
	}

}
