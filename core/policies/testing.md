# Policy: Behavior-First Testing

Tests are the executable spec of the domain behaviors. **Gesture coverage is the main signal.** Line coverage is needed but not enough.

## Test names

Use Ubiquitous Language terms. Say the business behavior, not a technical action. Pattern: `should_<outcome>_when_<condition>`.

| Bad | Good |
|-----|------|
| `testCreate()` | `should_issue_invoice_when_mandate_is_active()` |
| `testValidation()` | `should_reject_invoice_when_vat_number_is_missing()` |
| `testDB()` | `should_persist_mandate_when_customer_subscribes()` |

## What each gesture needs

- A test for the happy path.
- A test for each failure in the PO output.
- A test for each invariant in the PO output.
- Consider mutation testing to check test quality.

## Test levels

| Level | Tests | Infrastructure | Use for |
|-------|-------|----------------|---------|
| Domain unit | Aggregates, invariants, value objects | No | Every domain rule |
| Application unit | Use case orchestration | No (ports faked) | Every use case |
| Integration | Driven adapters (real DB, real HTTP) | Yes | Driven adapters |
| Acceptance | A gesture end to end through a driving adapter | Yes | Every gesture, when practical |

## Independence from infrastructure

Domain and application tests never need a database, a real HTTP call, the filesystem, or environment variables. Use in-memory versions of the driven ports.

## Traceability

Each test maps to a behavior in the PO output. The DEV-FS output lists, for each behavior, the tests that cover it. REVIEW checks this list.

```
Behavior: "Issue invoice when mandate is active"
Covered by: InvoiceIssuanceTest#should_issue_invoice_when_mandate_is_active
            InvoiceIssuanceTest#should_reject_invoice_when_mandate_is_revoked
```

## Fast and slow checks

- Early and often: linter, type checker, unit tests, coverage.
- At the end: the REVIEW agent judges gesture coverage and risks.

## Forbidden

| Anti-pattern | Instead |
|--------------|---------|
| Test named after a method | Test named after a behavior |
| One test for several behaviors | One test per behavior |
| Assertions on private state | Assertions on visible domain results |
| Test that passes with an empty implementation | Run it first and see it fail |
| Trivial assertions to raise coverage | Every test checks a real domain result |
