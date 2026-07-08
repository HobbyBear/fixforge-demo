---
name: "长事务检测"
description: "自动检测和记录超过阈值（8000ms）的数据库事务。用户讨论事务超时、锁等待、连接池耗尽或事务性能问题时会提及。与慢查询检测容易混淆——长事务关注 Begin→Commit 整体耗时，慢查询关注单条 SQL。"
---

# 概念：长事务检测

## 是什么

通过包装 database/sql 的 driver.Tx 接口（infra/sqlhooks.go 的 DriveTx），在 Commit() 和 Rollback() 时自动计算事务从 Begin 到结束的总耗时。当事务耗时超过 8000ms 阈值时，记录结构化长事务日志，包含耗时、调用栈信息，帮助定位长时间未提交或回滚的事务。

## 为什么有

长事务是数据库性能杀手：长时间持有锁会导致连接池耗尽、其他事务阻塞、主从延迟增加。在中小团队中，开发者可能无意中在事务内做了 RPC 调用或复杂计算。通过自动检测和记录调用栈，可以快速定位长事务的来源。

## 用户侧表现

- 长事务日志通过 ELK 管道进入 Elasticsearch
- 在 Kibana 中可按 MetricType="longTx" 搜索
- 日志包含完整调用栈（Stack 字段），可直接定位到发起事务的业务代码行

## 系统侧数据流

1. 业务代码调用 BeginTx() → Conn.BeginTx() 包装返回 DriveTx（记录 start 时间）
2. 业务代码调用 Commit() 或 Rollback() → DriveTx.Commit()/Rollback() 计算耗时
3. 耗时 > 8000ms → logrus 输出结构化日志（Cost、MetricType="longTx"、Stack 调用栈）
4. 日志被 filebeat 采集 → Logstash → Elasticsearch → Kibana / alerterserver

## 核心规则

- 长事务阈值固定 8000ms
- 调用栈通过 infra/stack.go 的 callers() 捕获，记录在日志的 Stack 字段中
- 长事务检测仅记录日志，不主动中断事务——即使检测到长事务，Commit/Rollback 仍正常执行
- DriveTx 包装对业务代码透明：业务代码操作的是 driver.Tx 接口，感知不到包装层

## 易混淆概念

- **不是慢查询检测**：慢查询检测单条 SQL 的耗时；长事务检测整个事务生命周期
- **不是数据库锁超时**：长事务检测的是事务持续时间，不是锁等待时间；两者相关但不同
- **不是连接池泄漏检测**：长事务可能指示连接池问题，但检测本身不监控连接池状态
