# GitHub Copilot Creator Baseline

Baseline: Copilot coding agent skill package.

Source:

- `https://docs.github.com/en/copilot/how-tos/agents/copilot-coding-agent/extending-copilot-coding-agent-with-skills`

Minimum package:

```text
<skill-id>/
  SKILL.md
```

Repository skills live under `.github/skills/`. Personal skills live under `~/.copilot/skills/`.

AgentX must not treat repository custom instructions as equivalent to skills; use them only when the requested output is instructions rather than a skill package.
