# REVIEW Output Template

> Fill it after reviewing the DEV-FS output. It feeds G3.
> Start the saved file with: `> Agent: REVIEW · Model: <model-id> · Gesture: <gesture> · Pass: <n>`

---

## Gesture: `[Name, same as the PO and DEV-FS outputs]`

**Bounded context**: `[Context]`

---

## Verdict

**`ACCEPTED`** | **`NEEDS_REVISION`** | **`REJECTED`**

> One sentence: the verdict and the main reason.

---

## Line coverage

| Scope | Coverage | Threshold | Status |
|-------|----------|-----------|--------|
| Domain | `XX%` | `.ai/local/` (default 90%) | ✅ / ❌ |
| Application | `XX%` | | ✅ / ❌ |
| Overall | `XX%` | | ✅ / ❌ |

**Notes**: why coverage is low, if it is (generated code, untested adapters).

---

## Gesture coverage

| Gesture | Happy path | Failures | Invariants | Status |
|---------|:----------:|:--------:|:----------:|--------|
| `[Name]` | ✅ | ✅ | ✅ | ✅ Covered |
| `[Name 2]` | ✅ | ❌ | ✅ | ⚠️ Partial |
| `[Name 3]` | ❌ | ❌ | ❌ | ❌ Not covered |

Total: `N` · Covered: `N` · Partial: `N` · Not covered: `N`

---

## Residual risks

| # | Risk | Level | Gesture | Fix |
|---|------|-------|---------|-----|
| 1 | `[Invariant X has no failure test]` | HIGH | `[Name]` | Add a test for `[condition]` |
| 2 | `[Failure Y not tested]` | MEDIUM | `[Name]` | Add a test for `[condition]` |
| 3 | `[Edge case Z]` | LOW | `[Name]` | Consider a test |

HIGH: `N` · MEDIUM: `N` · LOW: `N`

---

## Architecture notes

Information only, not a gate criterion.

| Note | Where | Level |
|------|-------|-------|
| | | |

---

## Fixes needed (if not `ACCEPTED`)

In priority order:

1. **[HIGH]** `[Action]` → risk #N
2. **[MEDIUM]** `[Action]`
3. **[LOW]** `[Action]`

---

## Input for G3

| Check | Status |
|-------|--------|
| Every gesture has at least one test | ✅ / ❌ |
| No HIGH or MEDIUM risk | ✅ / ❌ |
| Line coverage at or above the threshold | ✅ / ❌ |
| Failure paths tested | ✅ / ❌ |

**Recommended action**: `ACCEPT` | `REVISE with DEV-FS` | `REJECT and reopen PO`
