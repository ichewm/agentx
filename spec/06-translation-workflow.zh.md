# 06. 翻译工作流

## 目的

翻译是 AgentX 最有差异化的主流程。

AgentX 必须把一个 runtime 的源能力转成其他 runtime 的 target-native artifacts，并明确转换损耗。

## 输入

输入类型：

- 用户请求。
- 现有 `SKILL.md`。
- Claude Code plugin。
- Codex skill。
- GitHub Copilot skill。
- Cursor rule。
- OpenClaw skill pack。
- Hermes skill。
- MCP server 或 prompt pack。
- 文档和示例。

默认源输入优先级是路径、仓库 URL、目录。粘贴内容也被完整支持，不能被视为降级输入模式。工作流不能要求用户在多轮对话里反复复制粘贴。

如果源材料很大，AI 必须创建 intake summary，并保存到：

```text
.agentx/output/capabilities/<id>/sources/source-summary.md
```

AI 还必须在 capability brief 中记录 source provenance。

## Source Intake Modes

### 路径或文件夹

当用户提供本地路径或文件夹时，AI 必须检查文件树，识别候选能力文件，并总结发现。

如果文件本地可读，AI 不能要求用户粘贴文件。

### GitHub 或 Web URL

当用户提供 URL 时，如果工具允许，AI 必须获取或浏览它，记录 source URL，并总结相关文件或文档。

如果 URL 无法访问，AI 必须请求用户提供可下载压缩包、本地 clone 路径，或最小必要文件。

### 粘贴内容

当用户粘贴内容时，AI 必须把它保存为 intake entry。

如果内容是片段，AI 必须在生成最终产物前询问是否还有后续片段。

### 多轮

如果用户在多条消息中提供源材料，每条消息都必须成为一个带短标签的 intake entry；如果有时间戳，也必须记录。

AI 必须维护滚动 source summary，并避免要求用户重复复制粘贴。

## 步骤

1. 识别源 runtime。
2. 清点源文件。
3. 提取 capability map。
4. 创建或更新 capability brief。
5. 阅读 target profiles。
6. 产出 target plans。
7. 生成 target artifacts。
8. 运行审查。
9. 修订产物。
10. 产出 install/export plans。

## Target Generation Rule

始终优先使用目标 runtime 的原生能力格式。

例子：

- Codex：如果支持，使用 Agent Skills / `SKILL.md` package。
- Claude Code：按官方指导生成 skill 或 plugin layout。
- Copilot：使用支持的 skill directories 或官方 skill tooling。
- Cursor：生成 `.cursor/rules/*.mdc`，不要伪造成 skill。
- MCP：生成 tools、prompts 或 resources，不要伪造成 skill。

## Official Creator Baselines

当 target 有官方 creator skill 或指导时，AgentX 必须以它为生成基线。

AgentX 可以增加更严格的审查和跨 target 检查，但不能在有官方约定时自己发明 target conventions。

任何偏离官方 creator baseline 的地方都必须遵守 `spec/03-meta-skills.md` 中的 baseline adjustment protocol。翻译不能把 benchmark 驱动的变更静默当作官方 target 行为应用。

## Conversion Loss Report

每次翻译都必须包含：

```markdown
## Preserved

## Adapted

## Degraded

## Dropped

## Manual Setup Required

## Risks
```

除非目标 runtime 支持等价 primitives，否则不能声称无损转换。
