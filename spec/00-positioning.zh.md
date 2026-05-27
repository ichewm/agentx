# 00. 定位

## 产品定义

AgentX 是一个运行时无关的 Agent 能力工作台。

它帮助 AI Agent 创建、翻译、审查、打包并交付 skill、rule、prompt、tool、guard 等能力。

AgentX 不是模型，也不是 Codex、Claude Code、GitHub Copilot、Cursor、OpenClaw、Hermes Agent 的替代品。AgentX 是加载进这些 runtime 的元能力层。

## AgentX 是什么

AgentX 是：

- 一组用来构建其他 Agent 能力的 meta-skills。
- 一套描述目标 runtime 和其原生能力格式的参考资料。
- 一个把用户意图、文档、现有 skill/plugin 转成目标产物的工作流。
- 一个做语义一致性、可移植性、安全性、benchmark 准备的审查系统。
- 一个很薄的确定性 helper，用来做安装计划、文件交付、备份、回滚和验证。

## AgentX 不是什么

AgentX 不是：

- Codex、Claude Code、Copilot 等已有原生 skill discovery 的 runtime 的运行时 selector。
- 替代 Agent Skills / `SKILL.md` 的新通用标准。
- 把 skill 安装进模型的工具。
- 围绕 registry 文件构建的控制平面。
- 把所有能力编译进一个 generated instruction 文件的系统。
- 某个平台专属的插件管理器。

## 核心概念

逻辑是能力文档，runtime 是读取并执行它的 AI 平台。

因此 AgentX 首先是一组加载进现有 AI runtime 的能力。当前 AI 使用 AgentX meta-skills 去生产其他能力。

## 系统边界

创造性和语义性工作属于加载了 AgentX 的 AI：

- 理解用户意图。
- 阅读源资料。
- 提取 capability brief。
- 翻译现有 skill/plugin/rule。
- 生成目标产物。
- 审查语义一致性。
- 审计安全性和可移植性。
- 设计 benchmark 用例。

确定性工作属于 helper：

- 检测目标 runtime 路径。
- 复制或软链文件。
- 计算 hash。
- 写入 lock 记录。
- 创建备份。
- 回滚安装。
- 验证安装后的文件布局。

## 长期形态

AgentX 必须通过新增 target profile 和 playbook 来扩展，而不是修改核心代码。

支持一个新 runtime 必须通过增加：

- `targets/<runtime>/profile.md`
- `targets/<runtime>/creator-baseline.md`
- `targets/<runtime>/package-layout.md`
- `targets/<runtime>/install.md`
- `targets/<runtime>/review-checklist.md`

## 重写范围

原始控制平面实现已经删除。

仓库必须只保留 AI 原生工作台需要的内容：

- 规格。
- AgentX meta-skills。
- Target profiles。
- Reference documents。
- Thin helper source。
- `references/` 下的 reference examples 和 benchmark fixtures。

除非用户明确要求迁移桥接，否则不要为了兼容而重建旧文件。
