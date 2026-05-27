# GitHub Copilot Package Layout

Target-native artifact layout:

```text
<skill-id>/
  SKILL.md
  references/
  scripts/
  assets/
```

Repository install layout:

```text
.github/skills/<skill-id>/SKILL.md
```

Generated AgentX output:

```text
.agentx/output/capabilities/<id>/targets/copilot/.github/skills/<skill-id>/SKILL.md
```
