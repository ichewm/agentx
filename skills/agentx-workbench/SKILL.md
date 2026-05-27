---
name: agentx-workbench
description: Route AgentX capability creation, translation, review, benchmark, and install-planning requests to the correct specialized AgentX meta-skill.
---

# AgentX Workbench

Use this entry skill when the user asks to use AgentX, create a capability, translate an existing skill/plugin/rule, review a generated artifact, design benchmarks, or prepare an install/export plan.

## Routing Procedure

1. Identify the user's workflow:
   - create from zero
   - convert existing capability
   - review artifact
   - safety audit
   - benchmark design
   - install/export planning
2. Direct the AI to the specialized AgentX meta-skill for that workflow.
3. Keep the intake state under `.agentx/output/capabilities/<id>/`.
4. Do not duplicate the full instructions of specialized meta-skills.

## Mandatory Rules

- Use English files as authority: `spec/*.md`, `skills/*/SKILL.md`, `targets/*/*.md`, `models/*/*.md`, and `references/**/*.md`.
- Use `.zh.md` files only for human review.
- Create or reuse `.agentx/output/capabilities/<id>/` for every capability task.
- After every intake message, summarize what changed and ask:

```text
Do you want to add more input? Reply "no", "done", "不", or "不需要" to close intake and continue.
```

- Do not generate final target artifacts until intake is closed and the user confirms generation, unless the user explicitly asks for an early draft.
- If the user asks for installation, route to `agentx-install-planner`; do not execute installation directly.
- Do not execute external AI runtime CLIs for runtime benchmarks unless the user explicitly confirms that benchmark run.

## Workflow Map

- New capability from requirements: use `agentx-capability-architect`.
- Existing skill/plugin/rule conversion: use `agentx-capability-translator`.
- Semantic/portability review: use `agentx-capability-reviewer`.
- Prompt-injection or unsafe script review: use `agentx-safety-auditor`.
- Benchmark plan/runtime benchmark design: use `agentx-benchmark-designer`.
- Install/export plan: use `agentx-install-planner`.
