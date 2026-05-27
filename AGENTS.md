# AgentX

AgentX is a runtime-agnostic agent capability workbench.

The current project direction is defined by the specification in `spec/`. The old Zig registry/control-plane design is not part of AgentX and has been removed.

## Authoritative Documents

AI agents must use the English documents as the source of truth:

- `spec/*.md`
- `skills/*/SKILL.md`
- `targets/*/*.md`
- `references/**/*.md`
- `AGENTS.md`
- `CLAUDE.md`

Chinese files ending in `.zh.md` are human review translations only. Do not use `.zh.md` files as implementation authority. If English and Chinese differ, follow the English file and report the mismatch.

## Product Direction

AgentX is not a local registry control plane.

AgentX is:

- A set of meta-skills loaded into an existing AI runtime.
- A workbench for creating, translating, reviewing, packaging, and delivering agent capabilities.
- A reference system for target runtimes and their native capability formats.
- A thin deterministic helper for install planning, file delivery, backup, rollback, and verification.

AgentX is not:

- A runtime skill selector for Codex, Claude Code, Copilot, OpenClaw, Hermes, or other runtimes with native capability discovery.
- A new universal skill standard replacing Agent Skills / `SKILL.md`.
- A tool that installs skills into models.
- A generated-instructions sync system.
- A GLM target adapter unless a concrete GLM runtime with native capability installation is identified.

## Core Workflow

The primary workflow is AI-native:

1. The user loads AgentX meta-skills into the current AI runtime.
2. The user provides intent, documents, or an existing skill/plugin/rule.
3. The AI uses AgentX meta-skills to create a capability brief.
4. The AI reads target profiles and official creator baselines.
5. The AI generates target-native artifacts.
6. The AI runs semantic review, portability review, safety audit, and benchmark design.
7. The AI produces an install/export plan.
8. The helper performs deterministic delivery only after user confirmation.

## Repository Direction

New work must create or maintain only:

```text
skills/
targets/
models/
references/
helper/
spec/
```

The legacy `src/` Zig implementation and old specs have been removed. Do not recreate the old registry, sync, approve, guard, automation, or generated-instructions architecture.

## Implementation Guidance

- Prefer spec-first changes.
- Keep all directive, skill, target, and reference files in English.
- Keep matching Chinese `.zh.md` translations for every spec file. Chinese translations are not required for skills, targets, model profiles, or references.
- Treat runtime support as data: add target profiles and playbooks rather than hard-coding behavior.
- Prefer official target installers and official creator guidance when available.
- Use Go for the thin helper.
- Use shell only for bootstrap scripts.

## Benchmarks

AgentX capabilities must include benchmark plans and runtime benchmark records. Benchmarks must test trigger behavior, semantic adherence, safety boundaries, and cross-target drift. Runtime benchmark automation is mandatory for generated capabilities; if automation is unavailable, a manual runtime benchmark transcript is required before final delivery.
