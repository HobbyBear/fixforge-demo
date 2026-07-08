---
name: "EasyMonitor 监控体系"
description: "整体项目概念：面向中小企业的零侵入生产环境监控方案。用户讨论监控架构、指标设计、组件选型或接入流程时会提及。与具体某个指标的实现容易混淆——它描述的是整体方案而非单个指标。"
---

# 概念：EasyMonitor 监控体系

## 是什么

EasyMonitor 是一套面向中小企业的生产环境监控方案，通过 MySQL 驱动包装、Redis 客户端钩子、HTTP 中间件三种方式，对业务代码零侵入地采集 SQL 执行耗时、Redis 操作耗时、HTTP 请求指标，并通过 Prometheus + Grafana 可视化，Filebeat + ELK 做日志聚合和告警。

## 为什么有

生产环境需要对服务运行状态有可观测性，但中小团队通常缺乏自建监控体系的人力和经验。EasyMonitor 提供一套开箱即用的方案：业务代码只需引入 infra 包并做少量初始化配置，即可自动获得 SQL 慢查询、长事务、Redis 异常、HTTP 请求等关键指标的采集和可视化能力。

## 用户侧表现

- 开发者通过 Grafana 面板（app/mysql/redis/system 四类仪表盘）查看实时指标
- 当 SQL 查询超过慢查询阈值或事务超过长事务阈值时，结构化日志被 filebeat 采集到 Elasticsearch，最终可通过 Kibana 查询或在 alerterserver 触发告警
- 业务服务通过 8090 端口暴露 Prometheus 指标，Prometheus 定期拉取

## 系统侧数据流

1. 业务代码调用 database/sql 或 go-redis → infra 包装层拦截并记录耗时/次数
2. HTTP 请求经过 MetricMiddleware → 记录处理耗时和状态码
3. 所有指标写入 Prometheus MetricsReg → 通过 /metrics 端点暴露 → Prometheus 拉取 → Grafana 查询展示
4. 慢查询/长事务/Redis 错误通过 logrus 结构化日志输出 → filebeat 采集 → Logstash 处理 → Elasticsearch 存储 → Kibana 查询 / alerterserver 告警

## 核心规则

- MySQL 驱动必须通过 `infra.Wrap()` 包装才能被监控；直接使用原始驱动不会有任何指标
- Redis 监控基于 key 前缀白名单（`AddMonitorKey`），未注册的 key 不采集
- 慢查询默认阈值 1000ms，长事务默认阈值 8000ms
- 指标端口固定 8090，业务端口由各服务自行定义
- 监控体系采用 sidecar 模式：各监控组件（Prometheus、Grafana、ELK）与业务服务通过 docker-compose 统一编排

## 易混淆概念

- **不是 APM**：EasyMonitor 是应用级指标监控而非分布式链路追踪，不提供 trace ID 或跨服务调用链
- **不是 Prometheus**：Prometheus 是 EasyMonitor 使用的指标存储组件，EasyMonitor 是包含指标采集、日志聚合、告警、可视化的完整方案
- **不是单个指标**：慢查询检测、长事务检测、Redis 监控、HTTP 指标都是 EasyMonitor 体系下的子能力
