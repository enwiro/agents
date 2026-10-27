# Policy: Architecture — Hexagonal + Screaming

## Hexagonal (ports and adapters)

The business logic is at the center. Databases, HTTP, messaging and files reach it only through interfaces. The core never reaches outward.

```
 Driving side (REST, CLI, consumers, tests)
        │
 Driving adapters (controllers)
        │ call
 Driving ports (input interfaces)
        │
 ╔══════▼═══════════════════════╗
 ║ DOMAIN CORE                  ║
 ║ aggregates, domain services, ║
 ║ application services         ║
 ╚══════╤═══════════════════════╝
        │
 Driven ports (output interfaces)
        │ implemented by
 Driven adapters (DB repos, HTTP clients)
        │
 Driven side (databases, external APIs)
```

### Dependency rule

Dependencies point inward only.

| Layer | May depend on | Must not depend on |
|-------|---------------|--------------------|
| Domain | Nothing | Anything outside the domain |
| Application services | Domain, driven ports | Adapters, frameworks, DB drivers |
| Driving adapters | Application services, driving ports | Domain internals |
| Driven adapters | Driven port interfaces | The domain layer directly |

### Ports

- **Driving ports:** defined in the application layer, called by driving adapters. Example: `InvoiceCommandHandler`.
- **Driven ports:** defined in the domain or application layer, implemented by driven adapters. Example: `InvoiceRepository`.
- Ports speak the domain language. No SQL, no HTTP status codes.

### Checks at G2

- [ ] The domain imports nothing from adapters or frameworks.
- [ ] All DB access goes through a driven port.
- [ ] All HTTP in goes through a driving adapter.
- [ ] Domain tests use in-memory adapters.

---

## Screaming Architecture

The top-level structure shows the business, not the technology. A new developer should understand what the application does from the folder names.

**Bad:**
```
src/
  controllers/
  services/
  repositories/
  models/
```

**Good:**
```
src/
  invoicing/
    mandate/
    issuance/
    compliance/
  billing/
    subscription/
    payment/
  shared/
    currency/
    vatid/
```

### Rules

- Top-level packages are named after bounded contexts or core domain concepts.
- Technical layers (adapters, infrastructure) sit inside domain packages, not at the top.
- Entry points (`main.go`, `cmd/`, `app.ts`) stay at the edge.
- `shared/` or `common/` holds no domain logic.

### Checks at G2

- [ ] The top-level structure shows domain concepts.
- [ ] New packages use UL terms.
- [ ] No business logic in `utils/`, `helpers/` or `common/`.
- [ ] Infrastructure code is inside domain-named packages.
