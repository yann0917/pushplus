# pushplus

> go语言版 [pushplus推送加](https://www.pushplus.plus) 
> 
> [官方文档地址](https://www.pushplus.plus/doc/#%E5%BC%95%E8%A8%80)

## TODO

* [x] 消息接口
  * [x] 消息发送接口
* [x] 开放接口
  * [x] 获取 AccessKey
* [x] 消息接口
  * [x] 消息列表
  * [x] 查询消息发送结果
* [x] 用户接口
  * [x] 获取 token
  * [x] 个人资料详情
  * [x] 获取解封剩余时间
* [x] 群组接口
  * [x] 群组列表
  * [x] 获取我加入的群组详情
  * [x] 新增群组
  * [x] 获取群组二维码
  * [x] 退出群组
* [x] 群组用户接口
  * [x] 获取群组内用户
  * [x] 删除群组内用户
* [x] 渠道配置接口
  * [x] 获取 webHook 列表
  * [x] 获取 webhook 详情
  * [x] 新增 webhook
  * [x] 修改 webhook
* [x] 功能设置接口
  * [x] 获取默认发送渠道
  * [x] 修改默认发送渠道
  * [x] 修改接收消息限制

## 使用说明

1. 如果只要发送消息功能, 填写`token`即可，
2. 需要通过accessKey调用消息、用户、群组、渠道配置和功能设置接口时，需填写 `secretKey`。
3. `DebugLog` true 表示开启调试日志。
4. 具体使用方法参考测试用例  `*_test.go`。

```go
// 发送消息
token := "xxx"
HttpClient := NewClient(token, DebugLog(true))
msg, err := HttpClient.SendMsg("标题", "内容")
if err != nil {
    return
}
fmt.Println(msg)

// 开放接口举例
secretKey := "xxx"
Client := NewClient(token, SecretKey(secretKey), DebugLog(true))

result,err := Client.GetUserInfo()
if err != nil {
    return
}
fmt.Println(result)
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/yann0917/pushplus.svg)](https://starchart.cc/yann0917/pushplus)

## License

[MIT](./LICENSE) © yann0917

---
