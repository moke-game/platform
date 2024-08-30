# Platform

中台服务,包含许多公共服务

## Services:

* Auth: 用户认证服务,支持JWT token校验
* Analytics: BI日志服务，支持thinkingData，MixPanel，localFile，ClickHouse等方式
* Buddy: 好友服务，支持加好友/删除好友/拉黑/好友列表等功能
* Chat: 聊天服务，支持私聊/群聊/频道聊天，暂时不支持离线消息
* Knapsack: 背包服务，支持Item{ID, Count, Type,Expire}
* Leaderboard: 排行榜服务，更新排行榜，获取排行榜，支持多种排行榜类型
* Mail: 邮件服务，发送附件邮件，支持模板，过期时间，附件等功能
* Matchmaking: 匹配服务，分配房间 基于[open-match](https://open-match.dev/site/)+ [agones](https://agones.dev/site/)实现
* Party: 组队服务，组队/踢人/解散队伍等功能
* Profile: 玩家基本信息服务，创建/更新玩家基本信息
* Room: 房间管理服务，创建/加入/离开/销毁房间，房间内提供通用 lockstep 帧同步服务

## 目录结构规范 ：

工程目录结构参考[project-layout](https://github.com/golang-standards/project-layout)

## 运行：

* 你可以运行`cmd/platform/main.go`来启动所有服务， 你也可以参考`cmd/platform/main.go`来自定义组装服务。
* 你可以运行`cmd/{service-name}/service/main.go`来启动单个服务,例如：`cmd/auth/service/main.go` 来启动auth服务

## 容器化

```shell
# fix {appname} to service name
docker buildx build -t {appname}.registry.com:latest --build-arg APP_NAME={appname} -f ./build/package/docker/Dockerfile .  --push
```

## 如何测试？

### Integration Test

* build your interactive client:
   ```shell
     go build -o client.exe ./cmd/{game-name}/client/main.go 
   ```
* run your interactive client:
    ```shell
     # help
     ./{game-name}.exe help
    ```
  tips: http client use Postman to connect `localhost:8081`.

### Load Test

* install [k6](https://grafana.com/docs/k6/latest/get-started/installation/)
* run k6 load test
   ``` shell
    # fix the game-name to your game name
    k6 run ./tests/{game-name}/{game-name}.js
  ```

## 服务安全:

* JWT token: https://www.okta.com/identity-101/what-is-token-based-authentication/
* mTLS: https://www.cloudflare.com/zh-cn/learning/access-management/what-is-mutual-tls/

## 服务类型：

每个服务包含两种类型的服务：

* private service: 不会校验JWT token, 会走mTLS认证，适用于内部服务调用例如：gm, admin，或者其他内部服务调用。
* public service: 会校验JWT token, 会走mTLS认证 适用于其他外部服务调用。
