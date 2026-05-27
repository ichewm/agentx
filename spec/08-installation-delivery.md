# 08. Installation and Delivery

## Definition

Installation means delivering generated artifacts to a runtime's native consumption path.

AgentX does not install anything into a model.

## Delivery Principle

Use official target installers when available.

The default delivery mode is planning only. AgentX must produce an install/export plan and wait. AgentX must execute an install only after the user explicitly asks it to proceed.

The AgentX helper exists for:

- Install planning.
- Copy delivery when official installers are absent.
- Symlink delivery only as a manual or official-installer plan unless helper symlink support is implemented later.
- Backup.
- Rollback.
- Verification.
- Lock records.

## Artifact Location

Generated artifacts live first in:

```text
.agentx/output/capabilities/<id>/targets/<target-id>/
```

AgentX installs or exports from that location.

## Target-Ready Gate

AgentX must not produce an install-ready plan until the target artifact passes the target-ready gate.

Required target-ready evidence:

- `reviews/semantic-review.md` exists and does not mark the artifact blocked.
- `reviews/portability-review.md` exists and does not mark the artifact blocked.
- `reviews/safety-review.md` exists and does not mark the artifact blocked.
- `reviews/benchmark-plan.md` exists.
- `reviews/runtime-benchmark.md` exists and records either a successful automated runtime benchmark or a manual runtime benchmark transcript.
- `reviews/unknown-resolution.md` exists and has no unresolved Unknown, TBD, TODO, or blank resolution.
- `reviews/baseline-deviations.md` exists when the artifact deviates from an official creator baseline.
- Generated target files under `.agentx/output/capabilities/<id>/targets/<target-id>/` contain no unresolved placeholder terms such as `Unknown`, `TBD`, or `TODO`.

The install planner must output `Blocked` instead of an install-ready plan when any target-ready evidence is missing or blocked.

The deterministic helper may enforce this gate by checking required files, blocked markers, placeholder terms, and lock metadata. It must not treat those checks as proof of semantic correctness.

## Install Plan

Every delivery action must produce a plan before changing files.

```markdown
# Install Plan: <capability-id> -> <target-id>

## Source Artifact

## Target Runtime

## Method

## Destination

## Changes

## Backup

## Verification

## Rollback

## User Confirmation
```

## Methods

### official-cli

Use the runtime's official CLI.

AgentX must show the exact command and ask before running it.

### path-copy

Copy generated files to the runtime path.

Required:

- Backup existing destination.
- Copy files.
- Verify layout.
- Write lock record.

### path-symlink

Symlink generated artifact to runtime path.

Use only when the target runtime safely supports symlinks.

The current helper does not execute symlink delivery. A symlink plan must be marked manual or delegated to an official target installer until helper symlink support is explicitly implemented and benchmarked.

### export-only

Write generated files into project configuration.

Example:

- Cursor rule export to `.cursor/rules/*.mdc`.

### manual

Output instructions but make no filesystem changes.

## Verification

Verification must prove file delivery only. It must not claim that the model will definitely trigger the capability unless a runtime benchmark was run.
