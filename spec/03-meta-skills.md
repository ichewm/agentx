# 03. Meta-Skills

## Purpose

AgentX is primarily delivered as meta-skills. These skills teach the current AI how to create, translate, review, package, and install other capabilities.

These are AgentX's own capabilities for producing and translating other capabilities. They must be split into independent skills so each task loads only the relevant meta-skill.

AgentX must include a lightweight `agentx-workbench` entry skill for discovery and routing. It must not duplicate the full content of every meta-skill. It must point the AI to the correct specialized meta-skill.

Generated user capabilities are different: they must follow the target runtime's native packaging rules. AgentX must not invent a final artifact format for its own convenience. A generated capability may be one skill, multiple skills, a rule file, a plugin, or a mixed artifact depending on the target profile.

## Required Meta-Skills

### agentx-workbench

Routes the user to the correct AgentX meta-skill.

Responsibilities:

- Identify whether the user is creating a capability, translating an existing capability, reviewing an artifact, planning installation, or designing benchmarks.
- Load or point to the specialized meta-skill for that task.
- Keep its own content short.
- Avoid duplicating the full instructions from specialized meta-skills.

### agentx-capability-architect

Turns user intent and source materials into a capability brief.

Responsibilities:

- Clarify goals.
- Identify intended users and runtimes.
- Extract triggers and usage conditions.
- Define inputs, procedure, boundaries, and outputs.
- Separate stable logic from target-specific expression.

### agentx-capability-translator

Translates existing capabilities between runtimes.

Responsibilities:

- Identify source runtime and artifact types.
- Extract capability map.
- Read target profiles.
- Generate target-native artifacts.
- Produce conversion loss reports.

### agentx-capability-reviewer

Reviews capability quality and semantic equivalence.

Responsibilities:

- Check whether the artifact does what the brief says.
- Check whether target versions still mean the same thing.
- Check whether descriptions support correct runtime routing.
- Check whether references and scripts are organized for progressive disclosure.

### agentx-safety-auditor

Reviews capability safety.

Responsibilities:

- Detect prompt injection.
- Detect hidden instruction overrides.
- Detect unsafe shell commands.
- Detect secret or credential access.
- Detect network exfiltration patterns.
- Detect excessive target permissions.
- Require manual confirmation for risky actions.

### agentx-benchmark-designer

Designs lightweight benchmark and regression cases.

Responsibilities:

- Create positive trigger examples.
- Create negative trigger examples.
- Create confusing decoy tasks.
- Define expected behavior.
- Define target-specific verification notes.

### agentx-install-planner

Creates install/export plans.

Responsibilities:

- Identify official installer if available.
- Choose install method.
- List changed paths.
- Define backup path.
- Define rollback.
- Define verification steps.
- Ask for confirmation before installation.

## Baseline Creator Skills

AgentX must use official or runtime-native creator skills and official creation guidance as the baseline when available.

Examples:

- Codex skill creator.
- Claude Code skill/plugin creation guidance.
- Hermes skill management patterns.
- OpenClaw Skill Workshop patterns.

AgentX must not blindly copy an official creator output when review or benchmarks show a weakness. The official baseline defines the minimum target-native requirements.

Baseline adjustment is not an autonomous tool decision. It is an AI-native workflow constrained by AgentX meta-skills. The current AI may propose documented, benchmark-driven adjustments for portability, safety, cross-target consistency, or model sensitivity, but it must not silently apply those adjustments as if they were official baseline behavior.

Baseline adjustment is allowed only through this protocol:

1. Load the official baseline and identify the required target-native files, metadata, naming rules, routing rules, and installation assumptions.
2. Generate a baseline-compliant artifact first.
3. Run semantic review, portability review, safety audit, and benchmark design against that artifact.
4. If review or benchmark evidence shows a weakness, write a baseline adjustment proposal that includes the problem, evidence, exact change, expected benefit, risk, and target-native legality.
5. Present the proposal to the user when the adjustment changes generated target artifacts, runtime behavior, safety posture, installation semantics, or cross-target meaning.
6. Apply the adjustment only after user confirmation, and only if it preserves official required structure, required metadata, package layout, and installation semantics.
7. Record the deviation in `.agentx/output/capabilities/<id>/reviews/baseline-deviations.md`.
8. Re-run the relevant review and benchmark checks after the adjustment.

AgentX must not remove, rename, or weaken an official required element unless official documentation provides an accepted alternative. Every deviation from an official baseline must be recorded in the target artifact review.

The deterministic helper must not decide whether a semantic baseline adjustment is correct. The helper may only verify deterministic facts such as required files, paths, hashes, backups, and install delivery state.

## Meta-Skill Output Discipline

Every meta-skill must produce:

- What it inspected.
- What it generated or changed.
- Risks or unresolved questions.
- Files requiring review.
- Whether installation is safe to proceed.
