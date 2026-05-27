# Claude Code Package Layout

Target-native artifact layout:

```text
<skill-id>/
  SKILL.md
  references/
  scripts/
  assets/
```

Generated AgentX output:

```text
.agentx/output/capabilities/<id>/targets/claude-code/<skill-id>/SKILL.md
```

Do not generate Cursor rules, Codex-only metadata, or AgentX-specific intermediate formats as final Claude Code artifacts.
