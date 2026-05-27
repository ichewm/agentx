# 04. Target Profiles

## 目的

Target profiles 把平台知识外置。AgentX 必须通过增加 profile 文件和 playbook 来支持新 runtime，而不是修改核心工作流。

## Target 目录

每个 target 目录必须包含全部五个文件：

```text
targets/<target-id>/
  profile.md
  creator-baseline.md
  package-layout.md
  install.md
  review-checklist.md
```

## profile.md

描述 runtime 能力。

必需主题：

- Runtime 名称和作用域。
- 是否支持 native skills。
- 是否支持 rules。
- 是否支持 prompts、commands、tools、connectors、hooks、subagents。
- Runtime 如何路由能力。
- 是否由模型在运行时选择能力。
- 支持的 scope：project、user、organization、plugin。
- 已知 refresh 或 restart 要求。
- 已知不支持的功能。

## creator-baseline.md

总结该 target 的官方能力创建指导。

AgentX 必须把它作为 target-specific compiler reference。

## package-layout.md

定义该 target 的生成产物结构。

例子：

- `SKILL.md` package。
- Claude Code plugin layout。
- Cursor `.mdc` rule。
- MCP server tool/prompt/resource layout。

## install.md

定义交付方式。

允许的方法类型：

- `official-cli`
- `official-ui`
- `path-copy`
- `path-symlink`
- `export-only`
- `manual`

官方安装器优先于 helper copy。

## Unknown Values

Unknown 只允许出现在 target 和 model profile 的研究状态中，而且必须显式写出。

当某个事实尚未验证时，使用 `Unknown`。Unknown 用来防止编造平台能力声明。Unknown 不会阻止 profile 被加载，但会影响生成：

- AgentX 可以在存在 Unknown 字段时创建内部草稿。
- 对于标记为 Unknown 的功能，AgentX 不能声称最终 target compatibility。
- 当 install method 或 destination 是 Unknown 时，AgentX 不能执行安装。
- AgentX 不能把依赖未解决 Unknown 的行为放进最终生成产物。
- 最终交付前，AgentX 必须研究该 Unknown、询问用户、标记为 unsupported，或标记为 manual。
- 如果 Unknown 影响 routing、package layout、installation、safety 或 benchmark validity，该 target 的最终交付必须阻塞，直到 Unknown 被解决或被明确排除在范围外。

## Unknown Resolution Gate

每个 target artifact 在标记为 target-ready 前，都必须通过 Unknown resolution gate。

AI 必须产出：

```text
.agentx/output/capabilities/<id>/reviews/unknown-resolution.md
```

该文件必须列出 target profile、model profile、source capability、conversion notes 和 generated draft 中所有可能影响目标产物的 Unknown。

每一项都必须使用以下 resolution 之一：

- `verified`：通过引用官方文档、检查本地 runtime 行为，或用户提供的权威输入解决。
- `unsupported`：目标不支持该功能，并且最终产物排除了它。
- `manual`：该功能需要人工设置，并且已写入 install/export plan。
- `scoped-out`：该功能不在用户请求的本次产物范围内。

任何条目都不能保留为 `Unknown`、`TBD`、`TODO` 或空白。

最终 target artifact 目录中的 generated target files、install plans、lock records 或 final review summaries 不能包含 `Unknown`、`TBD`、`TODO` 这类未解决占位词。这些词只允许存在于 profile research files、历史 draft notes，或 `unknown-resolution.md` 中用于指称原始未解决 source fact 且已经分配 resolved status 的条目里。

如果确定性 helper 在 generated target files、install plans、lock records 或 final review summaries 中发现未解决占位词，必须让 finalization 失败。Helper 不能仅仅因为 `unknown-resolution.md` 提到了一个已有非 Unknown resolution 的原始 Unknown source item 就失败。Helper 不能判断语义 Unknown 是否被正确解决；这仍然是 AI review 的责任。

## review-checklist.md

定义 target-specific 审查要求。

例子：

- Routing description 质量。
- Script 权限假设。
- Progressive disclosure 结构。
- Target-specific frontmatter。
- Unsupported feature 处理。

## Model Profiles

模型不是安装目标。

Model profiles 从第一版实现开始就是完整 AgentX 产品的一部分。

GLM、GPT、Claude、Hermes 或其他模型可以有 model profiles，但这些 profile 只影响表达和审查敏感度。安装永远指向 runtime。

Model profiles 位于：

```text
models/<model-id>/
  profile.md
  expression-guidance.md
  review-sensitivity.md
```

Model profiles 不能定义安装路径。
