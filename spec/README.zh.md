# AgentX 规格说明

这个目录是 AgentX 的权威规格目录。

AgentX 不是本地 registry 控制平面，而是一个 AI 原生的能力工作台：一组加载到 Agent Runtime 里的元能力，加上一个很薄的确定性 helper，用来做文件交付、备份、回滚和校验。

英文文件是 AI 和实现工作的权威依据。以 `.zh.md` 结尾的中文文件只供人类审阅。

## 阅读顺序

1. [00-positioning.zh.md](00-positioning.zh.md)
2. [01-ai-native-workflow.zh.md](01-ai-native-workflow.zh.md)
3. [02-repository-layout.zh.md](02-repository-layout.zh.md)
4. [03-meta-skills.zh.md](03-meta-skills.zh.md)
5. [04-target-profiles.zh.md](04-target-profiles.zh.md)
6. [05-capability-brief.zh.md](05-capability-brief.zh.md)
7. [06-translation-workflow.zh.md](06-translation-workflow.zh.md)
8. [07-review-safety-benchmark.zh.md](07-review-safety-benchmark.zh.md)
9. [08-installation-delivery.zh.md](08-installation-delivery.zh.md)
10. [09-implementation-plan.zh.md](09-implementation-plan.zh.md)

## 核心原则

Runtime 自己选择 skill。AgentX 帮助当前 AI 创建、翻译、审查、打包并交付能力。
