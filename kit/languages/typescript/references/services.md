# TypeScript: services

For TypeScript projects, adding to the `standards` skill's
[layers.md](../../standards/references/layers.md). Where this project's own
code or instructions differ, follow the project. Put new code in this layout.
Don't move existing code into it unless asked.

A service is an external API or SDK the app talks to, over HTTP REST or
otherwise: the Acme API, Stripe, and so on. Each one gets a directory under
`src/services/`, named for it in camelCase: `src/services/acme/`,
`src/services/stripe/`.

## One module per resource

- Each resource the app uses gets its own module in its service's directory,
  named in camelCase for it: `src/services/acme/appointments.ts`,
  `src/services/acme/paymentMethods.ts`.
- Each module exports one function per operation, named for the operation and
  the resource: `listAppointments`, `getAppointment`, `createAppointment`,
  `updateAppointment` (`PATCH`), `replaceAppointment` (`PUT`),
  `deleteAppointment`, and actions by their verb, such as
  `rescheduleAppointment`. Add a function when something needs it, not before.
- Each module declares the TypeScript types for its resource and requests,
  mirroring the service's shapes.
- Each function takes an optional `AbortSignal` and passes it on to the
  request, so callers can cancel it.
- A list function takes the search, filters, sort, and page in the app's own
  terms, and turns them into whatever the service takes.

## Shared plumbing

Within a service, the base URL, credentials, request and response handling,
and turning the service's error responses into thrown errors live in one
module, `src/services/<service>/client.ts`. For an SDK, that's where it's set
up. Resource modules build on it. Each service has its own client. Don't
share one across services.
