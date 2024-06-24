# Platform

中台服务,包含许多公共性服务

## Services:

* 统计服务: analytics
* 校验服务: auth
* 好友服务: buddy
* 聊天服务: chat
* 背包服务: knapsack
* 排行榜服务: leaderboard
* 邮件服务: mail
* 匹配服务: matchmaking
* 组队服务: party
* 玩家信息服务: profile

## 目录结构 ：

工程目录结构参考[project-layout](https://github.com/golang-standards/project-layout)

## 运行：

你可以运行`cmd/platform/main.go`来启动所有服务，当然你可以参考`cmd/platform/main.go`来自定义组装服务。

## 容器化

```shell
# fix {appname} to service name
docker buildx build -t {appname}.registry.com:latest --build-arg APP_NAME={appname} -f ./build/package/docker/Dockerfile .  --push
```

## 服务安全:

* JWT token: https://www.okta.com/identity-101/what-is-token-based-authentication/
* MTls: https://www.cloudflare.com/zh-cn/learning/access-management/what-is-mutual-tls/

## 服务类型：

每个服务下都可以创建两种类型的服务：

* private service: 不会校验JWT token, 会走mTLS认证，适用于内部服务调用
* public service: 会校验JWT token, 适用于外客户端调用