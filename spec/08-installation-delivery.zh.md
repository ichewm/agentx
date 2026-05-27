# 08. 安装和交付

## 定义

安装是把生成产物交付到 runtime 的原生消费路径。

AgentX 不把任何东西安装进模型。

## 交付原则

如果目标平台有官方安装器，优先使用官方安装器。

默认交付模式只产出 plan。AgentX 必须先生成 install/export plan 并等待。只有当用户明确要求继续执行时，AgentX 才能执行安装。

AgentX helper 只负责：

- 安装计划。
- 在没有官方安装器时复制文件。
- 除非后续明确实现 helper symlink support，否则 symlink delivery 只能作为 manual 或官方安装器计划。
- 备份。
- 回滚。
- 验证。
- Lock 记录。

## 产物位置

生成产物首先位于：

```text
.agentx/output/capabilities/<id>/targets/<target-id>/
```

AgentX 从这个位置安装或导出。

## Target-Ready Gate

AgentX 不能在 target artifact 通过 target-ready gate 前产出 install-ready plan。

必需的 target-ready evidence：

- `reviews/semantic-review.md` 存在，并且没有把产物标记为 blocked。
- `reviews/portability-review.md` 存在，并且没有把产物标记为 blocked。
- `reviews/safety-review.md` 存在，并且没有把产物标记为 blocked。
- `reviews/benchmark-plan.md` 存在。
- `reviews/runtime-benchmark.md` 存在，并且记录成功的自动化 runtime benchmark 或人工 runtime benchmark transcript。
- `reviews/unknown-resolution.md` 存在，并且没有未解决的 Unknown、TBD、TODO 或空白 resolution。
- 当产物偏离官方 creator baseline 时，`reviews/baseline-deviations.md` 必须存在。
- `.agentx/output/capabilities/<id>/targets/<target-id>/` 下的生成目标文件不能包含 `Unknown`、`TBD` 或 `TODO` 这类未解决占位词。

如果任何 target-ready evidence 缺失或被标记为 blocked，install planner 必须输出 `Blocked`，而不是 install-ready plan。

确定性 helper 可以通过检查必需文件、blocked 标记、占位词和 lock metadata 来执行这个 gate。但这些检查不能被当作语义正确性的证明。

## Install Plan

任何交付动作在修改文件前都必须先产出 plan。

```markdown
# Install Plan: <capability-id> -> <target-id>

## Source Artifact

## Target Runtime

## Method

## Destination

## Changes

## Backup

## Verification

## Rollback

## User Confirmation
```

## 方法

### official-cli

使用 runtime 官方 CLI。

AgentX 必须展示确切命令，并在执行前请求确认。

### path-copy

把生成文件复制到 runtime 路径。

要求：

- 备份已有目标。
- 复制文件。
- 验证布局。
- 写入 lock 记录。

### path-symlink

把生成产物软链到 runtime 路径。

只有 target runtime 安全支持 symlink 时才能使用。

当前 helper 不执行 symlink delivery。在 helper symlink support 被明确实现并 benchmark 之前，symlink plan 必须标记为 manual，或交给官方 target installer。

### export-only

把生成文件写入项目配置。

例子：

- 导出 Cursor rule 到 `.cursor/rules/*.mdc`。

### manual

只输出说明，不修改文件。

## 验证

验证必须只能证明文件交付成功。除非运行过 runtime benchmark，否则不能声称模型一定会触发能力。
