---
name: "EasyMonitor 本地启动跑不起来"
description: "本地想直接运行 easymonitor（docker-compose up）却失败或 webapp 持续报错时查看。首查方向：是否先编译二进制、build.sh 的 GOARCH 架构、docker-compose 缺少 redis 服务、redis 地址硬编码。"
---

# 排查：EasyMonitor 本地启动跑不起来

## 当前状态

有效（基于当前代码）。

## 问题

想「直接运行」项目时，常见三类失败：

- 直接 `docker-compose up` 起不来，或 `ubuntu`/`node0` 容器报 `exec format error`
- webapp 起来后持续刷 Redis 连接错误
- webapp 连 MySQL 首次失败

## 排查路径

1. **是否先编译二进制**：启动分两步——先 `sh build.sh` 生成 `alerterserver`、`webapp` 到 `./program/`，再 `docker-compose up`。跳过 build 直接 up，`ubuntu` 容器（跑 `./alerterserver`）和 `node0` 容器（`start_node_exporter.sh` 后台拉 `webapp`）会找不到/跑不了二进制。
2. **架构是否匹配**：`build.sh` 写死 `CGO_ENABLED=0 GOOS=linux GOARCH=arm64`。在 x86/amd64 机器上产物无法执行，表现为 `exec format error`。→ 改 `GOARCH=amd64`（或本机架构）。
3. **Redis 是否可达**：`webapp/main.go` 硬编码 `Addr: "192.168.2.6:6379"`，而 `docker-compose.yaml` **没有 redis 服务**。webapp 内有每 3 秒读写 Redis 的 goroutine，连不上会持续报错。→ 换成可达的 Redis 地址，或在 compose 补一个 redis 服务。
4. **MySQL 启动顺序**：webapp 连 `mydb:3306`（root/1234567，库 test）。mysql 首次初始化慢，webapp 可能先连失败，需重试或等待。

## 根因

项目是演示/测试性质，启动配置里存在若干本地不可直接跑通的硬编码/缺口：build 架构写死 arm64、compose 缺 redis 服务、redis/mysql 连接为硬编码测试值、二进制需先本地编译再编排。

## 验证方式

- `sh build.sh` 后确认 `./program/` 下有 `webapp`、`alerterserver` 且架构与运行机器一致
- `docker-compose up -d` 后各容器无 `exec format error`、无 restart 抖动
- 访问 Grafana `:3000`、Prometheus `:9090`、webapp 业务端口 `:8080`，指标端口 `:8090/metrics` 有数据

## 涉及模块

- webapp（示例应用入口、硬编码 redis/mysql）
- alerterserver（`ubuntu` 容器运行的二进制）
- docker-compose.yaml / build.sh / program/start_node_exporter.sh（编排与启动入口）

## AI 注意事项

- 当前沙箱环境为只读，无法真正 `go build` / `docker-compose up`，只能给出启动步骤与坑点，不能声称已运行。
- redis 硬编码在 webapp 模块文档里已提及是「测试值」，但「compose 无 redis 服务导致启动即报错」是跨文件缺口，容易漏判。
- 端口区分：webapp 业务 8080，指标 8090，node-exporter 9100，alerterserver 16060。

## 修订记录

- 2026-07-10 新增：首次记录本地启动的三类失败与首查顺序（用户「直接运行我看看」触发，代码证据来自 docker-compose.yaml、build.sh、webapp/main.go、program/start_node_exporter.sh）。
