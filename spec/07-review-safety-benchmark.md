# 07. Review, Safety, and Benchmark

## Review Layers

AgentX uses AI-based semantic review plus deterministic helper checks.

AI review answers:

- Does this artifact express the intended capability?
- Can the target runtime route to it correctly?
- Is the target artifact too broad or too narrow?
- Did translation change meaning?
- Are there unsafe or hidden instructions?

Helper checks answer:

- Do required files exist?
- Is frontmatter parseable?
- Do referenced files exist?
- Are paths inside allowed directories?
- Do hashes match the lock file?

## Safety Audit

Every imported or translated capability must be audited for:

- Prompt injection.
- Hidden instruction override.
- Secret access.
- Credential access.
- Dangerous shell commands.
- Network exfiltration.
- Obfuscated scripts.
- Excessive tool permissions.
- Unpinned external dependencies.

High-risk findings require human confirmation before installation.

## Benchmark Purpose

Benchmarks are required because skill quality depends on the model and runtime.

AgentX benchmarks are not only model scores. They are regression checks for whether a capability is triggered and followed by a target runtime.

## Benchmark Levels

Level 1: Static benchmark

- Positive examples.
- Negative examples.
- Decoy tasks.
- Required behavior checklist.
- Forbidden behavior checklist.

Level 2: Runtime benchmark

- Run target runtime manually or through supported automation.
- Observe whether the runtime selects the capability.
- Record whether behavior follows the artifact.
- External AI runtime CLIs, including Codex, Claude Code, Copilot, Cursor, OpenClaw, and Hermes, must not be executed automatically without explicit user confirmation for that benchmark run.
- Runtime automation must not install temporary skills into real user runtime directories, create or alter authenticated runtime homes, or trigger runtime update/download flows without explicit user confirmation.

Level 3: Cross-target benchmark

- Compare Codex, Claude Code, Copilot, Cursor, OpenClaw, Hermes, or other targets.
- Record target-specific drift.

## Benchmark File

```text
.agentx/output/capabilities/<id>/reviews/benchmark-plan.md
```

The benchmark plan must be readable by humans and usable by automation.

## Required Benchmark Artifacts

Every generated capability must include:

```text
.agentx/output/capabilities/<id>/reviews/benchmark-plan.md
.agentx/output/capabilities/<id>/reviews/runtime-benchmark.<target-id>.md
```

Runtime automation is mandatory for generated capabilities. If a runtime cannot be automated yet, AgentX must record the blocking gap in that target's `runtime-benchmark.<target-id>.md` or record a manual runtime benchmark transcript in that same target-specific file.

For final delivery, `runtime-benchmark.<target-id>.md` must record either a successful automated runtime benchmark or a manual runtime benchmark transcript. If it records only a blocking gap, only that target artifact is not target-ready. A blocked runtime benchmark for one target must not block another target with passing runtime evidence.

Benchmark automation must not claim success unless it actually ran against the target runtime or a documented official test harness.

If the only available automation path requires launching an external AI runtime CLI or creating an authenticated runtime environment, the AI must ask first. Without explicit approval, record `Status: blocked` in that target's runtime benchmark file with the exact user-confirmation gap.
