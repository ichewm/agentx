---
name: agentx-capability-architect
description: Create a capability brief from user intent, documents, examples, constraints, and multi-turn intake.
---

# AgentX Capability Architect

Use this skill to build a capability from zero.

Follow `spec/01-ai-native-workflow.md`.
Follow `spec/05-capability-brief.md`.

## Procedure

1. Create or choose a stable capability id.
2. Create the workbench directory:

```text
.agentx/output/capabilities/<id>/
  intake.md
  sources/
  open-questions.md
  brief.md
```

3. Record every user-provided requirement, document, example, constraint, path, or URL in `intake.md`.
4. Store large pasted material or fetched files under `sources/`.
5. Maintain `open-questions.md`, but do not force answers during intake unless the next source cannot be understood. Create this file even when there are no open questions.
6. After each intake turn, summarize changes and ask whether the user wants to add more input.
7. After intake closes, summarize intent, sources, and decision points before generating final artifacts.
8. Write `brief.md` using the required sections from `spec/05-capability-brief.md`.

## Brief Rules

- Preserve intent, not file shape.
- Separate stable capability logic from target-specific expression.
- Label missing information explicitly.
- Do not convert incomplete assumptions into final target behavior.
- If the user asks for an early draft, mark it as draft and keep unresolved decisions in `open-questions.md`.

## Output

Produce:

- Updated `intake.md`
- Updated `sources/` files when needed
- Updated `open-questions.md`
- `brief.md`
- A short status message naming unresolved decisions and next specialized meta-skill.
