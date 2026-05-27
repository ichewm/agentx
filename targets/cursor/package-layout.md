# Cursor Package Layout

Target-native artifact layout:

```text
.cursor/
  rules/
    <rule-id>.mdc
```

Generated AgentX output:

```text
.agentx/output/capabilities/<id>/targets/cursor/.cursor/rules/<rule-id>.mdc
```

Do not generate `SKILL.md` as a final Cursor artifact unless a future Cursor target profile verifies native skill support.
