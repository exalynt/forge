---
paths:
  - "src/services/**"
  - "src/pages/**/*.tsx"
  - "src/components/**/*.tsx"
---

# React projects: services

General conventions for React projects. Where this project's CLAUDE.md
differs, follow it. Put new code in this layout. Don't move existing code into
it unless asked.

A service is an external API or SDK the app talks to, over HTTP REST or
otherwise: the Exalynt API, Stripe, and so on. Each one gets a directory under
`src/services/`, named for it in camelCase: `src/services/exalynt/`,
`src/services/stripe/`.

## One module per resource

- Each resource the app uses gets its own module in its service's directory,
  named in camelCase for it: `src/services/exalynt/appointments.ts`,
  `src/services/exalynt/paymentMethods.ts`.
- Each module exports one function per operation, named for the operation and
  the resource: `listAppointments`, `getAppointment`, `createAppointment`,
  `updateAppointment` (`PATCH`), `replaceAppointment` (`PUT`),
  `deleteAppointment`, and actions by their verb, such as
  `rescheduleAppointment`. Add a function when a page needs it, not before.
- Each module declares the TypeScript types for its resource and requests,
  mirroring the service's shapes.

## Shared plumbing

Within a service, the base URL, credentials, request and response handling,
and turning the service's error responses into thrown errors live in one
module, `src/services/<service>/client.ts`. For an SDK, that's where it's set
up. Resource modules build on it. Each service has its own client. Don't
share one across services.

## Pages and components

Pages and components call these functions. They never call `fetch`, build a
service's URLs, read status codes, or call an SDK directly. The exception is an
SDK's own React components and hooks, such as Clerk's, which pages use as
documented.
