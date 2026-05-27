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

### Changed

- Removed the legacy Zig registry/control-plane direction from the project scope.
- Made AgentX repository layout spec-first: `spec/`, `skills/`, `targets/`, `models/`, `references/`, `helper/`, and root agent instruction files.
- Clarified that helper delivery does not execute symlink delivery; symlink plans must be manual or delegated to official target installers until explicitly implemented and benchmarked.

### Verified

- Added 44 automated benchmark scenarios covering workflow gates, review status handling, Unknown resolution, placeholders, target-native package shape, symlink rejection, install/export confirmation, dangerous destination rejection, rollback restoration, skill quality, and randomized target-ready fuzzing.
- Verified `go test ./...`, `go vet ./...`, JSON benchmark output, and multiple deterministic fuzz seeds.
