# forge logo

The mark is an anvil with a blue ingot on it, and the ingot carries a `>_` prompt: the code that forge
shapes. The tagline is **Exalynt's way of working, for every coding agent**.

Colors come from the Exalynt design system
([`exalynt/design`](https://github.com/exalynt/design), `src/exalynt/tokens.ts`):

| Use | Anvil | Ingot | Background |
| --- | --- | --- | --- |
| On light | midnight `#080f1a` | blue `#3b82f6` | cloud `#f7f9fc` or white |
| On dark | white `#ffffff` | bright blue `#60a5fa` | midnight `#080f1a` |

The `>_` is cut out of the ingot with a mask, so the mark works on any background.

## Which file to use

| Need | File |
| --- | --- |
| Mark on a light background | `svg/forge-mark-light.svg` |
| Mark on a dark background | `svg/forge-mark-dark.svg` |
| Mark that follows the viewer's light or dark setting | `svg/forge-mark-adaptive.svg` |
| One-color mark | `svg/forge-mark-mono-midnight.svg`, `svg/forge-mark-mono-white.svg` |
| Mark with the wordmark | `svg/forge-lockup-light.svg`, `svg/forge-lockup-dark.svg`, or the PNGs in `png/` |
| GitHub avatar or app icon | `png/forge-app-icon-dark-512.png` (or `-white`), `svg/forge-app-icon-*.svg` |
| GitHub social preview | `social/forge-social-preview.png` (1280×640) |

## Small sizes

The `>_` can't be read below 64px, so smaller sizes use drawings fitted to the pixel grid, with a
plain ingot. They're in `svg/small/`:

- **48px and 24px:** drawn on a 24px grid.
- **32px:** edges aligned to whole pixels.
- **16px:** drawn pixel by pixel.

Use them instead of scaling the large mark down. `forge-mark-*-plain.svg` is the large mark without
the prompt.

The mark is drawn on a 64-unit grid with its straight edges on even units. When you export a new
size, keep it at a whole or half pixel per unit so those edges stay sharp.

## Wordmark

"FORGE" is set in bold capitals with 0.1em letter spacing, like Exalynt's own wordmark. In the SVGs
the letters are outlines from Roboto Bold, so no font is needed. In live UI, set it as text in the
system sans stack, weight 700, `letter-spacing: 0.1em`. Size the mark at about 1.375× the font size,
with a gap of about 0.45em.

## Don't

- Recolor the ingot outside the Exalynt blues, stretch the mark, or add effects.
- Use the light mark on a dark background, or the reverse.
- Scale the large mark below 64px. Use the files in `svg/small/`.
