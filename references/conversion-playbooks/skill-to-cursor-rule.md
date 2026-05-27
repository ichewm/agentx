# Skill-to-Cursor-Rule Conversion Playbook

Use this playbook when converting a skill-like source into Cursor project rules.

## Steps

1. Extract the skill's trigger conditions, constraints, and required behavior.
2. Decide whether the output is one `.mdc` rule or multiple focused rules.
3. Generate `.cursor/rules/<rule-id>.mdc` under the target output directory.
4. Remove skill-only installation language.
5. Convert bundled references into rule instructions only when they are necessary for Cursor routing or behavior.
6. Mark scripts, tools, commands, and connectors as `manual`, `unsupported`, or `scoped-out` unless Cursor rule support is verified.
7. Produce conversion loss report and runtime benchmark plan.

## Guardrails

- Do not generate `SKILL.md` as the final Cursor artifact.
- Do not claim Cursor will execute scripts from rules.
- Keep project-specific paths explicit.
