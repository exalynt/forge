# GA

## Purpose

People are allowed to depend on this feature. GA is a stronger commitment to
reliability, compatibility, maintainability, and long-term operation.

> We are intentionally allowing people to depend on this and will engineer and
> operate it accordingly.

## Expected (where applicable)

Reliability:

- behavior is dependable
- important workflows are comprehensively tested
- important failure modes have known behavior
- performance characteristics are understood
- realistic scale has been validated

Operations:

- operational ownership is clear
- monitoring and alerting are appropriate
- production incidents can be investigated effectively
- backups and recovery exist where relevant
- migrations and upgrades are safe

Contracts and lifecycle:

- compatibility expectations are understood
- public contracts are intentionally managed
- lifecycle and deprecation expectations are considered
- ongoing maintenance expectations are understood

Quality and documentation:

- accessibility expectations are satisfied
- security posture is appropriate for production dependence
- documentation is sufficient for users, operators, and developers
- significant architectural tradeoffs are documented

## Changing it

Changes must be non-breaking. If a request seems to need a breaking change,
clarify the direction with the user before making it. A feature with no
declared level is GA.

## What GA is not

GA does not mean perfect, defect-free, or that every conceivable edge case is
handled. It does not license speculative extensibility or scale beyond what
real dependence requires. The ceiling still applies: invest in what the
commitment demands, not in everything imaginable.
