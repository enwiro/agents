# DDD Agent Harness — Claude Code

Entry point for Claude Code (VS Code extension or CLI). For a short guide, read `agents/README.md` in `einvoicing-documentation`.

Claude Code does not read `AGENTS.md` by itself. The commands `/architect`, `/po`, `/dev-fs` and `/review` work only after you install them as slash commands.

## Install

Claude Code reads `.claude/commands/` and `.mcp.json` **at the session root** only. Where you link them depends on how you open Claude Code.

**A. You open one repo.** Run it at the repo root:

```sh
mkdir -p .claude/commands
for cmd in architect po dev-fs review; do
  ln -sfn "../../.ai/adapters/claude-vscode/commands/$cmd.md" ".claude/commands/$cmd.md"
done
```

**B. You open the parent folder that holds the repos and `einvoicing-documentation`.** Run it in the parent folder:

```sh
mkdir -p .claude/commands
for cmd in architect po dev-fs review; do
  ln -sfn "../../einvoicing-documentation/agents/adapters/claude-vscode/commands/$cmd.md" ".claude/commands/$cmd.md"
done
ln -sfn einvoicing-documentation/agents/adapters/claude-vscode/mcp.json .mcp.json
```

Both can exist together. Restart the session to load new commands.

The commands find the core and the target repo on their own: `.ai/agents-core` and `.` in case A, `einvoicing-documentation/agents/core` and the repo named in the request in case B.

## DDD book (MCP)

The `ddd` MCP server gives agents the DDD book by section, which costs far fewer tokens than the PDF. Build it once:

```sh
make -C einvoicing-documentation/agents/mcp/ddd build
```

Case B loads it from `.mcp.json`. For case A, or for any folder, register it once per machine:

```sh
claude mcp add --scope user ddd -- "$(pwd)/einvoicing-documentation/agents/mcp/ddd/ddd-mcp"
```

Without the MCP, agents read `einvoicing-documentation/agents/mcp/ddd/sections/` directly.

## What each command does

| Command | Role file | Output template |
|---------|-----------|-----------------|
| `/architect` | `roles/architect.md` | none |
| `/po` | `roles/po.md` | `templates/po-output.md` |
| `/dev-fs` | `roles/dev-fs.md` | `templates/dev-fs-output.md` |
| `/review` | `roles/review.md` | `templates/review-output.md` |

All paths are under `.ai/agents-core/`. The pipeline is in `workflow.md`.

## Read order

Later entries win:

1. `policies/`
2. `workflow.md`
3. `.ai/local/` in the target repo

## Local files (`.ai/local/`)

| File | Purpose |
|------|---------|
| `context.md` | Bounded context, team, stack |
| `ubiquitous-language.md` | Glossary of the repo |
| `coverage-thresholds.md` | Line coverage targets |
| `constraints.md` | Extra architecture or technical rules |
