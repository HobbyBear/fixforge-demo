# Graph Report - .  (2026-07-08)

## Corpus Check
- 81 files · ~51,740 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 230 nodes · 356 edges · 22 communities (15 shown, 7 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 40 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Client Monitor Hooks & Redis|Client Monitor Hooks & Redis]]
- [[_COMMUNITY_Database Driver & Connection Layer|Database Driver & Connection Layer]]
- [[_COMMUNITY_OpenSpec & RepoMind Skills|OpenSpec & RepoMind Skills]]
- [[_COMMUNITY_App Entry & HTTP Middleware|App Entry & HTTP Middleware]]
- [[_COMMUNITY_Transaction Driver & Logging|Transaction Driver & Logging]]
- [[_COMMUNITY_Stack Frame & Call Tracing|Stack Frame & Call Tracing]]
- [[_COMMUNITY_SQL Parser & Query Analysis|SQL Parser & Query Analysis]]
- [[_COMMUNITY_Config & OPSX Commands|Config & OPSX Commands]]
- [[_COMMUNITY_OpenSpec Skill Definitions|OpenSpec Skill Definitions]]
- [[_COMMUNITY_DB Driver Init & Hook Registry|DB Driver Init & Hook Registry]]
- [[_COMMUNITY_Docker Compose & ELK Stack|Docker Compose & ELK Stack]]
- [[_COMMUNITY_RepoMind Skill Definitions|RepoMind Skill Definitions]]
- [[_COMMUNITY_Prometheus Metrics Export|Prometheus Metrics Export]]
- [[_COMMUNITY_Grafana Dashboards & SME Monitoring|Grafana Dashboards & SME Monitoring]]
- [[_COMMUNITY_OpenSpec Artifact Model|OpenSpec Artifact Model]]
- [[_COMMUNITY_Build Script (singleton)|Build Script (singleton)]]
- [[_COMMUNITY_Grafana Docker (singleton)|Grafana Docker (singleton)]]
- [[_COMMUNITY_MySQL Docker (singleton)|MySQL Docker (singleton)]]
- [[_COMMUNITY_Redis Monitor (singleton)|Redis Monitor (singleton)]]
- [[_COMMUNITY_SQL Hooks (singleton)|SQL Hooks (singleton)]]
- [[_COMMUNITY_HTTP Handler (singleton)|HTTP Handler (singleton)]]

## God Nodes (most connected - your core abstractions)
1. `cacheWrapper()` - 9 edges
2. `DriveTx` - 9 edges
3. `Frame` - 9 edges
4. `OpenSpec Onboarding Skill` - 9 edges
5. `RepoMind Summary Skill` - 9 edges
6. `MetricMiddleware()` - 8 edges
7. `stack` - 8 edges
8. `openspec-onboard: Guided onboarding walkthrough for OpenSpec` - 8 edges
9. `HookDb` - 7 edges
10. `truncateKey()` - 7 edges

## Surprising Connections (you probably didn't know these)
- `OpenSpec Artifact-Driven Workflow Model` --semantically_similar_to--> `OpenSpec Change Lifecycle Pipeline`  [INFERRED] [semantically similar]
  openspec/config.yaml → .claude/commands/opsx/new.md
- `main()` --calls--> `MetricMiddleware()`  [INFERRED]
  webapp/main.go → infra/middleware.go
- `SMB-oriented monitoring system design` --rationale_for--> `MetricMonitor`  [EXTRACTED]
  ReadMe.md → infra/monitor.go
- `SMB-oriented monitoring system design` --rationale_for--> `Grafana App Dashboard`  [EXTRACTED]
  ReadMe.md → grafanadashbord/app_dashboard.json
