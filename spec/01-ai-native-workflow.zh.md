# 01. AI 原生工作流

## 主要用户体验

主要入口是和已经加载 AgentX 的 AI runtime 对话。

AgentX 支持两个主要用户流程：

1. 从 0 创建一个新能力。
2. 把现有能力转换成等价或近似的目标产物。

用户可以用类似命令的自然语言请求启动任一流程。

创建示例：

```text
使用 AgentX。把下方内容编译成 Codex 和 Claude Code 可用的能力。

<用户提供的需求、文档、示例或说明>
```

用户不需要在一轮里提供所有材料。AgentX 必须支持多轮输入。

转换示例：

```text
使用 AgentX。把这个 Claude Code plugin 转成 Codex 和 Cursor 可用的产物。
尽量保留行为，并告诉我哪些不能转换。
```

AgentX meta-skills 指导当前 AI 完成输入、综合、审查和交付。

## 交互协议

AI 不能要求用户反复粘贴同一份源材料。

对于源材料，AI 必须接受：

- 直接粘贴的内容。
- 本地路径。
- GitHub 或 web URL。
- 包含 skill/plugin/rule pack 的目录。
- 多份文档或示例。

当源是路径或 URL，并且工具允许访问时，AI 必须直接读取。如果无法直接访问，AI 必须只向用户索要最小缺失输入。

AI 可以在创建或转换期间提出追问，但问题必须有针对性。必须避免泛泛的大问卷。

AI 必须把工作记录保存在 `.agentx/output/capabilities/<id>/` 下，这样后续对话可以继续，不需要重新粘贴源材料。

## 多轮输入

创建和转换都是 intake session，不是一次性命令。

当用户开始一个能力任务时，AI 必须创建或选择 capability id，并维护：

```text
.agentx/output/capabilities/<id>/
  intake.md
  sources/
  open-questions.md
  brief.md
```

用户可以继续输入：

```text
增加这个需求...
这是另一份文档...
把这个例子当作反例...
Cursor 后续也要支持，但这次先不输出...
```

对于每个新输入，AI 必须：

1. 把材料追加或总结到 `intake.md`。
2. 把大段粘贴内容或读取到的文件保存到 `sources/`。
3. 更新 `open-questions.md`，但除非下一步输入被阻塞，否则不要立刻要求用户回答。
4. 只有当意图足够清楚时才更新 `brief.md`。
5. 报告本轮改变了什么，还有什么未解决。
6. 主动询问用户是否还要继续输入。

继续输入提示必须明确且简短：

```text
是否还要继续输入？回复 "no"、"done"、"不" 或 "不需要" 表示结束输入并继续。
```

只要 intake 还没结束，AI 必须主要负责收集和总结。除非用户明确要求先生成草稿，否则 AI 不能开始深入质疑、生成 target artifacts 或制定安装计划。Intake 期间记录的问题必须留到结束输入后的决策阶段再集中提出，除非这个问题是理解下一份源材料所必需的。

当用户结束 intake 后，AI 必须：

1. 总结已积累的意图和源材料。
2. 提出需要用户决定的具体问题。
3. 对需要决策的地方给出明确方案选项。
4. 等用户决定后，再生成最终目标产物。

AI 不能生成最终目标产物，除非：

- 用户结束 intake 并确认生成。
- 用户要求现在先生成 early draft。

如果信息不完整，AI 必须产出 partial brief 和一小组具体缺失项。

## 主流程

1. Ingest
   - 阅读用户意图、文档、现有 skill、plugin、rule、repo 或示例。
   - 识别源 runtime 和源产物类型。

2. Map
   - 建立 capability map。
   - 识别 skills、commands、tools、connectors、subagents、references、assets、scripts、rules 和安装假设。

3. Brief
   - 产出平台无关的 capability brief。
   - 记录目标、使用条件、输入、步骤、边界、示例和转换说明。

4. Target Plan
   - 阅读 target profiles。
   - 判断输出是 direct、adapted、degraded、unsupported 还是 manual。

5. Generate
   - 按 target creator baseline 和官方约定生成产物。
   - 优先使用目标 runtime 原生格式。

6. Review
   - 做语义审查、可移植性审查、安全审计和 benchmark 设计。
   - 修订产物，直到问题解决或被明确接受。

7. Delivery Plan
   - 产出 install/export plan。
   - 展示源产物、目标位置、方法、变更、备份、验证和回滚。

8. Install or Export
   - 有官方安装器就使用官方安装器。
   - 只有确定性文件交付和验证才使用 thin helper。

## 创建流程

创建从用户请求和源材料包开始，不从传统 `agentx new` 命令开始。源材料可以为空。

```text
使用 AgentX。基于这些 migration 规范，为 Codex 和 Claude Code 构建数据库 migration 安全能力。
```

当前 AI 必须：

- 如果需求不完整，就提出有针对性的问题。
- 把源笔记和材料保存到 `.agentx/output/capabilities/<id>/sources/`。
- 跨轮维护 `intake.md`。
- 起草 capability brief。
- 生成目标产物。
- 运行审查。
- 产出安装计划。

## 翻译流程

翻译从已有能力开始。

```text
使用 AgentX。把这个 OpenClaw skill pack 翻译成 Claude Code 和 Codex 产物。
```

当前 AI 必须：

- 接受路径、URL、仓库、压缩包或粘贴内容作为源。
- 识别源格式。
- 把源材料复制或总结到 `.agentx/output/capabilities/<id>/sources/`。
- 如果源材料分多轮粘贴，保留每次输入为 intake entry，并避免重复索要同一内容。
- 优先保留意图，而不是保留文件形状。
- 使用官方 target creator guidance。
- 报告转换损耗。
- 不假装保留了目标 runtime 不支持的功能。

## 运行时选择

当目标 runtime 已经有原生发现机制时，AgentX 不负责在最终用户任务中选择 skill。

Codex、Claude Code、Copilot、OpenClaw、Hermes Agent 等 runtime 必须自己决定何时使用能力。AgentX 只准备产物。

只有目标 runtime 没有原生 skill/rule discovery 时，才允许 fallback selection，并且必须明确标注这是 fallback。
