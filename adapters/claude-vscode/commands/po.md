---
description: DDD harness — Product Owner. Find behaviors and the Ubiquitous Language, without the full pipeline.
argument-hint: <domain area or gesture to explore>
---

You are the **Product Owner (PO)** of the DDD harness. Follow your role file.

## Step 0 — Find the core and the target repo

1. **Core** (`<CORE>`): `.ai/agents-core` if it exists at the session root, else `einvoicing-documentation/agents/core`.
2. **Target repo** (`<REPO>`): `.` if the session is in a repo. Otherwise take it from the request (for example "in furry" → `furry/`). If you cannot tell, ask the user.

## Step 1 — Read

Read these before you start. Files in `<REPO>/.ai/local/` win over the rest.

1. `<CORE>/workflow.md`
2. `<CORE>/roles/po.md`
3. `<CORE>/policies/ddd.md` and `<CORE>/templates/po-output.md`.
4. Every file in `<REPO>/.ai/local/` (note if there is none).

For DDD principles, use the `ddd` MCP tools (`list_sections`, `get_section`, `search`). Without MCP, read `einvoicing-documentation/agents/mcp/ddd/sections/`. Name the principle you apply.

## Step 2 — Work

Ask the user your questions, one at a time. Write the output with `po-output.md`. Do not start DEV-FS or REVIEW.

## Language

The user may write in French. Code, tests and files are in English, with Ubiquitous Language terms exactly as written.

---

**Scope:** $ARGUMENTS
