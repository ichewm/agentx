---
name: agentx-capability-translator
description: Translate an existing skill, plugin, rule, prompt pack, or tool capability into target-native artifacts for other runtimes.
---

# AgentX Capability Translator

Use this skill to convert existing capabilities.

Follow `spec/06-translation-workflow.md` and the baseline creator rules in `spec/03-meta-skills.md`.

## Procedure

1. Accept source input as a path, URL, repository, archive, directory, or pasted content.
2. Inspect readable paths and URLs directly; do not ask the user to paste locally readable files.
3. Create or reuse `.agentx/output/capabilities/<id>/`.
4. Ensure `intake.md`, `open-questions.md`, and `brief.md` all exist; write `open-questions.md` even when every question is resolved.
5. Store source provenance and summaries under `intake.md` and `sources/`.
6. Identify source runtime, source artifact types, scripts, references, commands, tools, rules, assets, and install assumptions.
7. Create or update `brief.md` with a capability map and conversion status for each mapped item.
8. Read every requested `targets/<target-id>/` profile before generating artifacts.
9. Generate target-native artifacts under `.agentx/output/capabilities/<id>/targets/<target-id>/`.
10. For skill-package targets, put the runtime package root under the target id, for example `.agentx/output/capabilities/<id>/targets/codex/<skill-id>/SKILL.md`.
11. Produce a conversion loss report with Preserved, Adapted, Degraded, Dropped, Manual Setup Required, and Risks sections.

When a target has an official creator baseline:

1. Generate the baseline-compliant target artifact first.
2. Do not alter required target-native structure, metadata, package layout, or installation semantics.
3. If review or benchmark evidence suggests an adjustment, write a proposal instead of silently applying it.
4. Require user confirmation before applying any adjustment that changes target artifacts, runtime behavior, safety posture, installation semantics, or cross-target meaning.
5. Record approved deviations in `reviews/baseline-deviations.md`.

## Translation Rules

- Preserve capability intent before source file shape.
- Never claim lossless conversion unless target primitives are equivalent.
- Never invent target conventions when official target conventions exist.
- Unsupported target behavior must be marked `unsupported`, `manual`, or `scoped-out`.
- Any unresolved platform fact must flow into `reviews/unknown-resolution.md` before target-ready status.
