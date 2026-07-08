---
name: "FixForge 测试项目"
description: "项目身份：本仓库是 FixForge 的 QA 测试沙箱，以 EasyMonitor 代码为基底，叠加 RepoMind 知识库管理和 OpenSpec 规约驱动开发工作流。用户讨论项目用途、仓库身份或工具配置时提及。与 EasyMonitor 监控方案本身是不同的概念层——FixForge 是容器，EasyMonitor 是内容。"
---

# 概念：FixForge 测试项目

## 是什么

本仓库（fixforge-qa-145-382）是一个 FixForge 平台的 QA 测试沙箱。它以 EasyMonitor（开源监控方案）的完整代码为测试基底，在此基础上叠加了两套工具能力的测试配置：

- **RepoMind**：业务知识库管理系统，通过 concepts / modules / troubles 三类知识文档和 graphify 代码图谱，支持 AI 辅助的代码问答、排查和方案设计
- **OpenSpec**：规约驱动的开发工作流，通过 proposal → specs → design → tasks 的制品序列管理变更

## 为什么有

FixForge 需要一套真实但不敏感的业务代码来验证 RepoMind、OpenSpec 等工具的集成效果。EasyMonitor 代码规模适中、结构清晰、有明确的业务语义（监控、埋点、告警），适合作为测试基底。本仓库不是 EasyMonitor 的官方副本，而是 FixForge 专用的测试沙箱。

## 核心事实

- 仓库名：fixforge-qa-145-382
- 基底项目：EasyMonitor（github.com/HobbyBear/performance-analyze）
- 测试能力：RepoMind 知识库管理 + OpenSpec 规约工作流
- 不用于生产：本仓库所有代码和数据仅用于工具测试验证

## 易混淆概念

- **不是 EasyMonitor 官方仓库**：EasyMonitor 原始作者是 HobbyBear，本仓库是其代码的测试副本，不向上游提交
- **不是生产环境**：这里是测试沙箱，所有配置、数据、知识库内容都是测试性质的
- **FixForge 不是被测试的代码**：被测试的是 EasyMonitor 代码；FixForge 是提供测试工具和环境的平台
