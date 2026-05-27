# 09. Implementation Plan

## Phase 0: Spec Reset

Status: completed in this repository.

Delivered:

- English specs.
- Chinese review translations.
- Root agent instructions pointing to English specs.
- Removal of the old control-plane code and old specs.

## Phase 1: Meta-Skill Pack

Deliver initial AgentX meta-skills:

- `agentx-workbench`
- `agentx-capability-architect`
- `agentx-capability-translator`
- `agentx-capability-reviewer`
- `agentx-safety-auditor`
- `agentx-benchmark-designer`
- `agentx-install-planner`

Each meta-skill must be usable by Codex and Claude Code as a standard skill package.

## Phase 2: Target Profiles

Deliver initial target profiles:

- `codex`
- `claude-code`
- `copilot`
- `cursor`
- `openclaw`
- `hermes`

Each target profile must cite official behavior or mark unknowns as unknown.

Deliver initial model profiles for target-relevant models. Model profiles affect expression and review sensitivity only.

## Phase 3: Workbench Output Convention

Deliver:

- `.agentx/output/capabilities/<id>/` layout.
- Source intake convention for paths, URLs, repositories, and pasted content.
- Multi-turn intake convention with `intake.md` and `open-questions.md`.
- Post-input continuation prompt and intake close signals.
- Capability brief template.
- Review templates.
- Install plan template.
- Lock file schema.

## Phase 4: Thin Helper

Helper language: Go.

Reason:

- Cross-platform single binary.
- Strong standard library for files, JSON, hashing, HTTP, archives, and subprocesses.
- Easier and faster for installer work than Zig.
- Safer and more testable than shell.
- Fewer runtime assumptions than TypeScript.

Shell must be used only for bootstrap scripts.

Zig is not a peer implementation choice for AgentX. It can be reconsidered later only for a concrete technical reason.

## Phase 5: Helper Commands

Initial helper commands:

```text
agentx detect <target>
agentx plan install <capability> --target <target>
agentx install <capability> --target <target>
agentx export <capability> --target <target>
agentx verify <capability> --target <target>
agentx rollback <capability> --target <target>
agentx list
```

No command may replace AI semantic review.

## Phase 6: Benchmarks

Deliver benchmark templates and at least one benchmark plan per sample capability.

Runtime benchmark automation is mandatory for generated capabilities. If a target runtime cannot be automated yet, AgentX must record the gap and require a manual runtime benchmark transcript before final delivery.

## Current Product Non-Goals

- No GLM target adapter unless a concrete GLM runtime with native capability installation is identified.
- No registry as the single source of truth.
- No all-target `sync`.
- No hook/automation control plane.
- No marketplace publishing.
- No GUI.
