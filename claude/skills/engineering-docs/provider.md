# Provider

`docs/engineering/providers/<type>/<vendor>.md`, one per third-party service in
each role it plays, such as `payments/stripe.md`.

A provider doc covers what's set up outside the repository and how the service
behaves in ways the code can't show. How the provider boundary works as a whole
is a system doc.

## Layout

```markdown
# <Vendor>

<What this service does for the product, and what stays on our side.>

**Code:** <where to start reading>

## What it does here

<How it's used, and behavior checked by hand, with the date it was checked.>

## Setup

<What's done by hand in the vendor's dashboard: settings, webhooks, and the
permissions a key needs.>

## Local development

## Production

## Tradeoffs

## Gotchas

## Known gaps

## Checklist for a new <account | environment>

1. <Each step, in order, ending with how to check it worked.>
```

## Sections

- **Code** and **What it does here**: required.
- **Setup**: required when anything is set up by hand. Say where a value comes
  from and what it needs. The list of environment variables goes wherever the
  stack's config rule puts it.
- **Local development** and **Production**: only what differs from the defaults.
- **Tradeoffs** and **Gotchas**: only if there are some.
- **Known gaps**: required. Write "None." when there are none.
- **Checklist**: required when Setup is. It's what someone follows, and it links
  to Setup rather than repeating it.

Never put a real key, secret, or account in the doc, only where to find one.

The format of Tradeoffs, Gotchas, and Known gaps is in
[SKILL.md](SKILL.md#tradeoffs-gotchas-and-known-gaps).
