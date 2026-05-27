# 03. Meta-Skills

## 目的

AgentX 主要以 meta-skills 交付。这些 skills 教当前 AI 如何创建、翻译、审查、打包和安装其他能力。

这些是 AgentX 自己用于生产和翻译其他能力的能力。它们必须拆成独立 skills，这样每个任务只加载相关的 meta-skill。

AgentX 必须包含一个轻量 `agentx-workbench` 入口 skill 用于发现和路由。它不能重复包含所有 meta-skill 的完整内容。它必须把 AI 指向正确的专门 meta-skill。

用户生成的能力是另一回事：它们必须遵循目标 runtime 的原生打包规则。AgentX 不允许为了自己方便而发明最终产物格式。一个生成能力可以是一个 skill、多个 skills、一个 rule 文件、一个 plugin，或按 target profile 生成的混合产物。

## 必需 Meta-Skills

### agentx-workbench

把用户路由到正确的 AgentX meta-skill。

职责：

- 判断用户是在创建能力、翻译已有能力、审查产物、规划安装，还是设计 benchmark。
- 加载或指向该任务需要的专门 meta-skill。
- 保持自身内容简短。
- 避免重复包含专门 meta-skills 的完整说明。

### agentx-capability-architect

把用户意图和源资料转成 capability brief。

职责：

- 澄清目标。
- 识别目标用户和 runtime。
- 提取触发条件和使用场景。
- 定义输入、流程、边界和输出。
- 区分稳定逻辑和 target-specific 表达。

### agentx-capability-translator

在不同 runtime 之间翻译已有能力。

职责：

- 识别源 runtime 和产物类型。
- 提取 capability map。
- 阅读 target profiles。
- 生成 target-native artifacts。
- 产出 conversion loss report。

### agentx-capability-reviewer

审查能力质量和语义一致性。

职责：

- 检查产物是否表达了 brief 的意图。
- 检查多个 target 版本是否仍然语义一致。
- 检查 description 是否支持正确触发。
- 检查 references 和 scripts 是否符合 progressive disclosure。

### agentx-safety-auditor

审查能力安全性。

职责：

- 检测 prompt injection。
- 检测隐藏的 instruction override。
- 检测危险 shell。
- 检测 secret 或 credential 访问。
- 检测网络外传模式。
- 检测过度 target 权限。
- 对高风险行为要求人工确认。

### agentx-benchmark-designer

设计轻量 benchmark 和回归用例。

职责：

- 创建正向触发示例。
- 创建反向示例。
- 创建混淆任务。
- 定义期望行为。
- 定义 target-specific 验证说明。

### agentx-install-planner

创建 install/export plan。

职责：

- 识别是否有官方安装器。
- 选择安装方式。
- 列出变更路径。
- 定义备份路径。
- 定义回滚方式。
- 定义验证步骤。
- 安装前请求确认。

## Baseline Creator Skills

只要存在官方或 runtime 原生 creator skill、官方创建指南，AgentX 必须把它作为 baseline。

例子：

- Codex skill creator。
- Claude Code skill/plugin 创建指导。
- Hermes skill management pattern。
- OpenClaw Skill Workshop pattern。

当审查或 benchmark 发现弱点时，AgentX 不能盲目复制官方 creator 输出。官方 baseline 定义 target-native 的最低要求。

Baseline 调整不是自动工具决策。它是受 AgentX meta-skills 约束的 AI 原生工作流。当前 AI 可以为了可移植性、安全性、跨 target 一致性或模型敏感度，提出有记录、由 benchmark 驱动的调整，但不能把这些调整静默当成官方 baseline 行为直接应用。

Baseline 调整只能按以下协议进行：

1. 读取官方 baseline，并识别必需的 target-native 文件、metadata、命名规则、路由规则和安装假设。
2. 先生成符合 baseline 的产物。
3. 对该产物运行语义审查、可移植性审查、安全审计和 benchmark 设计。
4. 如果 review 或 benchmark 证据显示存在弱点，写出 baseline adjustment proposal，包含问题、证据、精确变更、预期收益、风险和 target-native 合法性。
5. 如果调整会改变生成的 target artifacts、runtime behavior、safety posture、installation semantics 或 cross-target meaning，必须把 proposal 交给用户确认。
6. 只有在用户确认后，并且调整保留官方必需结构、必需 metadata、package layout 和 installation semantics 时，才能应用。
7. 把偏离记录到 `.agentx/output/capabilities/<id>/reviews/baseline-deviations.md`。
8. 调整后重新运行相关 review 和 benchmark 检查。

除非官方文档提供可接受替代方案，否则 AgentX 不能删除、重命名或削弱官方必需元素。任何偏离官方 baseline 的地方都必须记录在目标产物审查中。

确定性 helper 不能判断语义 baseline 调整是否正确。Helper 只能验证必需文件、路径、hash、备份和安装交付状态这类确定性事实。

## Meta-Skill 输出纪律

每个 meta-skill 都必须输出：

- 检查了什么。
- 生成或修改了什么。
- 风险或未解决问题。
- 需要审阅的文件。
- 是否可以安全安装。
