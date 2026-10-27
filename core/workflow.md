# Workflow

## Commands

| Command | Use it to |
|---------|-----------|
| `/architect <request>` | Build a new feature or gesture (full pipeline) |
| `/po <request>` | Explore a domain area or clarify the Ubiquitous Language |
| `/dev-fs <PO output>` | Implement a gesture that is already specified |
| `/review <scope>` | Check coverage of existing code |

A **gesture** is one named business action with a visible effect in the domain.

## Pipeline

```
Request
  → G0  Plan          Architect + PO + human      (human stop)
  → G1  Behaviors     PO                          (human stop)
  → Implementation    DEV-FS, one commit per gesture
  → G2  Architecture  Architect checks            (no stop)
  → Review            REVIEW
  → G3  Verdict       Architect + human           (human stop)
        ↑ not accepted: fix with DEV-FS (or PO), then a fresh REVIEW
```

The human is needed at G0, G1 and G3 only. Ask everything at G0 and G1 so that nothing is left open later.

## Sub-agents

- The Architect runs every PO, DEV-FS and REVIEW phase in a separate sub-agent. It never plays those roles itself.
- Each sub-agent has its own context. REVIEW is always a fresh agent.
- **Model:** a sub-agent uses the Architect's model or a smaller one, never a larger one. The Architect chooses for each task.
  - Default: one tier smaller.
  - Same model: complex work (new aggregate, linked invariants, financial or legal topic, large refactor, a review that already failed once).
  - The Architect writes the reason in one line in `architect-summary.md`.

## Phases

### G0 — Plan
The Architect and a PO sub-agent turn the request into a plan with the human. The plan is written in `documentation-agent/<feature>/plan.md` while it is built, so the human can read it at any time.

The plan has:
- The request, in the Architect's words.
- The gestures, and whether to process them in sequence or in parallel.
- Scope and out of scope.
- The ticket ID (for example `CM-300`). Ask once. If there is none, take it from the request or the branch name; if still none, leave it out.
- A high-level approach: bounded context and architecture direction, not a detailed design.
- Open questions.

**Pass when:** the human approved the plan and no question is open.

### G1 — Behaviors
A PO sub-agent writes the behaviors and the Ubiquitous Language (UL). The Architect asks the human every question about edge cases, invariants and failures before G1 closes.

**Pass when:**
- The UL is complete.
- Every behavior has invariants, a happy-path example and a failure-path example.
- No open question remains.
- The human approved.

### Implementation — DEV-FS
A DEV-FS sub-agent builds the gestures one by one: test, code, check, commit. Details are in `roles/dev-fs.md`. Commits follow `policies/commits-and-comments.md`.

### G2 — Architecture check
The Architect reads the DEV-FS output and the code, with no human stop.

**Pass when:**
- Every behavior is implemented.
- Every behavior has tests: happy path, each failure, each invariant.
- The architecture mapping is complete.
- There is no domain logic in adapters.
- There is one commit per gesture.

### Review — REVIEW
A fresh REVIEW sub-agent checks line coverage and gesture coverage, and lists residual risks.

### G3 — Verdict
The Architect gives the final verdict, then the human decides.

**Pass when:**
- REVIEW returned `ACCEPTED`.
- The Architect approves.
- The human approved.

The G3 report is short: what was built, commits, review verdict, revision count, open risks, and the choices made without the human (from the friction log). End it with: *accept, revise, or stop?*

On accept, the Architect reminds the human how to sign the commits, then asks for approval to push. Nobody pushes without it.

## Revision loop

When REVIEW returns `NEEDS_REVISION` or `REJECTED`, or the Architect rejects:
- Gap in code or tests: spawn a new DEV-FS with the findings, then a new REVIEW.
- Gap in the spec: spawn a new PO (back to G1), then DEV-FS, then REVIEW.

Repeat until REVIEW returns `ACCEPTED` and the Architect approves. Fixes are commits too.

After 3 passes without progress, ask the human. The final summary gives the number of passes and what each one fixed.

## Questions after G1

Stop and ask only when a doubt blocks the work and no safe default exists. Ask one grouped question with a recommendation.

Otherwise pick the safest option, log it in `friction-log.md` with the reason, and report it at G3.

## Skipping gates

Gates are never skipped by default. If the human asks, the Architect skips the named gate completely (no stop, no check). There is no special procedure.

## Parallel gestures

Independent gestures (no shared aggregate or invariant) may run in parallel, each in its own pipeline. The Architect decides. When in doubt, run them in sequence.

## Files

Every agent writes to `<REPO>/documentation-agent/`. The Architect creates the folder if needed.

| Who | File |
|-----|------|
| Architect + PO | `<feature>/plan.md` |
| PO | `<gesture>/po.md` |
| DEV-FS | `<gesture>/dev-fs.md` |
| REVIEW | `<gesture>/review-<n>.md` (one per pass, never overwritten) |
| Architect | `<gesture>/architect-summary.md` |
| Any agent | `<gesture>/friction-log.md` (append only) |
| Architect | `README.md` (roll-up of all gestures) |

Each file starts with: `> Agent: <role> · Model: <model-id> · Gesture: <gesture>`.

**Roll-up `README.md`:** short and kept up to date, never wiped. One line per gesture: name, ticket, bounded context, G3 verdict, number of passes, models used, and what was built. Then two lines per agent on what it did in this run.

**Friction log:** add a dated entry when something cost time: a wrong claim found by checking, a repeated correction, an unclear instruction, a tool surprise, or a choice made without the human. Say enough for a future agent to avoid it. Read the logs of earlier gestures before planning a new one.

## Precedence and DDD reference

Files in `<REPO>/.ai/local/` win over this workflow.

Agents check their choices against the DDD book with the `ddd` MCP tools (`list_sections`, `get_section`, `search`). Without the MCP, read `einvoicing-documentation/agents/mcp/ddd/sections/`. Name the principle when you apply it.
