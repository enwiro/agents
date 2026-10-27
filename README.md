# Agents Guide

Four AI agents help you build a business feature the DDD way. They write code, tests and documents. You decide at three key moments.

## What they are for

| Command | Agent | Use it to |
|---------|-------|-----------|
| `/architect <request>` | Architect | Build a new feature, start to end. **Start here.** |
| `/po <request>` | Product Owner | Clarify a business rule or the Ubiquitous Language |
| `/dev-fs <PO output>` | Developer | Build a gesture that is already specified |
| `/review <scope>` | Reviewer | Check test coverage of existing code |

A **gesture** is one named business action, like "issue an invoice".

## How it works

```
You ──► /architect "<request>"
          │
   G0  Plan        You and the agents agree on scope        ◄─ you decide
   G1  Behaviors   Rules, language, edge cases              ◄─ you decide
          │
   Build           One commit per gesture (test, code, check, commit)
   G2  Check       The Architect checks the architecture
   Review          A fresh agent checks coverage and risks
   G3  Verdict     Result and open points                   ◄─ you decide
```

- **G0 and G1:** the agents ask all their questions here. Answer well and the rest runs alone.
- **After G1:** the agents ask only if they are blocked. Their other choices are listed at G3.
- **Not accepted at G3?** The agents fix and review again, until the review passes.
- **Models:** the Architect uses the strongest model. The other agents use the same or a smaller one.

## Use it

```
/architect Add a way to reject an invoice with an invalid VAT number. Ticket CM-300, in furry.
/po Clarify what "invoice mandate" means in invoicing.
/dev-fs Build the gesture in documentation-agent/reject-invoice/po.md
/review Check coverage of the invoicing module.
```

You can write in French. Code and files are in English.

To skip a gate, say so: "skip G1". The gate is then skipped completely.

## Commits and push

- One commit per gesture, for example `feat(CM-300): reject invoice when VAT number is invalid`.
- Commits are local and not signed. At the end, sign them in one step:

```sh
git rebase <base> --exec 'git commit --amend --no-edit -S'
```

- Agents never push. You approve the push at G3.

## Where to find the results

Everything is in `documentation-agent/` at the root of the target repo.

| File | Content |
|------|---------|
| `README.md` | One line per gesture: verdict, passes, what was built |
| `<feature>/plan.md` | The approved plan |
| `<gesture>/po.md` | Behaviors and language |
| `<gesture>/dev-fs.md` | How it was built, tests, commits |
| `<gesture>/review-<n>.md` | Review of each pass |
| `<gesture>/architect-summary.md` | Verdicts and decisions |
| `<gesture>/friction-log.md` | Problems and choices made without you |

## Install

Full steps are in `adapters/claude-vscode/AGENTS.md` (Claude Code) and `HARNESS.md` (new repo, Cursor). In short, from the parent folder of your repos:

```sh
mkdir -p .claude/commands
for cmd in architect po dev-fs review; do
  ln -sfn "../../einvoicing-documentation/agents/adapters/claude-vscode/commands/$cmd.md" ".claude/commands/$cmd.md"
done
make -C einvoicing-documentation/agents/mcp/ddd build
ln -sfn einvoicing-documentation/agents/adapters/claude-vscode/mcp.json .mcp.json
```

Restart Claude Code. The `ddd` MCP server gives agents the DDD book by section. This costs far fewer tokens than the PDF.

## Per-repo settings

Put these files in `.ai/local/` of a repo. They win over the shared rules.

| File | Purpose |
|------|---------|
| `context.md` | Bounded context, team, stack |
| `ubiquitous-language.md` | Glossary |
| `coverage-thresholds.md` | Coverage targets (default 90%) |
| `constraints.md` | Extra rules |