- `RepoMind Graph Summary (fallback mode)` --references--> `OpenSpec Project Configuration (spec-driven)`  [EXTRACTED]
  .repomind/graph/summary.json → openspec/config.yaml

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **** — infra_monitor_metricmonitor, infra_middleware_metricmiddleware, infra_hooks_hookdb, infra_redis_redismonitor, grafanadashbord_app_dashboard, grafanadashbord_mysql_dashboard, grafanadashbord_redis_dashboard [EXTRACTED 1.00]
- **** — infra_sqlhooks_driver, infra_sqlhooks_conn, infra_sqlhooks_stmt, infra_sqlhooks_drivetx, infra_sqlhooks_hooks_interface, infra_sqlhooks_wrap [EXTRACTED 1.00]
- **** — easymonitor_sme_monitoring, infra_monitor_metricmonitor, grafanadashbord_system_dashboard, program_start_node_exporter, buildsh_buildscript [INFERRED 0.85]
- **ELK Logging Pipeline** — docker_compose_filebeat, docker_compose_logstash, docker_compose_elasticsearch, docker_compose_kibana, filebeat_config, logconf_api_input [EXTRACTED 1.00]
- **Prometheus-Grafana Monitoring Stack** — docker_compose_prometheus, docker_compose_grafana, docker_compose_mynode, prometheus_scrape_config [EXTRACTED 1.00]
- **OpenSpec OPSX Change Workflow Pipeline** — opsx_new, opsx_continue, opsx_propose, opsx_ff, opsx_apply, opsx_verify, opsx_archive, opsx_sync, opsx_bulk_archive, opsx_explore, opsx_onboard, openspec_artifact_model, opsx_change_lifecycle [EXTRACTED 1.00]
- **OpenSpec Change Lifecycle Pipeline** — openspec_new_change_skill_new_change, openspec_continue_change_skill_continue_change, openspec_propose_skill_propose_change, openspec_ff_change_skill_fast_forward, openspec_apply_change_skill_apply_change, openspec_verify_change_skill_verify_change, openspec_sync_specs_skill_sync_specs, openspec_archive_change_skill_archive_change, openspec_bulk_archive_change_skill_bulk_archive [INFERRED 0.85]
- **RepoMind Knowledge Lifecycle** — repomind_init_skill_repomind_init, repomind_prd_skill_repomind_prd, repomind_query_skill_repomind_query, repomind_summary_skill_repomind_summary, rules_repomind_repomind_rules [INFERRED 0.85]
- **RepoMind Query-Summary Mandatory Pipeline** — repomind_query_skill_repomind_query, repomind_summary_skill_repomind_summary, repomind_summary_skill_query_findings, rules_repomind_summary_gate, rules_repomind_frontmatter_routing [EXTRACTED 1.00]
- **OpenSpec Change Workflow Pipeline** — openspec_new_change_skill, openspec_continue_change_skill, openspec_apply_change_skill, openspec_archive_change_skill, openspec_propose_skill, openspec_ff_change_skill, openspec_explore_skill, openspec_verify_change_skill, openspec_sync_specs_skill, openspec_bulk_archive_change_skill [INFERRED]
- **RepoMind Knowledge Lifecycle** — repomind_init_skill, repomind_prd_skill, repomind_query_skill, repomind_summary_skill, repomind_knowledge_types, repomind_summary_gate [INFERRED]

## Communities (22 total, 7 thin omitted)

### Community 0 - "Client Monitor Hooks & Redis"
Cohesion: 0.10
Nodes (21): Client, Cmder, Transparent instrumentation via driver wrapping, Grafana MySQL Dashboard, Grafana Redis Dashboard, HookDb, Context, truncateKey() (+13 more)

### Community 1 - "Database Driver & Connection Layer"
Cohesion: 0.17
Nodes (19): Conn, Driver, ErrorHook, Hook, OnErrorer, Context, Hooks, handlerErr() (+11 more)

### Community 2 - "OpenSpec & RepoMind Skills"
Cohesion: 0.18
Nodes (25): OpenSpec Apply Change Skill, OpenSpec Archive Change Skill, OpenSpec Bulk Archive Change Skill, OpenSpec Continue Change Skill, OpenSpec Artifact, OpenSpec Explore Mode Skill, Explore as Stance Not Workflow, OpenSpec Fast-Forward Change Skill (+17 more)

### Community 3 - "App Entry & HTTP Middleware"
Cohesion: 0.10
Nodes (17): Build Script, Pluggable alert notification strategy, Shard table name normalization for metric aggregation, Grafana System Dashboard, Handler, start_node_exporter.sh script, Middleware, MetricMiddleware() (+9 more)

### Community 4 - "Transaction Driver & Logging"
Cohesion: 0.15
Nodes (13): Multi-table select query warning, Slow query detection (>1s) and long transaction detection (>8s), Const: Cost, MetricType, Stack, DriveTx, HookDb.After, HookDb.OnError, init(), client_handle_seconds histogram (+5 more)

### Community 5 - "Stack Frame & Call Tracing"
Cohesion: 0.33
Nodes (5): Frame, stack, funcname(), StackTrace, State

