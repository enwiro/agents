# DEV-FS Output Template

> Fill it after building the gestures. Every field is required.
> Start the saved file with: `> Agent: DEV-FS · Model: <model-id> · Gesture: <gesture>`

---

## Gesture: `[Name, same as the PO output]`

**Bounded context**: `[Context]`
**Ticket**: `[CM-300 or none]`

---

## Approach

2 to 5 sentences: how the gesture is built, which building blocks, any notable choice.

---

## Architecture mapping

### Domain

| Building block | Name (UL term) | Responsibility |
|----------------|----------------|----------------|
| Aggregate root | `AggregateRootName` | |
| Value object | `ValueObjectName` | |
| Domain event | `EventName` | |
| Domain service | `ServiceName` (if needed) | |

### Application

| Component | Name | Responsibility |
|-----------|------|----------------|
| Application service | `UseCaseName` | Runs the gesture |
| Driving port | `CommandHandlerName` | Entry for driving adapters |
| Driven port | `RepositoryName` | Persistence interface |

### Adapters

| Type | Name | Technology |
|------|------|------------|
| Driving | `ControllerName` | REST / CLI / consumer |
| Driven | `RepositoryImplName` | PostgreSQL / in-memory |

---

## Package structure

```
src/
  [bounded-context]/
    [aggregate-or-subdomain]/
      domain/
      application/
        ports/
      adapters/
        driving/
        driven/
```

---

## Tests

For each PO behavior, the tests that cover it.

### Behavior: `[Name]`

| Test | Type | Covers |
|------|------|--------|
| `should_[outcome]_when_[condition]` | domain unit | [rule or invariant] |
| `should_[failure]_when_[condition]` | domain unit | [failure from the PO output] |
| `should_[outcome]` | integration | [driven adapter behavior] |

### Traceability

| PO behavior | Tests | Status |
|-------------|-------|--------|
| `[Name]` | `TestClass#method` | ✅ Covered |
| `[Name 2]` | — | ⚠️ Missing |

---

## Commits

| Commit | Gesture |
|--------|---------|
| `feat(CM-300): [subject]` | `[Gesture]` |

---

## Architecture checklist

- [ ] The domain imports nothing from adapters or frameworks.
- [ ] All DB access goes through a driven port.
- [ ] All HTTP in goes through a driving adapter.
- [ ] Domain tests use in-memory adapters.
- [ ] Package names use UL terms.
- [ ] No business logic in utility packages.
- [ ] UL terms appear exactly as written in code.
- [ ] Code comments are short and self-contained.

---

## Deviations

Anything that breaks a policy, with the reason. The Architect checks it at G2.

| # | Deviation | Reason | Needs Architect approval |
|---|-----------|--------|--------------------------|
| 1 | | | YES |
