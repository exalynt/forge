---
paths:
  - "**/*.{go,ts,tsx,js,jsx,mjs,cjs,py,rb,java,kt,rs,php,cs,swift,sql}"
  - "**/docs/engineering/**"
---

# Engineering Docs

How a repository's software is built is documented in `docs/engineering/` at
its root:

```text
docs/engineering/
├── README.md                     # for people: what these docs are for and what goes where
├── architecture/
│   ├── systems/<slug>.md         # how something overarching works
│   └── features/<slug>.md        # how one feature works, and what to know before changing it
├── providers/<type>/<vendor>.md  # a third-party service: setup and behavior
└── runbooks/<situation>.md       # what to do when something goes wrong
```

- Before changing code, search `docs/engineering/` for its package, file, or
  main symbol, and read what covers it. In a repository that links to another's
  docs, read those too.
- These docs complement the code and never replace it. What a competent reader
  can learn in a few minutes from the code, spec, migrations, or tests stays
  out.
- Engineering docs are the how, in technical words. Product docs are the what,
  in the domain's words. A gap is in one or the other, never both: if someone
  using the product would notice it, it's in the capability doc. If only
  someone working on the code would, it's in the engineering docs.
- A change that makes one of these docs wrong updates it in the same change.
  Use the `engineering-docs` skill before creating or updating one.
- A feature doesn't get a doc for being a feature. Create one only when there's
  something the code can't say.
