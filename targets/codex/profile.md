# Codex Target Profile

Status: partially verified from local Codex runtime behavior and OpenAI skill documentation.

Sources:

- OpenAI Help: `https://help.openai.com/en/articles/20001066-skills-in-chatgpt`
- Local Codex runtime instruction: Codex skills are installed under `$CODEX_HOME/skills`.

## Runtime Capabilities

- Native skills: supported through Agent Skills / `SKILL.md` packages.
- Rules: no separate Codex rule artifact is verified in this profile.
- Prompts: supported as skill instructions and bundled references, not as a separate target artifact here.
- Commands: no target-native command package is verified in this profile.
- Tools: supported only when exposed by the current runtime; skills can instruct tool use but cannot install tools into the model.
- Connectors: no target-native connector package is verified in this profile.
- Hooks: no target-native hook package is verified in this profile.
- Subagents: no target-native subagent package is verified in this profile.

## Routing

- Runtime chooses capabilities: yes.
- Routing source: skill `name`, `description`, and `SKILL.md` body.

## Scopes

- Project: not verified.
- User: `$CODEX_HOME/skills/<skill-id>/`.
- Organization: not verified.
- Plugin: plugins can provide skills, but AgentX does not generate Codex plugins by default.

## Refresh Requirements

Restart or refresh behavior is runtime-dependent and not guaranteed by this profile.

## Unsupported Features

AgentX must not claim that a Codex skill installs new model weights, grants new tools, or overrides runtime tool permissions.
