# Standards: layers and boundaries

For every project, in any language. Language and stack skills add to this.
Where this project's own code or instructions differ, follow the project.

## Layers

Entry points → services → boundary interfaces (repositories, external
providers).

- **Services own the business logic and rules:** invariants, authorization and
  scoping, orchestration across repositories and providers, and transactions.
  Name operations in business terms where that fits (`Reschedule`, `Cancel`,
  `Approve`), not only CRUD verbs. One service per area of the domain.
- **Entry points are thin.** An HTTP handler, a command, or a job decodes and
  validates its input, calls a service, maps its error, and renders the
  result. No business rules.
- **In a UI, pages and components are the entry points.** They call the app's
  services layer for anything external, and never call an API or SDK
  directly.
- **Repositories only persist.** See [repositories.md](repositories.md).

## Domain models and translation boundaries

Core logic operates on the project's own domain models. Anything crossing to
or from an external system is translated at that boundary, whether it's the
database, a vendor API, a webhook payload, or outbound HTTP. The boundary's
interface takes and returns domain models, and its implementation does the
translation. Vendor types never leak past their implementation.

## Interfaces and implementations

- A boundary is named for the concept, and defines its interface in domain
  terms: `repositories`, `email`, `payments`, `storage`.
- Its implementations are named for the technology or vendor: `postgres`,
  `smtp`, `noop`, `stripe`. Only an implementation imports the vendor's SDK or
  driver.
- Services depend on the interface. The concrete implementation is chosen
  once, where the app is wired, usually from config.

Add an interface for a real boundary. A `noop` implementation is welcome when
it lets the app run locally without a vendor's credentials. Don't add
interfaces speculatively.
