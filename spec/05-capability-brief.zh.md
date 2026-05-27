# 05. Capability Brief

## 目的

Capability brief 是生成 target-specific artifacts 之前使用的稳定中间文档。

它不是新的公开 skill 标准，而是 AgentX 的工作产物，用来帮助当前 AI 在多个 target 间保持意图一致。

## 位置

```text
.agentx/output/capabilities/<id>/brief.md
```

Brief 会随着输入逐步从这些内容构建：

```text
.agentx/output/capabilities/<id>/intake.md
.agentx/output/capabilities/<id>/sources/
.agentx/output/capabilities/<id>/open-questions.md
```

AI 必须在用户提供更多材料时增量更新 brief。

## 必需章节

```markdown
# Capability Brief: <id>

## Goal

## Source Materials

## Intended Runtimes

## When To Use

## When Not To Use

## Inputs To Inspect

## Procedure

## Hard Constraints

## Safety Notes

## Expected Output

## Examples

## Non-Goals

## Conversion Notes
```

## Capability Map

翻译任务里，brief 必须包含 capability map：

```text
skills:
commands:
tools:
connectors:
subagents:
rules:
scripts:
references:
assets:
install assumptions:
```

## Conversion Status

每个映射项都必须标注：

- `direct`：可以在目标 runtime 中保留。
- `adapted`：可以改写成 target-native 形式。
- `degraded`：只能以降级方式表示。
- `manual`：需要人工配置。
- `unsupported`：不能诚实地表示。

## 设计规则

Brief 必须保留意图，而不是保留源文件形状。

目标产物必须由 brief 和 target profiles 生成，而不是盲目做文件到文件转换。

## Partial Briefs

多轮输入期间，brief 可以是不完整的。

不完整 brief 必须包含：

```markdown
## Missing Information

## Assumptions

## Questions For User
```

只有当用户明确要求在信息缺失时也先生成 draft artifact，AI 才可以继续生成草稿。
