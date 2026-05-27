# 04. Target Profiles

## Purpose

Target profiles externalize platform knowledge. AgentX must add new runtimes by adding profile files and playbooks, not by changing the core workflow.

## Target Directory

Each target directory must have all five files:

```text
targets/<target-id>/
  profile.md
  creator-baseline.md
  package-layout.md
  install.md
  review-checklist.md
```

## profile.md

Describes runtime capabilities.

Required topics:

- Runtime name and scope.
- Whether native skills are supported.
- Whether rules are supported.
- Whether prompts, commands, tools, connectors, hooks, or subagents are supported.
- How the runtime routes capabilities.
- Whether the model chooses capabilities at runtime.
- Supported scopes: project, user, organization, plugin.
- Known refresh or restart requirements.
- Known unsupported features.

## creator-baseline.md

Summarizes official guidance for creating capabilities for this target.

AgentX must use this as the target-specific compiler reference.

## package-layout.md

Defines generated artifact layout for this target.

Examples:

- `SKILL.md` package.
- Claude Code plugin layout.
- Cursor `.mdc` rule.
- MCP server tool/prompt/resource layout.

## install.md

Defines delivery method.

Possible methods:

- `official-cli`
- `official-ui`
- `path-copy`
- `path-symlink`
- `export-only`
- `manual`

Official installers take precedence over helper copy operations.

## Unknown Values

Unknowns are allowed only in target and model profile research state, and they must be explicit.

Use `Unknown` when a fact has not been verified. Unknown values prevent hallucinated platform claims. They do not prevent the profile from being loaded, but they do affect generation:

- AgentX may create internal drafts with Unknown fields.
- AgentX must not claim final target compatibility for a feature marked Unknown.
- AgentX must not execute installation for a target when the install method or destination is Unknown.
- AgentX must not put unresolved Unknown-dependent behavior into final generated artifacts.
- AgentX must either research the Unknown, ask the user, mark the feature unsupported, or mark the feature manual before final delivery.
- If an Unknown affects routing, package layout, installation, safety, or benchmark validity, final delivery for that target is blocked until the Unknown is resolved or explicitly scoped out.

## Unknown Resolution Gate

Every target artifact must pass an Unknown resolution gate before it can be marked target-ready.

The AI must produce:

```text
.agentx/output/capabilities/<id>/reviews/unknown-resolution.md
```

The file must list every Unknown from the target profile, model profile, source capability, conversion notes, and generated draft that could affect the target artifact.

Each entry must have one of these resolutions:

- `verified`: resolved through cited official documentation, inspected local runtime behavior, or user-provided authoritative input.
- `unsupported`: the feature is not supported by the target and is excluded from the final artifact.
- `manual`: the feature requires manual setup and is documented in the install/export plan.
- `scoped-out`: the feature is outside the user's requested scope for this artifact.

No entry may remain `Unknown`, `TBD`, `TODO`, or blank.

The final target artifact directory must not contain unresolved placeholder terms such as `Unknown`, `TBD`, or `TODO` in generated target files, install plans, lock records, or final review summaries. Those terms are allowed only in profile research files, historical draft notes, or `unknown-resolution.md` entries that name the original unresolved source fact while assigning a resolved status.

The deterministic helper must fail finalization if it detects unresolved placeholder terms in generated target files, install plans, lock records, or final review summaries. It must not fail merely because `unknown-resolution.md` mentions an original Unknown source item with a non-Unknown resolution. The helper must not decide whether a semantic Unknown was correctly resolved; that remains an AI review responsibility.

## review-checklist.md

Defines target-specific review requirements.

Examples:

- Routing description quality.
- Script permission assumptions.
- Progressive disclosure structure.
- Target-specific frontmatter.
- Unsupported feature handling.

## Model Profiles

Models are not installation targets.

Model profiles are part of the complete AgentX product from the first implementation.

GLM, GPT, Claude, Hermes, or other models may have model profiles, but those profiles affect expression and review sensitivity only. Installation always targets a runtime.

Model profiles live under:

```text
models/<model-id>/
  profile.md
  expression-guidance.md
  review-sensitivity.md
```

Model profiles must not define installation paths.
