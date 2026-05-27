# Codex Creator Baseline

Baseline: Agent Skills / `SKILL.md` package.

Sources:

- OpenAI Help: `https://help.openai.com/en/articles/20001066-skills-in-chatgpt`
- Local Codex skill-creator guidance: `/Users/hewm/.codex/skills/.system/skill-creator/SKILL.md`

Minimum package:

```text
<skill-id>/
  SKILL.md
```

`SKILL.md` must contain YAML frontmatter with `name` and `description`. The description must clearly state the trigger conditions because runtime selection depends on metadata.

Allowed bundled resources:

- `scripts/`
- `references/`
- `assets/`

AgentX must keep `SKILL.md` concise and move detailed references into bundled files when needed.
