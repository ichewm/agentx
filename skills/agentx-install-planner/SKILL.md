---
name: agentx-install-planner
description: Produce install/export plans for target-ready capability artifacts without executing installation by default.
---

# AgentX Install Planner

Use this skill to prepare install/export plans.

Follow `spec/08-installation-delivery.md`.

Before producing an install-ready plan:

1. Check the target-ready gate from `spec/08-installation-delivery.md`.
2. Require `reviews/unknown-resolution.md`.
3. Refuse installation planning if any required review is missing, blocked, or contains unresolved `Unknown`, `TBD`, or `TODO` placeholders.
4. Output `Blocked` with exact missing evidence instead of a plan when the artifact is not target-ready.
5. Never treat helper verification as semantic proof that the capability will be used correctly by the runtime.

## Install Plan

Write install/export plans under:

```text
.agentx/output/capabilities/<id>/install/<target-id>.plan.md
```

Use this structure:

```markdown
# Install Plan: <capability-id> -> <target-id>

Status: blocked | ready

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

If blocked, include exact missing or blocked target-ready evidence and do not include executable install commands as approved actions.

## Execution Boundary

Default mode is plan-only. Execute installation only when the user explicitly says to proceed. Helper delivery commands require explicit `--yes` confirmation; in other words, helper installation requires explicit --yes confirmation. Prefer official installers when the target profile defines them; otherwise use helper file delivery only for deterministic copy, backup, rollback, and verification. The helper does not execute symlink delivery.
