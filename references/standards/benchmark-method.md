# AgentX Benchmark Method

AgentX uses workflow-level benchmarks, not only output scoring.

## Borrowed Evaluation Ideas

- skills.sh documents skills as reusable agent capabilities and warns users to review skills before installing because quality and security are not guaranteed: `https://www.skills.sh/docs`
- Static skill security scanners check prompt injection, credential leaks, malicious code, and social engineering patterns before installation, as in skill-issue: `https://www.skill-issue.sh/`
- Skills security benchmarks such as OASB emphasize labeled attack categories, precision, recall, and false-positive control: `https://www.oasb.ai/benchmark`
- AgentDojo-style benchmarks separate ordinary user tasks, injected adversarial instructions, and utility-under-attack outcomes: `https://agentdojo.spylab.ai/`
- SkillsBench-style task benchmarks measure whether a skill improves real task completion across agent/model configurations: `https://www.skillsbench.ai/`
- Trigger testing for skills must include positive cases, negative cases, decoys, and routing confusion tests.

## AgentX Benchmark Layers

1. Static structure: required files, target-native layouts, status lines, and no unresolved placeholders.
2. Workflow state: intake closure, target-ready gate, install confirmation, and baseline deviation records.
3. Semantic review artifacts: semantic, portability, safety, unknown-resolution, and conversion loss.
4. Runtime benchmark artifacts: automated run result or manual transcript.
5. Adversarial safety: prompt injection, secret access, exfiltration, destructive commands, and hidden instruction override.
6. Cross-target drift: preserved/adapted/degraded/manual/unsupported behavior across targets.

## Required Automated Scenarios

- clean skill package passes target-ready gate
- missing workflow artifacts block finalization
- missing review blocks finalization
- missing final review status blocks finalization
- final review status that still needs a user decision blocks finalization
- multiple final review status lines block finalization
- blocked review blocks finalization
- unresolved placeholder text in final reviews blocks finalization
- unresolved Unknown resolution blocks finalization
- target placeholder blocks finalization
- target artifact symlinks block finalization and delivery
- target-native package shape mismatches block finalization
- runtime benchmark blocking gap blocks finalization
- manual runtime transcript allows finalization
- runtime benchmark evidence is target-specific, so one blocked target does not block another passing target
- unsafe safety review blocks finalization
- Cursor rule artifact passes without fake `SKILL.md`
- baseline deviation marker requires `baseline-deviations.md`
- install/export refuses execution without explicit confirmation
- install/export refuses dangerous destinations such as repository root, repository subdirectories, home, filesystem root, `.agentx`, symlink destinations, and source-overlapping paths
- helper commands refuse capability ids that look like paths or parent traversal
- helper commands refuse target ids that look like paths, parent traversal, or missing target profiles
- ready artifact export copies files and writes a lock record
- rollback restores the latest lock-recorded backup
- ready artifact creates a plan without executing installation
- install plans include manual requirements from `unknown-resolution.md`
- every AgentX meta-skill has clear frontmatter, an actionable operational section, and authoritative spec references
- the workbench skill stays a lightweight router instead of duplicating specialized instructions
- translation coverage includes preserved, adapted, degraded, dropped, manual setup, and risk reporting
- capability architecture coverage includes multi-turn intake, explicit closure, open questions, and decision points before final generation
- fixed-seed randomized target-ready fuzzing combines edge cases so gate behavior remains stable under mixed invalid states
- external AI runtime CLI benchmarks require explicit user confirmation before execution

## Real Corpus Strategy

Use skills.sh as a source of real-world skill shapes and risk patterns:

- high-install official skills
- community skills with scripts
- document-only skills
- workflow/planning skills
- security/audit skills
- frontend/design skills
- cloud/provider skills

Do not copy external skill content into AgentX unless license and provenance are verified. Use local synthetic fixtures for repeatable CI, and use real corpus URLs for periodic manual or networked benchmark review.
