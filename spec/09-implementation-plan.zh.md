# 09. 实施计划

## Phase 0：规格重置

状态：已在当前仓库完成。

已交付：

- 英文 specs。
- 中文审阅翻译。
- 根目录 agent instructions 指向英文 specs。
- 已删除旧控制平面代码和旧 specs。

## Phase 1：Meta-Skill Pack

交付初始 AgentX meta-skills：

- `agentx-workbench`
- `agentx-capability-architect`
- `agentx-capability-translator`
- `agentx-capability-reviewer`
- `agentx-safety-auditor`
- `agentx-benchmark-designer`
- `agentx-install-planner`

每个 meta-skill 都必须能作为标准 skill package 被 Codex 和 Claude Code 使用。

## Phase 2：Target Profiles

交付初始 target profiles：

- `codex`
- `claude-code`
- `copilot`
- `cursor`
- `openclaw`
- `hermes`

每个 target profile 必须引用官方行为；无法确认的内容必须标为 unknown。

交付与目标相关模型的初始 model profiles。Model profiles 只影响表达和审查敏感度。

## Phase 3：Workbench Output Convention

交付：

- `.agentx/output/capabilities/<id>/` 布局。
- 路径、URL、仓库和粘贴内容的 source intake convention。
- 使用 `intake.md` 和 `open-questions.md` 的多轮输入 convention。
- 每次输入后的继续输入提示和 intake 结束信号。
- Capability brief template。
- Review templates。
- Install plan template。
- Lock file schema。

## Phase 4：Thin Helper

Helper 语言：Go。

原因：

- 跨平台单二进制。
- 标准库适合文件、JSON、hash、HTTP、压缩包和子进程。
- 比 Zig 更适合快速写安装器。
- 比 shell 更安全、更可测试。
- 比 TypeScript 少 Node runtime 依赖。

Shell 必须只用于 bootstrap scripts。

Zig 不是 AgentX 的同级实现选择。以后只有出现具体技术理由时，才重新考虑。

## Phase 5：Helper Commands

初始 helper commands：

```text
agentx detect <target>
agentx plan install <capability> --target <target>
agentx install <capability> --target <target>
agentx export <capability> --target <target>
agentx verify <capability> --target <target>
agentx rollback <capability> --target <target>
agentx list
```

任何命令都不得替代 AI 语义审查。

## Phase 6：Benchmarks

交付 benchmark templates，并且每个 sample capability 至少有一个 benchmark plan。

生成能力必须做 runtime benchmark 自动化。如果某个 target runtime 暂时无法自动化，AgentX 必须记录该 gap，并要求人工 runtime benchmark transcript 后才能最终交付。

## 当前产品非目标

- 没有明确 GLM runtime 原生能力安装机制之前，不做 GLM target adapter。
- 不做 registry single source of truth。
- 不做 all-target `sync`。
- 不做 hook/automation 控制平面。
- 不做 marketplace publishing。
- 不做 GUI。
