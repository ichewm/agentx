# Changelog

All notable changes to AgentX are documented in this file.

## 0.1.0 - 2026-05-27

### Added

- Defined AgentX as a runtime-agnostic agent capability workbench.
- Added authoritative English specifications under `spec/` with Chinese `.zh.md` review translations.
- Added AgentX meta-skills for workbench routing, capability architecture, capability translation, review, safety audit, benchmark design, and install planning.
- Added target profiles for Codex, Claude Code, Copilot, Cursor, OpenClaw, and Hermes.
- Added model profiles for GPT, Claude, GLM, and Hermes.
- Added reference standards, conversion playbooks, templates, benchmark matrices, and interaction simulations.
- Added a Go helper for detection, target-ready verification, install/export planning, deterministic copy delivery, backup, rollback, lock records, and benchmark execution.
- Added target-specific runtime benchmark evidence using `runtime-benchmark.<target-id>.md`.
- Added regression coverage for target runtime isolation and install-plan manual requirements.

### Changed

- Removed the legacy Zig registry/control-plane direction from the project scope.
- Made AgentX repository layout spec-first: `spec/`, `skills/`, `targets/`, `models/`, `references/`, `helper/`, and root agent instruction files.
- Clarified that helper delivery does not execute symlink delivery; symlink plans must be manual or delegated to official target installers until explicitly implemented and benchmarked.
- Changed install-ready checks so one target's blocked runtime benchmark does not block another target with passing runtime evidence.
- Changed install plans to include manual requirements recorded in `unknown-resolution.md`.
- Changed skill-package delivery to use `.agentx/output/capabilities/<id>/targets/<target-id>/<skill-id>/` as the package root for Codex, Claude Code, OpenClaw, and Hermes, without requiring `<skill-id>` to equal the capability id.
- Changed install plans to scope manual requirements to the requested target while preserving shared manual requirements.
- Required every capability workspace to include `open-questions.md`, even when there are no unresolved questions.

### Verified

- Added 44 automated benchmark scenarios covering workflow gates, review status handling, Unknown resolution, placeholders, target-native package shape, symlink rejection, install/export confirmation, dangerous destination rejection, rollback restoration, skill quality, and randomized target-ready fuzzing.
- Expanded the benchmark suite to 46 scenarios after adding target-specific runtime benchmark isolation.
- Expanded the benchmark suite to 48 scenarios after adding target-scoped manual requirement checks and package-id independence coverage.
- Verified `go test ./...`, `go vet ./...`, JSON benchmark output, and multiple deterministic fuzz seeds.
