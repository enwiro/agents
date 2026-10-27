# PO Output Template

> Fill it for each gesture. Every field is required.
> Start the saved file with: `> Agent: PO · Model: <model-id> · Gesture: <gesture>`

---

## Gesture: `[Name in Ubiquitous Language]`

**Bounded context**: `[Context]`
**Actor**: `[Who triggers it]`
**Domain event**: `[EventName, past tense]`

---

## Ubiquitous Language

| Term | Definition | Bounded context | Notes |
|------|------------|-----------------|-------|
| `TermA` | Plain business definition | `Context` | |
| `TermB` | Plain business definition | `Context` | Differs from `OtherTerm` in `OtherContext` |

---

## Behaviors

### Behavior: `[Name]`

**Description**: One sentence on what happens when it succeeds.

**Preconditions**:
- ...

**Command payload**:
- `field_name` (type): description

**Postconditions**:
- ...

**Rules and invariants**:
- ...

**Failures**:
| Condition | Result |
|-----------|--------|
| `[condition]` | `[domain error or reason]` |

**Examples**:

*Happy path*:
> Given [context], when [actor] does [action], then [result].

*Failure path*:
> Given [context], when [actor] does [action] with [invalid input], then [failure result].

---

## Side effects

| Other aggregate or context | Effect | How (sync or event) |
|----------------------------|--------|---------------------|
| | | |

---

## Open questions

All answered before G1 closes.

| # | Question | Answer | Answered by |
|---|----------|--------|-------------|
| 1 | | | |
