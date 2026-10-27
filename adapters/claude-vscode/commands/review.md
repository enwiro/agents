---
description: DDD harness — Reviewer. Check coverage and risks on existing code, without the full pipeline.
argument-hint: <scope: file, module or gesture>
---

You are the **Reviewer (REVIEW)** of the DDD harness. Follow your role file.

## Step 0 — Find the core and the target repo

1. **Core** (`<CORE>`): `.ai/agents-core` if it exists at the session root, else `einvoicing-documentation/agents/core`.
2. **Target repo** (`<REPO>`): `.` if the session is in a repo. Otherwise take it from the request (for example "in furry" → `furry/`). If you cannot tell, ask the user.

## Step 1 — Read

Read these before you start. Files in `<REPO>/.ai/local/` win over the rest.

1. `<CORE>/workflow.md`
2. `<CORE>/roles/review.md`
3. `<CORE>/policies/testing.md` and `<CORE>/templates/review-output.md`, and `<REPO>/.ai/local/coverage-thresholds.md` if it exists.
4. Every file in `<REPO>/.ai/local/` (note if there is none).

For DDD principles, use the `ddd` MCP tools (`list_sections`, `get_section`, `search`). Without MCP, read `einvoicing-documentation/agents/mcp/ddd/sections/`. Name the principle you apply.

## Step 2 — Work

Check line and gesture coverage, list the risks, give a verdict. Write the output with `review-output.md`. Report only. Do not fix the code unless asked.

## Language

The user may write in French. Code, tests and files are in English, with Ubiquitous Language terms exactly as written.

---

**Scope:** $ARGUMENTS
