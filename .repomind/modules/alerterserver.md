---
name: "alerterserver 告警服务"
description: "easymonitor 的告警接收与分发服务。用于定位告警通知的接收端点、路由策略和扩展点，涉及告警管道和通知渠道接入。"
keywords:
- "告警"
- "alert"
- "alerterserver"
- "通知"
- "钉钉"
- "webhook"
- "alerterserver 告警服务"
---

# alerterserver 告警服务

## 业务描述

alerterserver 是 easymonitor 的告警接收与分发服务，监听 16060 端口，提供 `/alert_log` 和 `/alert_grafana` 两个 Webhook 端点。它接收来自 ELK（通过日志分析触发）和 Grafana（通过告警规则触发）的告警通知，当前以日志打印方式展示告警内容，设计意图是支持根据应用名称和告警内容匹配不同的通知策略（如钉钉群消息）。

## 关键代码

- `alerterserver/main.go:main()` — 服务入口，注册两个告警路由并启动 HTTP 服务
- `alerterserver/main.go` `/alert_log` 端点 — 接收来自日志分析的告警通知
- `alerterserver/main.go` `/alert_grafana` 端点 — 接收来自 Grafana 的告警通知

## 常见修改场景

- **新增告警通知渠道**（钉钉、企微等）：在告警处理函数中添加 HTTP 调用通知渠道的 Webhook
- **新增告警路由策略**：根据请求体中的应用名称、告警内容字段做分支处理
- **新增告警端点**：仿照现有路由注册新的 http.HandleFunc

## AI 注意事项

- alerterserver 当前仅打印告警日志到 stdout，没有实际的通知发送逻辑——TODO 注释说明了预期的扩展方向
- 告警通知渠道扩展是可插拔的：根据应用服务名称和告警内容匹配不同策略
- alerterserver 是独立编译的二进制文件，通过 docker-compose 与监控组件一起部署
