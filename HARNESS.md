# DDD Agent Harness — Setup and Maintenance

A shared multi-agent harness for domain-driven development, for **Claude Code** and **Cursor**.

This folder is the single source for agent roles, rules, policies and templates. Each repo reuses it through symlinks and can add local overrides.

---

## How It Works

Role files, policies and templates guide the agents before they act. The REVIEW agent and the G3 verdict check the result after.

```
/architect
  G0 plan (human)  →  G1 behaviors (human)  →  DEV-FS, one commit per gesture
  →  G2 architecture check  →  REVIEW  →  G3 verdict (human)
```

See `core/workflow.md` for the rules.

---

## Onboarding a New Repository

### Step 1 — Create the `.ai/` directory

```bash
mkdir -p <your-repo>/.ai
```

### Step 2 — Add symlinks

From the root of your repository:

```bash
# Symlink to shared harness core
ln -s ../einvoicing-documentation/agents/core .ai/agents-core

# Symlink to shared adapters
ln -s ../einvoicing-documentation/agents/adapters .ai/adapters
```

> Adjust the relative path (`../einvoicing-documentation/`) to match the actual relative position of this repository from your repo root.

### Step 3 — Wire up your tool

**Claude (VS Code / Claude Code CLI)**:
```bash
ln -s .ai/adapters/claude-vscode/AGENTS.md AGENTS.md
```

**Cursor**:
```bash
mkdir -p .cursor/rules
ln -s ../../.ai/adapters/cursor/rules.mdc .cursor/rules/ddd-harness.mdc
```

### Step 4 — Create local overrides (optional but recommended)

```bash
mkdir -p .ai/local
```

Add any of these files as needed:

| File | Purpose |
|------|---------|
| `.ai/local/context.md` | Bounded context description, team, tech stack |
| `.ai/local/ubiquitous-language.md` | Living domain glossary for this repo |
| `.ai/local/coverage-thresholds.md` | Line coverage targets (overrides workflow.md defaults) |
| `.ai/local/constraints.md` | Additional architectural or technical constraints |

Local overrides take **precedence** over shared rules. Use them for repo-specific decisions.

### Step 5 — Verify

```bash
# Check symlinks resolve
ls -la .ai/
cat .ai/agents-core/workflow.md
cat .ai/adapters/claude-vscode/AGENTS.md
```

---

## Using the Agents

```
/architect <the business gesture or feature to build>
/po        <domain area to clarify>
/dev-fs    <PO output to implement>
/review    <scope to check>
```

More in [README.md](README.md).

---

## DDD book MCP

`mcp/ddd/` (next to this file) holds a small Go MCP server with the DDD book split by section. Build it with `make -C agents/mcp/ddd build`. Setup is in `adapters/claude-vscode/AGENTS.md`.

---

## Override Precedence

```
.ai/local/          ← highest priority (repo-specific)
core/                ← default (shared across all repos)
```

When a local file conflicts with a shared policy, the local file wins. The Architect enforces this order automatically.

---

## Language Convention

- Code, tests, documentation, and all agent outputs: **English**.
- Ubiquitous Language terms: **English** (verbatim in code).
- Conversational exchanges with the user: French is fine.

---

## Contributing to the Harness

Changes to `core/` affect all consuming repositories. Before modifying shared files:

1. Test the change in a single repo's `.ai/local/` override first.
2. Once validated, promote to `core/` with a clear commit message.
3. Communicate significant changes to all repo owners.

No formal versioning. Git history is the audit trail.