### Community 6 - "SQL Parser & Query Analysis"
Cohesion: 0.20
Nodes (11): HookDb.Before, TestParseSql(), SqlOp, getOp(), getTable(), getTables(), sqlMonitor, sqlMonitor.parseTable (+3 more)

### Community 7 - "Config & OPSX Commands"
Cohesion: 0.18
Nodes (14): OpenSpec Project Configuration (spec-driven), OPSX Apply Command, OPSX Archive Command, OPSX Bulk Archive Command, OPSX Continue Command, OPSX Explore Command, OPSX Fast Forward Command, OPSX New Command (+6 more)

### Community 8 - "OpenSpec Skill Definitions"
Cohesion: 0.29
Nodes (13): openspec-apply-change: Implement tasks from an OpenSpec change, openspec-archive-change: Archive a completed change, openspec-bulk-archive-change: Batch archive multiple changes, OpenSpec Change: a container for work holding proposal, specs, design, and tasks artifacts, openspec-continue-change: Continue working on a change by creating the next artifact, openspec-explore: Thinking partner for exploring ideas before/during a change, openspec-ff-change: Fast-forward through artifact creation, openspec-new-change: Start a new OpenSpec change (+5 more)

### Community 9 - "DB Driver Init & Hook Registry"
Cohesion: 0.29
Nodes (7): Conn, DbInfo, decorateMySQLConn(), sqlMonitor, sqlMonitor.InitHookDb, LocalDbClient, registerHookDriver()

### Community 10 - "Docker Compose & ELK Stack"
Cohesion: 0.24
Nodes (10): Alerterserver Service (ubuntu container), Elasticsearch Service, Filebeat Service, Kibana Service, Logstash Service, Node Exporter Service (mynode), Prometheus Service, Filebeat Configuration (+2 more)

### Community 11 - "RepoMind Skill Definitions"
Cohesion: 0.27
Nodes (10): repomind-init OpenAI agent configuration, repomind-init: Initialize RepoMind knowledge base with graphify rebuild, RepoMind Knowledge Types: concepts, modules, troubles, repomind-prd OpenAI agent configuration, repomind-prd: Extract business concepts from historical PRD documents, repomind-query OpenAI agent configuration, repomind-query: Route business/code questions through RepoMind knowledge before answering, Summary Gate: lightweight check before writing to RepoMind knowledge base (+2 more)

### Community 13 - "Grafana Dashboards & SME Monitoring"
Cohesion: 0.29
Nodes (7): SMB-oriented monitoring system design, Grafana App Dashboard, MetricMonitor, MetricMonitor.RecordServerHandlerSeconds, MetricMonitor.RegPrometheusClient, server_handle_seconds histogram, EasyMonitor Project

### Community 14 - "OpenSpec Artifact Model"
Cohesion: 0.67
Nodes (3): OpenSpec Artifact-Driven Workflow Model, OpenSpec Change Lifecycle Pipeline, Test Change Artifacts

## Knowledge Gaps
- **52 isolated node(s):** `build.sh script`, `sqlMonitor`, `Handler`, `T`, `Client` (+47 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **7 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `DriveTx` connect `Transaction Driver & Logging` to `Database Driver & Connection Layer`, `Stack Frame & Call Tracing`?**
  _High betweenness centrality (0.095) - this node is a cross-community bridge._
- **Why does `Wrap()` connect `Database Driver & Connection Layer` to `Client Monitor Hooks & Redis`, `DB Driver Init & Hook Registry`?**
  _High betweenness centrality (0.085) - this node is a cross-community bridge._
- **Why does `registerHookDriver()` connect `DB Driver Init & Hook Registry` to `Database Driver & Connection Layer`?**
  _High betweenness centrality (0.079) - this node is a cross-community bridge._
- **Are the 2 inferred relationships involving `RepoMind Summary Skill` (e.g. with `OpenSpec Archive Change Skill` and `OpenSpec Verify Change Skill`) actually correct?**
  _`RepoMind Summary Skill` has 2 INFERRED edges - model-reasoned connections that need verification._
- **What connects `build.sh script`, `sqlMonitor`, `Handler` to the rest of the system?**
  _57 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Client Monitor Hooks & Redis` be split into smaller, more focused modules?**
  _Cohesion score 0.10256410256410256 - nodes in this community are weakly interconnected._
- **Should `App Entry & HTTP Middleware` be split into smaller, more focused modules?**
  _Cohesion score 0.10276679841897234 - nodes in this community are weakly interconnected._