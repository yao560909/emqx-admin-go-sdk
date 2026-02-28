# emqx-admin-go-sdk
Go client library for EMQX Admin API.

## 安装

```go
go get github.com/yao560909/emqx-admin-go-sdk
```
## API Client

开发者在调用 API 前，需要先创建一个 API Client，然后才可以基于 API Client 发起 API 调用，Client对象线程安全，全局初始化一次即可。
### 创建API Client

```go
import sdk "github.com/yao560909/emqx-admin-go-sdk"
const (
    appId = "b36b87e3f5ba05d9"
    appSecret= "paP9BI3F7p9BjfcuxI89CI3T184rDxOxQf4tbdrHz9AEGKL"
    target="172.27.106.198:30239"
    scheme="http"
)
sdk.WithLogLevel(core.LogLevelInfo)
var client = client := sdk.NewClient(sdk.OpenSource, appId, appSecret, target, scheme)
```

每个配置选项的具体含义，如下表格：

<table>
  <thead align=left>
    <tr>
      <th>
        配置选项
      </th>
      <th>
        配置方式
      </th>
       <th>
        描述
      </th>
    </tr>
  </thead>
  <tbody align=left valign=top>
    <tr>
      <th>
        <code>LogLevel</code>
      </th>
      <td>
        <code>sdk.WithLogLevel(logLevel projectcore.LogLevel)</code>
      </td>
      <td>
设置 API Client 的日志输出级别(默认为 Info 级别)，枚举值如下：

- LogLevelDebug
- LogLevelInfo
- LogLevelWarn
- LogLevelError

</td>
</tr>
<tr>
      <th>
        <code>Logger</code>
      </th>
      <td>
        <code>sdk.WithLogger(logger projectcore.Logger)</code>
      </td>
      <td>
设置API Client的日志器，默认日志输出到标准输出。

开发者可通过实现下面的 Logger 接口，来设置自定义的日志器:

```go
type Logger interface {
Debug(context.Context, ...interface{})
Info(context.Context, ...interface{})
Warn(context.Context, ...interface{})
Error(context.Context, ...interface{})
}
```

</td>
</tr>
  </tbody>
</table>

## API调用

创建完毕 API Client，我们可以使用 ``Client.业务域.方法名称`` 来定位具体的 API 方法，然后对具体的 API 发起调用。

EMQX Admin开放的所有 API
列表(v5.4版本)，可点击[这里查看](https://docs.emqx.com/zh/emqx/v5.4/admin/api-docs.html)

### 基本用法

如下示例我们通过 client 调用EMQX 节点相关的 ListNodes
方法，获取EMQX节点列表：

``` go
import (
    "context"
    "fmt"
    "github.com/yao560909/emqx-admin-go-sdk/service/nodes"
    sdk "github.com/yao560909/emqx-admin-go-sdk"
)

const (
    appId = "b36b87e3f5ba05d9"
    appSecret= "paP9BI3F7p9BjfcuxI89CI3T184rDxOxQf4tbdrHz9AEGKL"
    target="172.27.106.198:30239"
    scheme="http"
)


func main() {
	ctx:=context.Background()
	client := sdk.NewClient(sdk.OpenSource, appId, appSecret, target, scheme)
	resp, err := client.Nodes.ListNodes(ctx, nodes.NewListNodesReqBuilder().Build())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	code:=resp.APIResp.StatusCode
	fmt.Sprintf("code:%d\n",code)
	header := resp.APIResp.Header
	for k, v := range header {
		fmt.Println("key:", k, "value:", v)
	}
	switch code {
	case 200:
		fmt.Printf("data:%v\n",resp.Data)
	case 400:
		fmt.Printf("codeError:%v\n",resp.CodeError)
	}
}
```

更多 API 调用示例：[./sample/demo.go](./sample/demo.go)
