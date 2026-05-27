# 07. 审查、安全和 Benchmark

## 审查层次

AgentX 使用 AI 语义审查加确定性 helper 检查。

AI 审查回答：

- 这个产物是否表达了预期能力？
- 目标 runtime 能否正确路由到它？
- 目标产物是否过宽或过窄？
- 翻译是否改变了语义？
- 是否存在不安全或隐藏指令？

Helper 检查回答：

- 必需文件是否存在？
- Frontmatter 是否可解析？
- 引用文件是否存在？
- 路径是否在允许目录中？
- Hash 是否匹配 lock file？

## 安全审计

每个导入或翻译的能力都必须审计：

- Prompt injection。
- Hidden instruction override。
- Secret 访问。
- Credential 访问。
- 危险 shell。
- 网络外传。
- 混淆脚本。
- 过度 tool permissions。
- 未锁版本的外部依赖。

高风险发现必须在安装前要求人工确认。

## Benchmark 目的

必须做 benchmark，因为 skill 质量依赖模型和 runtime。

AgentX benchmark 不只是模型分数，而是检查目标 runtime 是否会触发并遵守能力。

## Benchmark 层级

Level 1：静态 benchmark

- 正例。
- 反例。
- 混淆任务。
- 必须行为 checklist。
- 禁止行为 checklist。

Level 2：Runtime benchmark

- 手动或通过支持的自动化运行目标 runtime。
- 观察 runtime 是否选择能力。
- 记录行为是否遵守产物。

Level 3：跨 target benchmark

- 比较 Codex、Claude Code、Copilot、Cursor、OpenClaw、Hermes 或其他 targets。
- 记录 target-specific drift。

## Benchmark 文件

```text
.agentx/output/capabilities/<id>/reviews/benchmark-plan.md
```

Benchmark plan 必须人可读，并能被自动化复用。

## 必需 Benchmark 产物

每个生成能力都必须包含：

```text
.agentx/output/capabilities/<id>/reviews/benchmark-plan.md
.agentx/output/capabilities/<id>/reviews/runtime-benchmark.md
```

生成能力必须做 runtime 自动化 benchmark。如果某个 runtime 暂时无法自动化，AgentX 必须在 `runtime-benchmark.md` 中记录 blocking gap，或者在 `runtime-benchmark.md` 中记录人工 runtime benchmark transcript。

最终交付前，`runtime-benchmark.md` 必须记录成功的自动化 runtime benchmark，或人工 runtime benchmark transcript。如果它只记录 blocking gap，则该 target artifact 不是 target-ready。

Benchmark 自动化不能声称成功，除非它确实运行过目标 runtime 或有记录的官方测试 harness。
