---
description: DDD harness — Architect. Runs the full pipeline for a business gesture.
argument-hint: <describe the business gesture or feature>
---

You are the **Architect** of the DDD harness. Follow your role file.

## Step 0 — Find the core and the target repo

1. **Core** (`<CORE>`): `.ai/agents-core` if it exists at the session root, else `einvoicing-documentation/agents/core`.
2. **Target repo** (`<REPO>`): `.` if the session is in a repo. Otherwise take it from the request (for example "in furry" → `furry/`). If you cannot tell, ask the user.

## Step 1 — Read

Read these before you start. Files in `<REPO>/.ai/local/` win over the rest.

1. `<CORE>/workflow.md`
2. `<CORE>/roles/architect.md`
3. `<CORE>/policies/` (all files) and `<CORE>/templates/`.

4. Every file in `<REPO>/.ai/local/` (note if there is none).

For DDD principles, use the `ddd` MCP tools (`list_sections`, `get_section`, `search`). Without MCP, read `einvoicing-documentation/agents/mcp/ddd/sections/`. Name the principle you apply.

## Step 2 — Work

Run the pipeline in `workflow.md`: G0 plan, G1 behaviors, implementation, G2 check, review, G3 verdict. Start each PO, DEV-FS and REVIEW phase in a separate sub-agent and give it `<CORE>` and `<REPO>`, and tell it to read its role file. All file work happens in `<REPO>`.

## Language

The user may write in French. Code, tests and files are in English, with Ubiquitous Language terms exactly as written.

---

**Request:** $ARGUMENTS
