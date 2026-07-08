---
name: "webapp 示例应用"
description: "easymonitor 的示例 Web 应用，展示如何集成 infra 模块做监控埋点。用于定位监控 SDK 的接入方式和典型使用模式。"
keywords:
- "webapp"
- "示例"
- "demo"
- "接入"
- "HTTP"
- "webapp 示例应用"
---

# webapp 示例应用

## 业务描述

webapp 是 easymonitor 的示例 Web 应用，展示如何使用 infra 模块进行 MySQL、Redis、HTTP 三层监控埋点。它演示了：数据库连接池初始化与驱动包装、Redis 客户端钩子注册、HTTP 中间件集成、分表名归一化回调等典型接入模式。同时包含一个模拟的业务协程持续执行 SQL 和 Redis 操作以产生监控数据。

## 关键代码

- `webapp/main.go:main()` — 应用入口，初始化 SqlMonitor（连接池、分表名归一化）、RedisMonitor（钩子、key 白名单）、HTTP 中间件
- `webapp/main.go:handleRequest()` — HTTP 请求处理器，经 MetricMiddleware 包装后自动采集指标
- `webapp/main.go` 中 `SetFixName()` 调用 — 演示分表名归一化：将 `t_user_*` 映射为 `t_user`
- `webapp/main.go` 中 `AddMonitorKey("name")` — 演示 Redis key 前缀白名单配置

## 常见修改场景

- **新增业务服务接入监控**：参照 webapp/main.go 中的初始化流程复制 SqlMonitor/RedisMonitor/MetricMiddleware 三部曲
- **修改分表规则**：在 SetFixName 回调中调整表名映射逻辑
- **新增 Redis 监控 key**：调用 AddMonitorKey 添加新的 key 前缀

## AI 注意事项

- webapp 是示例代码，不是生产服务——它的数据库密码、Redis 地址都是硬编码的测试值
- webapp 和 alerterserver 共用 infra 模块，但各自独立编译部署
- webapp 的监控指标通过 8090 端口（/metrics）暴露，业务端口是 8080——不要混淆
