# 02. 仓库结构

## 顶层结构

原始 source tree 已经删除。仓库必须保持很小，并以 spec 和 meta-skills 为中心。

```text
agentx/
  AGENTS.md
  CLAUDE.md
  spec/
    *.md
    *.zh.md
  skills/
    agentx-workbench/
    agentx-capability-architect/
    agentx-capability-translator/
    agentx-capability-reviewer/
    agentx-safety-auditor/
    agentx-benchmark-designer/
    agentx-install-planner/
  targets/
    codex/
    claude-code/
    copilot/
    cursor/
    openclaw/
    hermes/
  models/
  references/
    standards/
    examples/
    conversion-playbooks/
  helper/
    README.md
```

旧的 `src/`、legacy `spec/*.md` 和 Zig build 文件已经删除。不要重建它们。不要把业务型应用源码树引入这个仓库。

## 生成输出工作区

AgentX 生成的工作内容必须放在 AgentX 仓库的 `.agentx/output/` 下。

原因：用户是在加载这个 AgentX 仓库之后，才开始能力创建或转换对话。因此 AgentX 仓库是 workbench 状态持有者。外部仓库、文件夹、URL 或粘贴文档是 source；目标 runtime 或目标项目是 install/export destination。

```text
.agentx/
  output/
    capabilities/
      <capability-id>/
        intake.md
        open-questions.md
        brief.md
        sources/
        targets/
          codex/
            <skill-id>/
          claude-code/
            <skill-id>/
          cursor/
        reviews/
        install/
        lock.json
  backups/
  logs/
```

`.agentx/` 默认必须被忽略。只有当用户明确要求提交某个特定生成产物时，才允许提交生成工作。Install plan 必须提醒用户 `.agentx/` 是否已忽略。

如果用户明确希望把生成能力工作区放到另一个项目里，必须显式提出，并记录在 install/export plan 中。

## 权威文件

英文文件是 AI 执行的权威依据：

- `AGENTS.md`
- `CLAUDE.md`
- `spec/*.md`
- `skills/*/SKILL.md`
- `targets/*/*.md`
- `models/*/*.md`
- `references/**/*.md`

以 `.zh.md` 结尾的中文文件只供人类审阅。

英文 spec 变更必须同步到对应 `.zh.md` 文件。Skills、targets、model profiles 和 references 不要求中文翻译。

## 能力产物结构

每个生成的能力必须按 capability id 分组。

```text
.agentx/output/capabilities/<id>/
  intake.md
  open-questions.md
  brief.md
  sources/
  targets/
    <target-id>/
      <target-native-package-root>/
  reviews/
    semantic-review.md
    portability-review.md
    safety-review.md
    benchmark-plan.md
    runtime-benchmark.<target-id>.md
    unknown-resolution.md
    baseline-deviations.md
  install/
    <target-id>.plan.md
  lock.json
```

目标产物必须已经可以复制、安装或导出，不需要 AI 再重新推导文件布局。当某个 target 的原生格式是 skill package directory 时，生成输出必须在 target id 下包含这个 package directory，例如 `.agentx/output/capabilities/<id>/targets/codex/<skill-id>/SKILL.md`。

## 旧代码

旧的 Zig 控制平面实现已经删除。

新的工作流不能依赖 `.agents/registry.yaml`、generated instruction files、`approve`、`sync`、`guard`、hook 控制平面，或把 GLM 当成 target runtime。
