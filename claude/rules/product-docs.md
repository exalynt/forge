# Product Docs

What a product is and does is documented in `docs/product/` at the root of its
repository:

```text
docs/product/
├── personas/         # who uses the product
├── capabilities/     # what people can do, with its stability level and known gaps
├── flows/            # how work gets done from start to finish
└── domain/
    ├── rules/        # what must or can't happen
    └── concepts/     # what something in the domain is
```

- Look here first for personas, capabilities, flows, business rules, and domain
  definitions.
- Whenever one is created or updated, it goes here, in the format for its type.
  Use the `docs` skill before writing one.
- Each capability's stability level and known gaps are in its capability doc.
