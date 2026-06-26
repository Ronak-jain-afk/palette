# Product

## Register

product

## Users

Designers and developers who need quick color exploration in the terminal. They already live in the command line — working on frontend projects, prototyping UIs, or building design systems. Their context is a focused work session where GUI apps feel interruptive: they want to generate, tweak, and export a palette in seconds without leaving their editor or terminal.

## Product Purpose

Palette generates professional, mood-driven color palettes directly in the terminal. It replaces the browser-tab detour to Coolors or Adobe Color with a CLI command or a lightweight TUI session. Success looks like: a designer or dev types `palette generate --mood dark --count 5`, sees five well-chosen swatches, locks two, iterates, exports as CSS variables — and never opens a browser for color work again.

## Brand Personality

Precise, calm, confident. The tool is intentional without being sterile, creative without being frivolous. It inspires trust: the colors are mathematically sound, the output is clean, the interface stays out of the way. Voice is direct and informative — no flourish, no fluff.

## Anti-references

- **Generic AI output.** Not bland, follow-the-scaffold designs. Every surface should feel considered, not assembled by a template.
- **Adobe-ware.** Not heavy, bloated, or enterprise-slow. Not Adobe Color. Palette is lightweight and immediate.
- **Overly gamified.** No flashy animations, no cutesy elements, no confetti on generation. A serious tool for serious creative work.

## Design Principles

1. **Opinionated defaults, expert overrides.** The tool chooses well for you, but never locks you out of fine control.
2. **Mathematical confidence.** Colors follow real color theory (harmony rules, perceptual deltas, WCAG contrast). The output isn't random — it's grounded.
3. **Terminal-native, not web-in-terminal.** Respect terminal conventions: clear keybindings, no mouse dependency, no pixel-perfect mimicry of a GUI. Work with the medium, not against it.
4. **Out of your way.** Every feature earns its place. No splash screens, no startup tips, no telemetry prompts. Start fast, quit faster.
5. **Accessible by default.** High-contrast compatible, screen-reader aware, reduced-motion friendly from day one.

## Accessibility & Inclusion

Target terminal-native accessibility:
- Respect system color schemes and high-contrast terminal themes
- TUI navigable by keyboard alone (no mouse requirement)
- Screen-reader compatible output structure
- Reduced-motion support for any animated transitions (status messages, palette refresh)
- No reliance on color alone for conveying state (lock indicators use symbols, not just hues)
