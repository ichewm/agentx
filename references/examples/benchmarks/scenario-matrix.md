# Benchmark Scenario Matrix

## Static Gate

- `clean-skill-ready`: all required review files passed and target files are clean.
- `missing-brief-blocks`: required workflow artifacts are absent.
- `missing-review-blocks`: one required review file is absent.
- `missing-status-blocks`: a required final review exists but has no final status line.
- `needs-user-decision-blocks`: a required final review still needs user decision.
- `multiple-status-blocks`: a required final review has ambiguous multiple status lines.
- `blocked-review-blocks`: one required review file has `Status: blocked`.
- `review-placeholder-blocks`: a final review still contains unresolved placeholder text.
- `target-placeholder-blocks`: generated target file contains `Unknown`, `TBD`, or `TODO`.
- `target-symlink-blocks`: generated target artifact contains a symlink.
- `randomized-target-ready-fuzz`: fixed-seed random combinations of missing workflow files, missing reviews, invalid statuses, unresolved Unknowns, placeholders, runtime gaps, and missing baseline deviation records.

## Unknown Gate

- `unknown-resolution-passes`: original Unknown is recorded with `Resolution: verified`.
- `unknown-resolution-blocks`: resolution remains Unknown, TBD, TODO, or blank.

## Runtime Benchmark

- `runtime-automated-passes`: `runtime-benchmark.md` has `Status: passed`.
- `runtime-manual-passes`: `runtime-benchmark.md` has `Status: manual-transcript`.
- `runtime-gap-blocks`: `runtime-benchmark.md` has `Status: blocked`.

## Baseline

- `baseline-deviation-recorded`: review marks `Baseline deviation: yes` and `baseline-deviations.md` exists.
- `baseline-deviation-missing-blocks`: review marks `Baseline deviation: yes` and `baseline-deviations.md` is missing.

## Installation Boundary

- `install-without-confirmation-blocks`: helper refuses delivery without `--yes`.
- `plan-blocked-artifact`: helper writes blocked plan instead of install-ready plan.
- `plan-ready-artifact`: helper writes a ready plan for a target-ready artifact without executing installation.
- `dangerous-destination-blocks`: helper refuses repository root, home, filesystem root, `.agentx`, and source-overlapping destinations.
- `repo-subdir-destination-blocks`: helper refuses destinations inside the AgentX repository.
- `invalid-capability-id-blocks`: helper refuses capability ids that are paths or parent traversal.
- `invalid-target-id-blocks`: helper refuses target ids that are paths, parent traversal, or missing target profiles.
- `export-copies-ready-artifact`: helper copies a target-ready artifact and writes a lock record.
- `rollback-restores-backup`: helper restores the latest lock-recorded backup.

## Target Shape

- `codex-skill-package`: generated artifact is `SKILL.md` package.
- `codex-missing-skill-blocks`: a Codex artifact without `SKILL.md` is blocked even if the target directory exists.
- `cursor-rule-package`: generated artifact is `.cursor/rules/*.mdc` and no fake `SKILL.md`.

## Workflow Coverage

- `multi-turn-intake-coverage`: intake rules require repeated input capture and an explicit close signal.
- `final-generation-gate-coverage`: final artifacts require closed intake and user confirmation unless explicitly requested as an early draft.
- `source-intake-coverage`: source conversion accepts paths, URLs, repositories, directories, archives, and pasted content.
- `architect-decision-point-coverage`: the architect summarizes sources and decision points before final generation.

## Skill Quality

- `skill-metadata-quality`: each AgentX skill has high-signal frontmatter with a clear action-oriented description.
- `skill-operational-sections`: each AgentX skill has an actionable procedure, routing procedure, required-file section, or equivalent operational instruction.
- `skill-progressive-disclosure-quality`: each AgentX skill points to the correct authoritative specs and specialized downstream outputs.
- `skill-directory-minimal`: AgentX meta-skills remain instruction-only unless a future spec explicitly adds local assets.
- `workbench-does-not-duplicate`: `agentx-workbench` remains a lightweight router and does not absorb specialized skill bodies.
- `translator-conversion-loss-coverage`: translation always accounts for preserved, adapted, degraded, dropped, manual, and risk items.
- `benchmark-method-source-coverage`: benchmark design remains grounded in external benchmark and skill-quality references.
