# Platform

中台服务,包含许多公共性服务

## Services:
* 校验服务: auth  token校验服务，支持JWT token校验
* 统计服务: analytics BI日志服务，支持thinkingData，MixPanel，localFile，ClickHouse等方式
* 好友服务: buddy 支持加好友/删除好友/拉黑/好友列表等功能
* 聊天服务: chat  支持私聊/群聊/频道聊天，暂时不支持离线消息
* 背包服务: knapsack 抽象背包功能，支持Item{ID, Count, Type,Expire}
* 排行榜服务: leaderboard 更新排行榜，获取排行榜，支持多种排行榜类型
* 邮件服务: mail  发送附件邮件，支持模板，过期时间，附件等功能
* 匹配服务: matchmaking  分配房间 (TODO refactor with openMatch)
* 组队服务: party  组队/踢人/解散队伍等功能
* 玩家信息服务: profile 创建/更新玩家基本信息
* 房间服务: room (TODO 帧同步房间)

## 目录结构 ：

工程目录结构参考[project-layout](https://github.com/golang-standards/project-layout)

## 运行：

* 你可以运行`cmd/platform/main.go`来启动所有服务， 你也可以参考`cmd/platform/main.go`来自定义组装服务。
* 你可以运行`cmd/{service}/main.go`来启动单个服务

## 容器化

```shell
# fix {appname} to service name
docker buildx build -t {appname}.registry.com:latest --build-arg APP_NAME={appname} -f ./build/package/docker/Dockerfile .  --push
```

## 服务安全:

* JWT token: https://www.okta.com/identity-101/what-is-token-based-authentication/
* mTLS: https://www.cloudflare.com/zh-cn/learning/access-management/what-is-mutual-tls/

## 服务类型：

每个服务下都可以创建两种类型的服务：

* private service: 不会校验JWT token, 会走mTLS认证，适用于内部服务调用
* public service: 会校验JWT token, 适用于外客户端调用