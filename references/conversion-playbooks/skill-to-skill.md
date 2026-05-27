# Skill-to-Skill Conversion Playbook

Use this playbook when converting one `SKILL.md`-style package to another `SKILL.md`-style target such as Codex, Claude Code, Copilot coding agent, OpenClaw, or Hermes Agent.

## Steps

1. Read the source `SKILL.md` frontmatter and body.
2. Inventory bundled `references/`, `scripts/`, and `assets/`.
3. Extract intent into `.agentx/output/capabilities/<id>/brief.md`.
4. Read the target profile and creator baseline.
5. Generate a target-native `SKILL.md` first, preserving official required metadata.
6. Copy only relevant bundled resources.
7. Produce conversion loss report.
8. Run semantic, portability, safety, unknown-resolution, and benchmark reviews.

## Guardrails

- Do not preserve source folder shape when it conflicts with target baseline.
- Do not claim tools, commands, or connectors are available unless the target profile verifies them.
- Keep long details in references and keep the target `SKILL.md` concise.
