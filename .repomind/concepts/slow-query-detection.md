---
name: "SQL 慢查询检测"
description: "自动检测和记录超过阈值的 SQL 查询。用户讨论慢查询监控、数据库性能优化或 SQL 耗时告警时会提及。与长事务检测容易混淆——慢查询关注单条 SQL 耗时，长事务关注事务整体耗时。"
---

# 概念：SQL 慢查询检测

## 是什么

通过 MySQL 驱动包装层（infra/hooks.go 的 HookDb.After()），自动检测每条 SQL 的执行耗时，当超过阈值（默认 1000ms）时记录结构化慢查询日志，包含 SQL 文本、耗时、指标类型和调用栈信息。

## 为什么有

生产环境中慢 SQL 是数据库性能劣化的主要根因之一。通过驱动层透明埋点，开发者无需在业务代码中手动打点就能捕获所有 SQL 的执行耗时，及时发现性能瓶颈。

## 用户侧表现

- 慢查询日志通过 filebeat → Logstash → Elasticsearch 进入 ELK 管道
- 可在 Kibana 中按应用、耗时、表名搜索慢查询日志
- 可在 Grafana MySQL 仪表盘中查看 SQL 耗时分布

## 系统侧数据流

1. 业务代码执行 SQL → 经过 HookDb.Before() 记录开始时间
2. SQL 执行完毕 → HookDb.After() 计算耗时
3. 耗时超过阈值 → logrus 输出结构化日志（含 Cost、MetricType="slowQuery"、表名、操作类型）
4. 同时通过 MetricMonitor 记录 client_handle_seconds 指标至 Prometheus

## 核心规则

- 慢查询阈值默认 1000ms，是编译时常量
- 慢查询检测依赖 SQL 解析器（sqlparser.go）从 SQL 文本提取表名和操作类型
- 分表场景下，必须先通过 SetFixName 注册归一化回调，否则分表指标不聚合
- 多表查询（如 JOIN、子查询）会被单独标记

## 易混淆概念

- **不是长事务检测**：慢查询检测衡量单条 SQL 的耗时；长事务检测衡量整个事务（Begin→Commit/Rollback）的耗时
- **不是 MySQL 慢查询日志**：EasyMonitor 慢查询是应用层采集，不依赖 MySQL 服务端配置
