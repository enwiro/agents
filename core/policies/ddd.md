# Policy: Domain-Driven Design

The business domain is at the center. Code, tests and talks use one language, the **Ubiquitous Language** (UL), everywhere.

## Ubiquitous Language

- Every concept named by the PO appears exactly as written in code: classes, methods, variables, events.
- No translation to technical words. An "invoice mandate" is `InvoiceMandate`, not `Document` or `Request`.
- The UL is bounded. One word can mean different things in two bounded contexts. Do not merge their vocabularies.
- Keep a glossary in `.ai/local/ubiquitous-language.md`.

## Bounded contexts

- A bounded context is an explicit boundary where one model applies.
- Each context owns its data. Other contexts use a contract (anti-corruption layer or domain event), never a database join.
- Boundaries follow team ownership when possible.
- Unsure about a boundary: the Architect asks the human.

## Building blocks

**Aggregate**
- A cluster of objects changed as one unit. The root is the only entry point.
- It enforces its own invariants: a command either succeeds or fails with a domain error.
- Keep aggregates small. Refer to other aggregates by ID only.

**Entity**
- Has an identity that stays the same when its state changes.

**Value object**
- No identity, defined by its attributes, immutable.
- Use it for money, addresses, invoice numbers, VAT codes.

**Domain event**
- Immutable fact, named in the past tense with UL terms: `InvoiceIssued`, `MandateRevoked`.
- Other aggregates or contexts react through handlers, not direct calls.

**Domain service**
- Only for domain behavior that fits no single aggregate. No state.
- Not a place to dump logic that belongs in an aggregate.

**Application service**
- Runs a use case: coordinates domain objects and ports. No domain logic.
- It is the entry point for adapters.

## Forbidden

| Anti-pattern | Instead |
|--------------|---------|
| Anemic model (only getters and setters) | Behavior inside domain objects |
| Domain logic in controllers or services | Move it to aggregates or domain services |
| Database schema drives the model | The domain model drives the schema |
| One database shared by several contexts | Each context owns its schema |
| Technical names (`UserDTO`, `InvoiceDAO`) | Business names (`Customer`, `Invoice`) |
| Object reference between aggregates | Reference by ID |
