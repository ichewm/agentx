# AgentX Helper

The helper is a thin deterministic program.

Language: Go.

Responsibilities:

- Detect target runtime paths.
- Produce install/export plans.
- Copy files only after explicit user confirmation.
- Create backups.
- Roll back installations.
- Verify file layout and hashes.

The helper must not replace AI semantic review.

## Commands

```text
agentx detect <target>
agentx list
agentx benchmark
agentx verify <capability> --target <target>
agentx plan install <capability> --target <target>
agentx export <capability> --target <target> --dest <path> --yes
agentx install <capability> --target <target> --dest <path> --yes
agentx rollback <capability> --target <target> --yes
```

`install` and `export` require `--yes`. The helper refuses delivery when the target-ready gate fails.

`<capability>` is a capability id, not a path. It may contain only letters, numbers, dot, dash, and underscore; parent traversal and path separators are refused.

`<target>` is a configured target id under `targets/<target>/`; path-like target values and missing target profiles are refused.

The helper also refuses dangerous destinations: repository root, repository subdirectories, home directory, filesystem root, paths under `.agentx`, symlink destinations, source artifact paths, source descendants, and source ancestors. Existing destinations must be directories.

`benchmark` runs the repeatable AgentX benchmark suite. It creates temporary benchmark capabilities, validates expected pass and block behavior, and removes its temporary `.agentx` artifacts before exiting.

Set `AGENTX_BENCH_SEED=<integer>` to run the randomized target-ready fuzz scenario with a different repeatable seed.

## Target-Ready Gate

The helper checks deterministic evidence only:

- Required workflow files exist: `intake.md`, `open-questions.md`, and `brief.md`.
- Required review files exist.
- Required final review files contain exactly one status line, and that line is `Status: passed`.
- `runtime-benchmark.<target-id>.md` is `Status: passed` or `Status: manual-transcript` for the requested target.
- `unknown-resolution.md` has no unresolved resolution.
- `baseline-deviations.md` exists when any review records `Baseline deviation: yes`.
- Final review files and generated target files contain no `Unknown`, `TBD`, or `TODO` placeholder text, except `unknown-resolution.md` may name an original unknown source fact while recording a resolved status.
- Generated target artifacts contain no symlinks.
- Target artifacts match the target-native package shape.

Passing this gate does not prove semantic correctness. The AI review files are still required authority for meaning, safety, routing quality, and conversion loss.

Install plans include manual requirements from `unknown-resolution.md` entries marked `Resolution: manual`.

For skill-package targets such as Codex, Claude Code, OpenClaw, and Hermes, the target-native package root is `.agentx/output/capabilities/<capability>/targets/<target>/<skill-id>/`. Delivery copies that package root, so the destination receives `SKILL.md` directly rather than an extra nested directory.
