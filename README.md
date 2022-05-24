# scf-apigw-wrap

腾讯云云函数工具，无需 web 项目任何变动，就可以迁移到云函数中，并返回标准的 APIGatewayResponse

实现方式是通过 httptest 来调用 ServeHTTP，只要实现了`http.Handler`接口的框架都能使用，比如`gin.Engine`

## 注意

为了兼容 CloudBase, 升级到 `v1.1.1` 后代码不兼容

# 使用

```bash
go get -u github.com/TMaize/scf-apigw-wrap@v1.1.1
```

```go
import (
	gw "github.com/TMaize/scf-apigw-wrap"
	"github.com/TMaize/scf-apigw-wrap/model"
)

func main() {

	// 以gin框架为例
	server := gin.Default()

	// 接口的访问地址 => /api/hello
	server.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	cloudfunction.Start(func(event map[string]interface{}) (model.Response, error) {

		// 把访问的path转换为内部的path
		// https://service-xxxx-xxxx.sh.apigw.tencentcs.com/test/webapp/api/hello
		// prefix     => webapp 已知
		// event.Path => /webapp/api/hello
    request, _ := model.GetRequest(event)
		uri := strings.TrimPrefix(request.Path, "/webapp")

		// 调用 gin server
		resp := gw.Wrap(event, uri, server)
		return resp, nil
	})

}
```

# 使用 v1.0.4

```bash
go get -u github.com/TMaize/scf-apigw-wrap@v1.0.4
```

```go
import (
  gw "github.com/TMaize/scf-apigw-wrap"
  "github.com/tencentyun/scf-go-lib/events"
)

func main() {

	// 以gin框架为例
	server := gin.Default()

	// 接口的访问地址 => /api/hello
	server.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	cloudfunction.Start(func(event events.APIGatewayRequest) (events.APIGatewayResponse, error) {

		// 把访问的path转换为内部的path
		// https://service-xxxx-xxxx.sh.apigw.tencentcs.com/test/webapp/api/hello
		// prefix     => webapp 已知
		// event.Path => /webapp/api/hello
		uri := strings.TrimPrefix(event.Path, "/webapp")

		// 调用 gin server
		resp := gw.Wrap(event, uri, server)
		return resp, nil
	})

}
```
