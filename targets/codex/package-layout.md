# Codex Package Layout

Target-native artifact layout:

```text
<skill-id>/
  SKILL.md
  references/
  scripts/
  assets/
```

Only `SKILL.md` is required. Other directories are bundled resources included only when needed.

Generated AgentX output:

```text
.agentx/output/capabilities/<id>/targets/codex/<skill-id>/SKILL.md
```

Do not emit a fake registry file or generated monolithic instruction file as the final Codex artifact.
