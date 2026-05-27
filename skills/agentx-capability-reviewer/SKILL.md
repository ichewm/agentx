---
name: agentx-capability-reviewer
description: Review capability artifacts for semantic equivalence, routing quality, portability, and target-native correctness.
---

# AgentX Capability Reviewer

Use this skill to review generated target artifacts.

Follow `spec/07-review-safety-benchmark.md` and the baseline creator rules in `spec/03-meta-skills.md`.

When reviewing a target artifact, verify:

1. The artifact follows the official creator baseline when one exists.
2. Required target-native structure, metadata, package layout, and installation semantics are preserved.
3. Any proposed deviation has evidence, risk analysis, target-native legality, and user confirmation when required.
4. No helper command is treated as proof of semantic correctness.
5. `reviews/unknown-resolution.md` resolves every Unknown that affects routing, package layout, installation, safety, benchmark validity, runtime behavior, or cross-target meaning.
6. Final target files do not contain unresolved placeholder terms such as `Unknown`, `TBD`, or `TODO`.

## Required Review Files

Write or update these files under `.agentx/output/capabilities/<id>/reviews/`:

- `semantic-review.md`
- `portability-review.md`
- `unknown-resolution.md`
- `baseline-deviations.md` when the artifact deviates from an official baseline

## Review Status

Each review file must contain one status line:

```text
Status: passed
Status: blocked
Status: needs-user-decision
```

Use `blocked` when the artifact depends on unresolved platform facts, unsafe behavior, missing target-ready evidence, or unsupported target features that are not marked manual or scoped-out.

When the artifact deviates from an official creator baseline, include:

```text
Baseline deviation: yes
```

and require `reviews/baseline-deviations.md`. Otherwise include `Baseline deviation: no`.

## Unknown Resolution

For every Unknown that affects the final artifact, record:

```markdown
## <short-name>

- Source: <target profile, model profile, source artifact, generated draft>
- Impact: <routing | package-layout | installation | safety | benchmark | behavior | cross-target meaning>
- Resolution: <verified | unsupported | manual | scoped-out>
- Evidence: <official URL, inspected runtime behavior, or user-provided authority>
- Final artifact impact: <what changed>
```

No resolution may remain Unknown, TBD, TODO, or blank.
