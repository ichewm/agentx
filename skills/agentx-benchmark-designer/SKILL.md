---
name: agentx-benchmark-designer
description: Design static, runtime, and cross-target benchmark plans for generated capabilities.
---

# AgentX Benchmark Designer

Use this skill to create required benchmark artifacts.

Follow `spec/07-review-safety-benchmark.md`.

Benchmark results may justify proposing a baseline adjustment, but they do not automatically authorize applying it. If a benchmark exposes a weakness in an official-baseline artifact, record the evidence so `agentx-capability-translator` and `agentx-capability-reviewer` can prepare a user-confirmed baseline adjustment proposal.

## Required Files

Write:

```text
.agentx/output/capabilities/<id>/reviews/benchmark-plan.md
.agentx/output/capabilities/<id>/reviews/runtime-benchmark.md
```

## Benchmark Plan Sections

`benchmark-plan.md` must include:

- positive trigger cases
- negative trigger cases
- decoy tasks
- required behavior checklist
- forbidden behavior checklist
- target-specific routing checks
- cross-target drift checks when more than one target exists

## Runtime Benchmark Status

`runtime-benchmark.md` must use one status line:

```text
Status: passed
Status: manual-transcript
Status: blocked
```

Final delivery requires `passed` or `manual-transcript`. If automation is unavailable and no manual transcript exists, record `Status: blocked` and the blocking gap.

## Runtime CLI Safety

Do not execute external AI runtime CLIs without explicit user confirmation for that benchmark run. This includes Codex, Claude Code, Copilot, Cursor, OpenClaw, Hermes, and future runtime CLIs.

Do not install temporary skills into real user runtime directories, create or alter authenticated runtime homes, or trigger runtime update/download flows without explicit user confirmation.

When confirmation is missing, write `Status: blocked` in `runtime-benchmark.md` and record the exact confirmation gap instead of running the CLI.
