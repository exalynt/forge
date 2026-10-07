# Engineering Philosophy

The best implementation is the simplest one that appropriately meets the
feature's needs and stability level today. The goal is appropriate complexity,
not minimal code. Sometimes the right answer is more structure or more rigor,
when a present, concrete need justifies it.

## Design for what we currently know

- Build against concrete, current requirements. A requirement nobody has asked
  for, and no evidence supports, is not a requirement.
- "We will probably need it eventually" is not a reason to build it now. When
  eventually arrives, we will know more and build it better.
- When uncertain, prefer the decision that is cheapest to reverse.
- An explicitly stated known gap is better than speculative complexity that
  handles imagined cases.

## Complexity needs a reason

Complexity is justified by things that exist now: a demonstrated failure, a
real second use case, a measured performance problem, an explicit requirement,
a security or privacy need, or the feature's current stability level. It is
not justified by what might happen later.

## Abstraction

- Introduce an abstraction when it solves a problem that exists now: real
  duplication, a second real implementation, or a boundary that is genuinely
  hard to reason about without it.
- Do not add interfaces, extension points, plugin mechanisms, or configuration
  because they could be useful someday.
- Prefer clear, direct code to clever, general code. A few similar concrete
  blocks are often better than one premature abstraction.
- Code that is easy to delete or replace is often better than code generalized
  to survive requirements we cannot see yet.

## Replacing early work

- Early implementations exist to get us to the next decision. Replace or delete
  them once they have served that purpose; don't preserve them out of sunk cost.
- Don't contort new work to stay compatible with code that was always meant to
  be temporary.

## Performance and operations

- Optimize when there is evidence of a problem at realistic load, or when the
  current stability level calls for validated performance, not on intuition
  about future scale.
- Don't add operational infrastructure (dashboards, alerting, elaborate
  retries, caches, queues, flag systems, scaling machinery) ahead of a
  demonstrated need or the stability level that requires it.

## Knowing what not to build

Good engineering includes deciding what not to build. When you deliberately
leave something out, say so briefly. A plainly stated known gap is part of a
complete answer. When existing code already exceeds what the feature needs,
point that out instead of extending the pattern.

## Not negotiable at any level

- Security and privacy are never speculative. Authentication, authorization,
  validation at trust boundaries, secret handling, and protection of personal
  data meet an appropriate baseline from the first line of code.
- Simple does not mean sloppy. What we do build should be correct and clear.
