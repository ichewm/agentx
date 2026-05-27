# 01. AI-Native Workflow

## Primary User Experience

The primary interface is conversation with an AI runtime that has AgentX loaded.

AgentX supports two primary user flows:

1. Create a new capability from zero.
2. Convert an existing capability into equivalent or approximate target artifacts.

The user may start either flow with a command-like natural language request.

Creation example:

```text
Use AgentX. Build a capability from the content below for Codex and Claude Code.

<user-provided requirements, documents, examples, or notes>
```

The user does not need to provide all material in one turn. AgentX must support multi-turn intake.

Conversion example:

```text
Use AgentX. Convert this Claude Code plugin into Codex and Cursor artifacts.
Preserve behavior where possible and tell me what cannot be translated.
```

AgentX meta-skills guide the current AI through intake, synthesis, review, and delivery.

## Interaction Protocol

The AI must not require the user to paste the same source repeatedly.

For source materials, the AI must accept any of:

- Inline pasted content.
- A local path.
- A GitHub or web URL.
- A directory containing a skill/plugin/rule pack.
- Multiple documents or examples.

When the source is a path or URL, the AI must inspect it directly when tools allow. If direct access is not available, the AI must ask the user for the smallest missing input.

The AI may ask follow-up questions during creation or conversion, but questions must be targeted. It must avoid broad open-ended questionnaires.

The AI must keep a work record under `.agentx/output/capabilities/<id>/` so later turns can continue without re-pasting source material.

## Multi-Turn Intake

Creation and conversion are intake sessions, not one-shot commands.

When a user starts a capability task, the AI must create or choose a capability id and maintain:

```text
.agentx/output/capabilities/<id>/
  intake.md
  sources/
  open-questions.md
  brief.md
```

The user may then continue with messages such as:

```text
Add this requirement...
Here is another document...
Treat this example as a negative case...
Also support Cursor later, but not in this first output...
```

For each new input, the AI must:

1. Append or summarize the material in `intake.md`.
2. Store large pasted material or fetched files under `sources/`.
3. Update `open-questions.md` without forcing the user to answer immediately unless the next intake step is blocked.
4. Update `brief.md` only when the intent is clear enough.
5. Report what changed and what remains unresolved.
6. Ask whether the user wants to continue providing input.

The continuation prompt must be explicit and short:

```text
Do you want to add more input? Reply "no", "done", "不", or "不需要" to close intake and continue.
```

While intake is open, the AI must primarily collect and summarize. It must not start deep critique, target generation, or installation planning unless the user explicitly asks for an early draft. Questions recorded during intake must be held for the intake-close decision phase unless they are required to understand the next user-provided source.

When the user closes intake, the AI must:

1. Summarize the accumulated intent and sources.
2. Present specific questions that require user decisions.
3. Provide concrete options when a decision is needed.
4. Wait for the user's decision before generating final target artifacts.

The AI must not generate final target artifacts until either:

- The user closes intake and confirms generation.
- The user asks it to generate an early draft now.

If information is incomplete, the AI must produce a partial brief and a short list of specific missing items.

## Main Workflow

1. Ingest
   - Read user intent, documents, existing skills, plugins, rules, repositories, or examples.
   - Identify source runtime and source artifact types.

2. Map
   - Build a capability map.
   - Identify skills, commands, tools, connectors, subagents, references, assets, scripts, rules, and installation assumptions.

3. Brief
   - Produce a platform-neutral capability brief.
   - Record goal, usage conditions, inputs, procedure, boundaries, examples, and conversion notes.

4. Target Plan
   - Read target profiles.
   - Decide which outputs are direct, adapted, degraded, unsupported, or manual.

5. Generate
   - Use target creator baselines and official conventions to produce artifacts.
   - Prefer native target formats.

6. Review
   - Run semantic review, portability review, safety audit, and benchmark design.
   - Revise artifacts until review issues are resolved or explicitly accepted.

7. Delivery Plan
   - Produce install/export plans.
   - Show source artifact, destination, method, changes, backup path, verification, and rollback.

8. Install or Export
   - Use the official target installer when available.
   - Use the thin helper only for deterministic file delivery and verification.

## Creation Workflow

Creation starts from a user request and a source bundle. The source bundle may be empty. Creation does not start from a traditional `agentx new` command.

```text
Use AgentX. Build a reusable database migration safety capability for Codex and Claude Code.
Here are our migration rules and examples...
```

The current AI must:

- Ask targeted questions if the request is underspecified.
- Save source notes and materials under `.agentx/output/capabilities/<id>/sources/`.
- Maintain `intake.md` across turns.
- Draft a capability brief.
- Generate target artifacts.
- Run reviews.
- Produce install plans.

## Translation Workflow

Translation starts from an existing capability.

```text
Use AgentX. Translate this OpenClaw skill pack into Claude Code and Codex artifacts.
```

The current AI must:

- Accept a path, URL, repository, archive, or pasted content as the source.
- Identify source format.
- Copy or summarize source materials under `.agentx/output/capabilities/<id>/sources/`.
- If source material is pasted across multiple turns, preserve each input as an intake entry and avoid asking for the same content again.
- Preserve intent before preserving file shape.
- Reuse official target creator guidance.
- Report conversion loss.
- Avoid pretending unsupported runtime features were preserved.

## Runtime Selection

AgentX does not select skills during end-user tasks when the target runtime has native discovery.

For runtimes such as Codex, Claude Code, Copilot, OpenClaw, or Hermes Agent, AgentX prepares the artifacts and the runtime chooses when to use them.

Fallback selection is allowed only for targets without native skill or rule discovery, and must be described as a fallback.
