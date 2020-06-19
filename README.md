# scf-apigw-wrap

腾讯云云函数工具，把 api 网关的请求封装成 http 请求模拟调用本地 http 服务，在把 http 的返回内容转化为 APIGatewayResponse

# 使用

```bash
go get -u github.com/TMaize/scf-apigw-wrap@v1.0.3
```

```go
import wrap "github.com/TMaize/scf-apigw-wrap"
```

```go
func Event(req events.APIGatewayRequest) (interface{}, error) {
  // 实际路径 /{req.Context.Stage}/{req.Path}
  // 传入的路径 /{req.Path}
  // 网关配置的前缀为 req.Context.Path
  // 根据需求转化自己http服务的请求路径
  innerUrl := "/" + req.Path[len(req.Context.Path):]
  return Wrap.wrap(req, innerUrl, Engine), nil
}
```
