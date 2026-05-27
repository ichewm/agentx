# Codex Install

Install methods:

- `path-copy`: copy `<skill-id>/` to `$CODEX_HOME/skills/<skill-id>/`.
- `export-only`: leave generated files under `.agentx/output/capabilities/<id>/targets/codex/`.

Default destination:

```text
$CODEX_HOME/skills/<skill-id>/
```

If `$CODEX_HOME` is unset, the helper may default to `~/.codex`.

Installation requires explicit user confirmation. Verification checks file delivery only; runtime trigger quality requires runtime benchmark evidence.
