---
name: "infra 监控基础设施"
description: "easymonitor 核心监控埋点库。用于定位 SQL/Redis/HTTP 的指标采集入口、驱动包装机制和改动影响面，涉及慢查询、长事务、告警策略等业务场景。"
keywords:
- "infra"
- "监控"
- "埋点"
- "指标"
- "metrics"
- "钩子"
- "hook"
- "慢查询"
- "infra 监控基础设施"
---

# infra 监控基础设施

## 业务描述

infra 模块是 easymonitor 的核心监控埋点库，承担所有业务服务的指标采集、日志记录和告警触发职责。它通过 MySQL 驱动包装、Redis 客户端钩子、HTTP 中间件三种渠道，以对业务代码零侵入的方式采集运行指标，并通过 Prometheus 暴露，最终被 Grafana/Prometheus/ELK 消费。适用于中小企业的生产环境监控。

## 关键代码

- `infra/sqlhooks.go:Wrap()` — MySQL 驱动包装入口，注入 Before/After/OnError 钩子以采集每条 SQL 的执行耗时和调用次数
- `infra/sqlhooks.go:DriveTx.Commit()/Rollback()` — 长事务检测：事务耗时 > 8000ms 时自动记录慢事务日志和调用栈
- `infra/hooks.go:Before()/After()` — HookDb 实现，驱动慢查询检测（默认阈值 1000ms）和表级多表查询告警
- `infra/redis.go:AddRedisHook()` — Redis 客户端钩子注册入口，基于 key 前缀白名单过滤监控范围
- `infra/redis.go:AddMonitorKey()` — 注册需要监控的 Redis key 前缀，只有匹配的 key 才被采集
- `infra/middleware.go:MetricMiddleware()` — HTTP 中间件，自动采集每个 HTTP 请求的处理耗时和状态码
- `infra/monitor.go:metricMonitor` — Prometheus 指标注册与暴露，包括 server_handle_seconds、client_handle_seconds 等四类指标
- `infra/sqlparser.go:parseTable()` — SQL 解析器，从 SQL 文本提取表名和操作类型
- `infra/stack.go:callers()` — 调用栈捕获，用于长事务日志中定位调用来源

## 常见修改场景

- **调整慢查询阈值**：在 `infra/hooks.go` 中修改 `After()` 方法的耗时判断常量
- **调整长事务阈值**：在 `infra/sqlhooks.go:DriveTx.Commit()/Rollback()` 中修改 8000ms 常量
- **新增监控指标类型**：在 `infra/monitor.go` 中新增 Prometheus HistogramVec/CounterVec，并在对应的钩子函数中调用 Record 方法
- **新增数据库驱动支持**：仿照 `infra/hooks.go` 的 `registerHookDriver()` 注册新驱动，实现 Hooks 接口
- **修改 Redis 监控 key 过滤规则**：调用 `RedisMonitor.AddMonitorKey()` 或在 `infra/redis.go:matchKey()` 中修改匹配逻辑
- **新增告警策略**：在 Before/After/OnError 钩子中添加自定义告警逻辑

## AI 注意事项

- infra 模块通过 driver.Wrap() 实现透明的数据库驱动包装，修改驱动包装链时注意不要破坏 database/sql 的 driver 接口契约
- Prometheus 指标通过 MetricsReg 全局注册表管理，所有指标在 init() 中注册；新增指标需同时更新 init() 和对应的 Record 方法
- Redis 监控基于 key 前缀白名单（monitorKeys），不在白名单内的 key 不会被采集指标——排查 Redis 指标缺失时先检查 AddMonitorKey 调用
- 慢查询和长事务通过 logrus 结构化日志输出（Fields 带 Cost/MetricType/Stack），而非直接告警——日志由 filebeat 采集到 ELK 后再触发告警
- SQL 解析器（sqlparser.go）有分表名归一化回调（SetFixName），调用方需注册此回调才能正确聚合分表指标
- webapp 和 alerterserver 各自独立编译为二进制文件，infra 是它们共享的库——修改 infra 会影响所有引用它的服务
