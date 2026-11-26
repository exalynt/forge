# Standards: configuration

For every project, in any language. Language and stack skills add to this.
Where this project's own code or instructions differ, follow the project.

## One config, read in one place

- Configuration comes from environment variables, read in one place into one
  config. Nothing else reads the environment.
- Group related settings (`HTTP`, `Database`, `Email`, `Payments`), with
  environment variable prefixes to match (`HTTP_*`, `DATABASE_*`, `EMAIL_*`,
  `PAYMENTS_*`).
- Each setting says what it controls and what happens when it's unset.
- Commands build the config once and pass each component the parts it needs.
  Components never read config globally.

## Defaults that make local setup simpler

- Provide sensible defaults as named constants, with a comment saying why.
  Typical ones are a localhost database URL for the local stack, `noop`
  providers instead of third-party ones, port `8080`, and `info` logging.
- The aim is for the project's documented local run command to work with no
  extra setup.
- Secrets and production-only values have no defaults.
- When a required setting is missing, the affected feature fails closed: its
  route responds `503`, or the command refuses to run. Never silently fall
  back to something insecure.
- Document every variable in the project's README.
