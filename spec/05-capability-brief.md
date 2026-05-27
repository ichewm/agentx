# 05. Capability Brief

## Purpose

The capability brief is the stable intermediate document used before generating target-specific artifacts.

It is not a new public skill standard. It is an AgentX work product that helps the current AI keep intent consistent across targets.

## Location

```text
.agentx/output/capabilities/<id>/brief.md
```

The brief is built over time from:

```text
.agentx/output/capabilities/<id>/intake.md
.agentx/output/capabilities/<id>/sources/
.agentx/output/capabilities/<id>/open-questions.md
```

The AI must update the brief incrementally as the user provides more material.

## Required Sections

```markdown
# Capability Brief: <id>

## Goal

## Source Materials

## Intended Runtimes

## When To Use

## When Not To Use

## Inputs To Inspect

## Procedure

## Hard Constraints

## Safety Notes

## Expected Output

## Examples

## Non-Goals

## Conversion Notes
```

## Capability Map

For translation work, the brief must include a capability map:

```text
skills:
commands:
tools:
connectors:
subagents:
rules:
scripts:
references:
assets:
install assumptions:
```

## Conversion Status

Each mapped item must be labeled:

- `direct`: can be preserved in the target runtime.
- `adapted`: can be rewritten into a target-native form.
- `degraded`: can be represented only with reduced behavior.
- `manual`: requires human setup.
- `unsupported`: cannot be represented honestly.

## Design Rule

The brief must preserve intent, not source file shape.

Target artifacts must be generated from the brief and target profiles. They must not be blind file-to-file conversions.

## Partial Briefs

During multi-turn intake, the brief may be incomplete.

Incomplete briefs must include:

```markdown
## Missing Information

## Assumptions

## Questions For User
```

The AI may proceed with a draft artifact only if the user explicitly asks for a draft despite missing information.
