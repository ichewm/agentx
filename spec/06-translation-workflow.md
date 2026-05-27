# 06. Translation Workflow

## Purpose

Translation is the main differentiated AgentX workflow.

AgentX must turn a source capability from one runtime into target-native artifacts for other runtimes while making conversion loss explicit.

## Inputs

Possible inputs:

- User request.
- Existing `SKILL.md`.
- Claude Code plugin.
- Codex skill.
- GitHub Copilot skill.
- Cursor rule.
- OpenClaw skill pack.
- Hermes skill.
- MCP server or prompt pack.
- Documentation and examples.

The default source input priority is path, repository URL, then directory. Pasted content is fully supported and must not be treated as a degraded input mode. The workflow must not require repeated copy/paste across turns.

If source materials are large, the AI must create an intake summary and store it in:

```text
.agentx/output/capabilities/<id>/sources/source-summary.md
```

The AI must also record source provenance in the capability brief.

## Source Intake Modes

### Path or Folder

When the user provides a local path or folder, the AI must inspect the file tree, identify candidate capability files, and summarize what it found.

The AI must not ask the user to paste files that are locally readable.

### GitHub or Web URL

When the user provides a URL, the AI must fetch or browse it when tools allow, record the source URL, and summarize relevant files or docs.

If the URL cannot be accessed, ask the user for a downloadable archive, local clone path, or the smallest needed files.

### Pasted Content

When the user pastes content, the AI must store it as an intake entry.

If the content is partial, the AI must ask whether more chunks are coming before generating final artifacts.

### Multiple Turns

If the user provides source material over multiple messages, each message must become a separate intake entry with a short label and timestamp if available.

The AI must maintain a running source summary and avoid requiring repeated copy/paste.

## Steps

1. Identify source runtime.
2. Inventory source files.
3. Extract capability map.
4. Create or update capability brief.
5. Read target profiles.
6. Produce target plans.
7. Generate target artifacts.
8. Run reviews.
9. Revise artifacts.
10. Produce install/export plans.

## Target Generation Rule

Always prefer the target runtime's native capability format.

Examples:

- Codex: Agent Skills / `SKILL.md` package if supported.
- Claude Code: skill or plugin layout according to official guidance.
- Copilot: supported skill directories or official skill tooling.
- Cursor: `.cursor/rules/*.mdc`, not fake skills.
- MCP: tools, prompts, or resources, not fake skills.

## Official Creator Baselines

When a target has an official creator skill or guidance, AgentX must use that as the generation baseline.

AgentX may add stricter review and cross-target checks, but must not invent target conventions when official ones exist.

Any deviation from an official creator baseline must follow the baseline adjustment protocol in `spec/03-meta-skills.md`. Translation must not silently apply benchmark-driven changes as if they were official target behavior.

## Conversion Loss Report

Every translation must include:

```markdown
## Preserved

## Adapted

## Degraded

## Dropped

## Manual Setup Required

## Risks
```

No translation may claim lossless conversion unless the target runtime supports equivalent primitives.
