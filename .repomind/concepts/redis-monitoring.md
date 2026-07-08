---
name: "Redis 键监控"
description: "基于 key 前缀白名单的 Redis 操作监控。用户讨论 Redis 性能、缓存命中率、key 设计或 Redis 异常排查时会提及。与 MySQL 监控容易混淆——Redis 监控有白名单过滤机制，MySQL 监控是全量采集。"
---

# 概念：Redis 键监控

## 是什么

通过 go-redis 的 WrapProcess/WrapProcessPipeline 钩子，对 Redis 客户端操作进行透明拦截，采集每次 Redis 命令的执行耗时和调用次数。监控范围通过 key 前缀白名单（AddMonitorKey）控制，只有匹配白名单前缀的 key 才会被采集指标和记录日志。

## 为什么有

Redis 作为缓存层，其性能和可用性直接影响服务质量。但 Redis 命令非常高频，全量采集会产生海量指标。通过 key 前缀白名单机制，团队可以只监控核心业务的 Redis key，平衡可观测性和资源开销。

## 用户侧表现

- Redis 操作指标在 Prometheus 的 client_handle_seconds/client_handle_total 指标中，type=redis
- Grafana Redis 仪表盘展示各 key 的调用次数和耗时分布
- Redis 错误日志（rediserrlog）通过 ELK 管道可搜索

## 系统侧数据流

1. 业务代码调用 Redis 命令 → go-redis 执行 → WrapProcess 拦截
2. matchKey() 检查命令中的 key 是否匹配白名单 → 匹配则记录指标
3. cacheWrapper() 记录耗时指标到 Prometheus，错误时记录 rediserrlog

## 核心规则

- 必须调用 AddMonitorKey 注册 key 前缀，否则该 key 的所有操作都不会被监控
- key 匹配是前缀包含匹配（strings.Contains），不是精确匹配
- Redis Nil 错误（key 不存在）不算异常，不记录错误日志
- key 超过 100 字符会被 truncateKey() 截断

## 易混淆概念

- **不是全量 Redis 监控**：仅监控白名单 key；要监控更多 key 需显式添加
- **不是 Redis Slowlog**：EasyMonitor 在客户端侧采集，Redis 服务端的 Slowlog 是另一套机制
