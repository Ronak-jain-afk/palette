# Design

## Theme

Dark theme (terminal-native default); light theme as optional alternative.

**Mood phrase:** Terminal session at dusk — quiet precision, muted confidence, the calm of a well-typed command.

## Color Strategy

**Restrained.** Tinted neutrals plus one accent used sparingly (<10% of surface). The brand color (mossy green) primarily appears in headers, selection highlights, borders, and the logo/wordmark. The TUI itself is near-black — the interface recedes and the generated colors take center stage.

## Color Palette

### Dark Mode

| Role | OKLCH Value | Description |
|------|-------------|-------------|
| **bg** | `oklch(0.08 0 0)` | Near-black. Terminal-native background. No hue tint. |
| **surface** | `oklch(0.14 0.005 160)` | Card/panel background. Faint green warmth, barely perceptible. |
| **ink** | `oklch(0.88 0.005 160)` | Body text. Faint green tint reads as cohesive without announcing itself. |
| **muted** | `oklch(0.55 0.006 160)` | Secondary text, labels. Clear contrast against bg (~5:1). |
| **primary** | `oklch(0.55 0.119 160.0)` | Mossy green. Headers, selection highlights, active borders, brand mark. White text on fills. |
| **accent** | `oklch(0.65 0.105 75)` | Warm amber-copper. Status indicators, badges, links. White text on fills. |

### Light Mode

| Role | OKLCH Value | Description |
|------|-------------|-------------|
| **bg** | `oklch(0.995 0 0)` | Pure white. |
| **surface** | `oklch(0.95 0.005 160)` | Card background. Traces of brand hue. |
| **ink** | `oklch(0.12 0.005 160)` | Body text. Near-black with brand-ward warmth. |
| **muted** | `oklch(0.55 0.006 160)` | Secondary text. |
| **primary** | `oklch(0.55 0.119 160.0)` | Mossy green. Same as dark. |
| **accent** | `oklch(0.65 0.105 75)` | Warm amber-copper. Same as dark. |

### Checks

- **ink vs bg (dark):** ~15:1. **ink vs bg (light):** ~16:1. Both exceed WCAG AAA.
- **muted vs bg (dark):** ~5:1. **muted vs bg (light):** ~5.5:1. Both exceed WCAG AA for body text (4.5:1).
- **primary chroma:** 0.119 ≤ 0.23 ✓. **primary L:** 0.55 — white text on filled backgrounds ✓.
- **accent chroma:** 0.105 ≥ 0.10 ✓. Not in the muddy mid-tone danger zone.
- **primary vs accent:** Distinct hue (160° → 75°) and lightness (0.55 → 0.65). No confusion risk.

## Typography

### TUI (terminal interface)

- **Font:** System monospace (`"SF Mono", "JetBrains Mono", "Fira Code", "Cascadia Code", "Consolas", monospace`)
- **Scale:** Terminal's own font size; no explicit scale needed — the TUI layout respects the user's terminal preferences.
- **Weights:** Regular (400) for body, Bold (700) for headers and active selection.

### Marketing / documentation (website, README, help text)

- **Display:** System sans-serif (`-apple-system, "Segoe UI", "Inter", "Roboto", sans-serif`)
- **Mono:** Same as TUI stack for code blocks and CLI examples.
- **Body line length:** 65-75ch max.

## Components

### TUI Components

| Component | Description |
|-----------|-------------|
| **Color swatch** | Filled block with hex label. Active swatch gets a `primary` border. Locked swatch shows a lock symbol. |
| **Header bar** | Mood + scheme indicators, separated by `·`. Uses `muted` for secondary info. |
| **Info panel** | Footer area with stats (color count, lock count, contrast info). `muted` labels, `ink` values. |
| **Key binding help** | Overlay listing shortcuts. Grouped by category. `surface` background, `primary` for modifier keys. |
| **Status line** | Bottom line showing focused color's RGB/HSL values. `muted` text. |

### Layout

- **Swatch row:** Flexbox, horizontal. Each swatch is a colored block with hex underneath. Min swatch width: 12 chars. Focused swatch expands slightly (highlight border).
- **Panel dividers:** Single `muted`-colored horizontal rules (`─` characters). No heavy borders.
- **Spacing:** Single blank line between sections. No padding within the terminal beyond what the character grid provides.

## Motion & Animation

- **Palette regeneration:** Instant (no fade). A CLI tool responds immediately.
- **Lock toggle:** Instant symbol change. No animation needed.
- **Mode transitions** (view → export): Instant swap. No slide or crossfade within the TUI.
- **Reduced motion:** All transitions are instant by design — no `prefers-reduced-motion` override needed.

## Terminal Capabilities

| Feature | Detection | Fallback |
|---------|-----------|----------|
| Truecolor (24-bit) | `$COLORTERM` = `truecolor` or `24bit` | Map to nearest xterm-256 index |
| 256-color | `$TERM` = `xterm-256color` or similar | 16-color ANSI |
| No color | `--no-color` flag or unsupported terminal | Hex text only |
| Unicode | Check terminal locale | ASCII fallback for symbols |

## Accessibility

- All keyboard-navigable. No mouse dependency.
- Status indicators use symbols + text, never color alone.
- Lock states show `[L]` / `[ ]`. Same symbols in all terminal modes.
- Color values always shown as hex text — the meaning is never in the hue alone.
- Respects system-level reduced motion and high-contrast settings.
