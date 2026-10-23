# Observability

How to add traces, logs, and metrics, in any language or framework. How much
to add is the feature's stability level's call: Prototype adds none beyond what
the experiment needs, Alpha needs basic logging, and Beta adds metrics and
tracing. Whatever is added follows this.

## OpenTelemetry is the standard

- Traces, metrics, and logs use the OpenTelemetry (OTel) API and SDK for the
  language, exported over OTLP. No vendor SDKs or agents in application code;
  where telemetry goes is configuration, not code.
- Configure it with OTel's standard environment variables
  (`OTEL_SERVICE_NAME`, `OTEL_EXPORTER_OTLP_ENDPOINT`,
  `OTEL_RESOURCE_ATTRIBUTES`) and set `service.name`, `service.version`, and
  `deployment.environment.name` on the resource. With no endpoint set, export
  nothing or write to the console, so local setup needs nothing extra.
- Use OTel's instrumentation libraries for the framework, HTTP clients,
  database drivers, and queues instead of hand-writing spans and metrics for
  them. Hand-write only what they can't know: the domain.
- Name attributes by OTel's semantic conventions (`http.route`,
  `db.system.name`, `messaging.system`). Domain attributes get a domain
  namespace that can't collide with them, such as `auth.method` or
  `payment.provider`.
- Propagate context with W3C Trace Context: across HTTP calls, into queued
  messages and job payloads, and through the code (Go's `ctx`, for example).

## Traces

Every unit of work someone might need to investigate gets a span, with the kind
that matches its role:

- **SERVER**: each inbound request (HTTP, gRPC), usually from the framework's
  instrumentation. Named `{method} {route}` with the route template
  (`GET /appointments/{id}`), never the raw path.
- **CLIENT**: each outbound call, such as an HTTP API, the database, or a
  third-party SDK.
- **PRODUCER**: enqueuing a job or publishing a message. The trace context goes
  in the message's metadata.
- **CONSUMER**: processing that job or message, continuing the producer's
  trace, as a child, or as a link when one run handles many messages.
- **INTERNAL**: meaningful in-process work worth timing on its own, such as a
  service operation or a step of a long job. Not every function.

Long-running and scheduled jobs that no request started get a root span per
run, with child spans for each meaningful step or batch, not each item.

Every database query gets its own CLIENT span, whatever the database: SQL,
document, key-value, or search. The span's duration is the query's time, and
it carries:

- `db.query.text`: the query with every value replaced by a placeholder
  (`SELECT * FROM appointments WHERE id = $1`), never the values themselves.
- `db.system.name`, `db.operation.name`, and `db.collection.name` (the table
  or collection).
- A name of `{operation} {table}`, such as `SELECT appointments`.

Use the driver's OTel instrumentation when it records these. When it doesn't,
or there isn't one, wrap the driver once so every query gets a span, rather
than adding spans by hand at each call site.

- Span names stay low-cardinality. IDs and other specifics go in attributes.
- On failure, record the exception and set the span's status to Error. A 4xx
  is an error on a CLIENT span but not on a SERVER span, where it's the
  caller's mistake.
- Spans may carry high-cardinality attributes such as tenant, user, or record
  IDs. Never secrets, tokens, passwords, request bodies, or personal data such
  as emails.

## Logs

- Structured key-value logs, not formatted strings, through the language's
  standard logger bridged to OTel.
- Every log written while handling a request or job carries `trace_id` and
  `span_id`, so it can be joined to its trace. Log with the active context
  (`logger.InfoContext(ctx, ...)` in Go) so the bridge can add them. If the
  console output doesn't include them, add them in the log handler.
- Log at the top of the stack: HTTP handlers, job runners, message consumers,
  and the program's entry point. Lower layers, such as services, repositories,
  and clients, don't log. They return errors wrapped with context, and results
  that say what happened, and the top level decides what to log and at which
  level. When a lower layer handles something itself, such as a retry or a
  fallback, it records that on its span as an event or attribute.
- Levels:
  - **Error**: an operation failed and someone may need to act, such as an
    unexpected 500 or a job that gave up. Log it once, at the top, not at
    every layer it passes through.
  - **Warn**: something unexpected that was handled, such as a request that
    only succeeded through a fallback, or a limit being approached.
  - **Info**: significant events, low-volume enough to leave on in production:
    startup and shutdown, a job run starting and finishing, important domain
    state changes.
  - **Debug**: detail for diagnosing a problem. Off in production by default.
- An expected outcome for the caller, such as failed validation, not found, or
  401/403, is not an error. Log it at Debug, or not at all.
- The default level is Info, configurable without a code change.
- Never log secrets, tokens, passwords, or request bodies. Identify people and
  records by ID, not by email, name, or other personal data.

## Metrics

Instrumentation libraries already cover infrastructure metrics such as request
latency, database pools, and runtime. Hand-written metrics are about the
domain: sign-ups, sign-ins, payments, bookings, jobs processed. Record each
where its outcome is known, usually the service layer, and only once.

- Pick the instrument by what's measured: a **Counter** for things that happen
  (sign-ups), a **Histogram** for distributions (payment amounts, job
  durations), an **UpDownCounter** or **Gauge** for current values (active
  sessions, queue depth).
- Name with dot-separated lowercase namespaces, such as `auth.sign_ins` or
  `billing.payment.amount`. Put the unit in the instrument's unit (`s`, `By`,
  `{payment}`), not the name, and don't add `_total`. Record money in minor
  units, split by a `currency` attribute.
- Choose attributes by the questions the metric should answer. "How many
  sign-ins failed this week, and why?" means `auth.sign_ins` with `outcome`,
  `reason`, and `method`.
- Every attribute has a small, fixed set of values, tens rather than
  thousands, and their combinations stay small too. Good ones are `outcome`
  (`success`/`failure`), a failure reason as an enum (`invalid_password`,
  `locked`), method or provider, plan tier, currency, and region.
- Never use unbounded values as attributes: user, tenant, or organization IDs,
  emails, request or trace IDs, raw URLs or paths, error messages, free text,
  or amounts. Per-user and per-tenant questions are for traces and logs.

| Metric                   | Instrument | Unit           | Attributes                        |
| ------------------------ | ---------- | -------------- | --------------------------------- |
| `auth.sign_ups`          | Counter    | `{sign_up}`    | `method`, `plan`                  |
| `auth.sign_ins`          | Counter    | `{sign_in}`    | `method`, `outcome`, `reason`     |
| `billing.payments`       | Counter    | `{payment}`    | `provider`, `currency`, `outcome` |
| `billing.payment.amount` | Histogram  | `{minor_unit}` | `provider`, `currency`            |
| `jobs.duration`          | Histogram  | `s`            | `job` (its name), `outcome`       |
