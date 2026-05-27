# AgentX

AgentX is a runtime-agnostic agent capability workbench.

Use the English specifications in `spec/` as the source of truth. The old Zig registry/control-plane design is not part of AgentX and has been removed.

## Read First

1. `spec/README.md`
2. `spec/00-positioning.md`
3. `spec/01-ai-native-workflow.md`
4. `spec/02-repository-layout.md`
5. `spec/03-meta-skills.md`
6. `spec/04-target-profiles.md`
7. `spec/05-capability-brief.md`
8. `spec/06-translation-workflow.md`
9. `spec/07-review-safety-benchmark.md`
10. `spec/08-installation-delivery.md`
11. `spec/09-implementation-plan.md`

Chinese `.zh.md` files are human review translations only. Do not treat them as implementation authority.

## Project Direction

AgentX is delivered primarily as meta-skills. These meta-skills help the current AI runtime create, translate, review, package, and deliver other agent capabilities.

The deterministic helper is intentionally thin. It must handle detection, install planning, file delivery, backup, rollback, hash records, and verification. It must not replace AI semantic review.

## Do Not Build Against Legacy Assumptions

Do not base new work on:

- `.agents/registry.yaml` as the single source of truth.
- `agentx sync` as the central workflow.
- generated target instruction files as the main artifact.
- `approve`, `guard`, hook, automation, or policy control-plane flows.
- GLM as an install target without a concrete GLM runtime.

## Language Rule

All directive files, specs, skills, target profiles, and implementation references must be written in English.

English spec changes must be synchronized to matching `.zh.md` files. Chinese translations are required for specs, but not for skills, targets, model profiles, or references.
