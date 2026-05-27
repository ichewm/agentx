# 02. Repository Layout

## Top-Level Layout

The original source tree has been removed. The repository must stay small and spec/meta-skill centered.

```text
agentx/
  AGENTS.md
  CLAUDE.md
  spec/
    *.md
    *.zh.md
  skills/
    agentx-workbench/
    agentx-capability-architect/
    agentx-capability-translator/
    agentx-capability-reviewer/
    agentx-safety-auditor/
    agentx-benchmark-designer/
    agentx-install-planner/
  targets/
    codex/
    claude-code/
    copilot/
    cursor/
    openclaw/
    hermes/
  models/
  references/
    standards/
    examples/
    conversion-playbooks/
  helper/
    README.md
```

The old `src/`, legacy `spec/*.md`, and Zig build files have been removed. Do not recreate them. Do not introduce business application source trees into this repository.

## Generated Output Workspace

AgentX-generated work must live under `.agentx/output/` in the AgentX repository.

Reason: AgentX is loaded from this repository before the user starts a capability creation or conversion conversation. The AgentX repository is therefore the workbench state holder. External repositories, folders, URLs, or pasted documents are sources; target runtimes or target projects are install/export destinations.

```text
.agentx/
  output/
    capabilities/
      <capability-id>/
        intake.md
        open-questions.md
        brief.md
        sources/
        targets/
          codex/
          claude-code/
          cursor/
        reviews/
        install/
        lock.json
  backups/
  logs/
```

`.agentx/` must be ignored by default. Generated work may be committed only when the user explicitly asks to commit a specific generated artifact. The install plan must remind the user whether `.agentx/` is ignored.

If the user intentionally wants a generated capability workspace inside another project, that must be requested explicitly and recorded in the install/export plan.

## Authoritative Files

English files are authoritative for AI execution:

- `AGENTS.md`
- `CLAUDE.md`
- `spec/*.md`
- `skills/*/SKILL.md`
- `targets/*/*.md`
- `models/*/*.md`
- `references/**/*.md`

Chinese files ending in `.zh.md` are human review translations only.

English spec changes must be synchronized to matching `.zh.md` files. Skills, targets, model profiles, and references do not require Chinese translations.

## Capability Artifact Layout

Every generated capability must be grouped by capability id.

```text
.agentx/output/capabilities/<id>/
  intake.md
  open-questions.md
  brief.md
  sources/
  targets/
    <target-id>/
  reviews/
    semantic-review.md
    portability-review.md
    safety-review.md
    benchmark-plan.md
    runtime-benchmark.md
    unknown-resolution.md
    baseline-deviations.md
  install/
    <target-id>.plan.md
  lock.json
```

Target artifacts must be ready to copy, install, or export without requiring the AI to reconstruct file layout.

## Legacy Code

The old Zig control-plane implementation has been removed.

No new workflow may depend on `.agents/registry.yaml`, generated instruction files, `approve`, `sync`, `guard`, hook control planes, or GLM as a target runtime.
