# 00. Positioning

## Product Definition

AgentX is a runtime-agnostic agent capability workbench.

It helps an AI agent create, translate, review, package, and deliver capabilities such as skills, rules, prompts, tools, and guards across agent runtimes.

AgentX is not a model. AgentX is not a replacement for Codex, Claude Code, GitHub Copilot, Cursor, OpenClaw, or Hermes Agent. AgentX is a meta-capability layer loaded into one of those runtimes.

## What AgentX Is

AgentX is:

- A set of meta-skills for building other agent capabilities.
- A reference library describing target runtimes and their native capability formats.
- A workflow for turning user intent, documents, or existing skills/plugins into target-ready artifacts.
- A review system for semantic equivalence, portability, safety, and benchmark readiness.
- A thin deterministic helper for install planning, file delivery, backup, rollback, and verification.

## What AgentX Is Not

AgentX is not:

- A runtime-time skill selector for Codex, Claude Code, Copilot, or other runtimes with native skill discovery.
- A new universal skill standard replacing Agent Skills / `SKILL.md`.
- A tool that installs skills into models.
- A control plane centered on a registry file.
- A system that compiles every capability into one generated instruction file.
- A platform-specific plugin manager.

## Core Concept

The logic is the capability document. The runtime is the AI platform that reads and executes it.

AgentX therefore starts as capabilities loaded into an existing AI runtime. The current AI uses AgentX meta-skills to produce other capabilities.

## System Boundary

Creative and semantic work belongs to the loaded AI:

- Understand user intent.
- Read source materials.
- Extract a capability brief.
- Translate existing skills/plugins/rules.
- Produce target artifacts.
- Review semantic equivalence.
- Audit safety and portability.
- Design benchmark cases.

Deterministic work belongs to the helper:

- Detect target runtime paths.
- Copy or link files.
- Compute hashes.
- Write lock records.
- Create backups.
- Roll back installations.
- Verify installed file layout.

## Long-Term Shape

AgentX must be extensible by adding target profiles and playbooks, not by changing core code.

New runtime support must be added through:

- `targets/<runtime>/profile.md`
- `targets/<runtime>/creator-baseline.md`
- `targets/<runtime>/package-layout.md`
- `targets/<runtime>/install.md`
- `targets/<runtime>/review-checklist.md`

## Rewrite Scope

The original control-plane implementation has been removed.

The repository must keep only the materials needed for the AI-native workbench:

- Specifications.
- AgentX meta-skills.
- Target profiles.
- Reference documents.
- Thin helper source.
- Reference examples and benchmark fixtures under `references/`.

Do not recreate legacy files for compatibility unless the user explicitly asks for a migration bridge.
