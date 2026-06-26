# Palette — Agent Guide

## State

Pre-implementation. Only docs committed; no Go code or build config exists yet.

## Source documents

| File | What it defines |
|------|----------------|
| `plan.md` | Product spec: data structures, CLI flags, TUI layout, key bindings, export formats |
| `IMPLEMENTATION_GUIDE.md` | Build order: which file to create first, exact function signatures, dependency map |
| `PRODUCT.md` | Strategy: register (product), audience, brand personality, design principles |
| `DESIGN.md` | Visual: dark/light OKLCH palette, typography, components, motion, accessibility |

## Architecture

```
cmd/           — cobra commands (root, generate, interactive, export, configure)
internal/
  color/       — palette.go (Color/Palette structs), harmony.go, mood.go
  display/     — renderer.go (ANSI escapes), swatches.go (color block rendering)
  tui/         — bubbletea model, update, view
  export/      — json.go, css.go, hex.go, yaml.go, clipboard.go
pkg/
  colorspace/  — conversions.go (RGB↔HSL↔Hex), colorful.go (DeltaE, contrast)
  terminal/    — capability detection
```

## Dependencies

```
cobra (CLI), bubbletea (TUI), lipgloss (TUI styling), go-colorful (color math),
termenv (ANSI), viper (config), clipboard (clipboard), go-runewidth (Unicode)
```

## Build order (do not skip)

1. Scaffold: `main.go` → `cmd/root.go`
2. Color: `pkg/colorspace/conversions.go` → test round-trips → `internal/color/*.go`
3. CLI: `internal/display/*.go` → `cmd/generate.go`
4. Export: `internal/export/*.go` → `cmd/export.go`
5. TUI: `internal/tui/*.go` → `cmd/interactive.go`
6. Config: `cmd/configure.go` → wire viper

## Design constraints

- Dark terminal-native default; light theme as alternative
- OKLCH only for all color values. Primary: `oklch(0.55 0.119 160.0)` (mossy green). Accent: `oklch(0.65 0.105 75)` (warm amber)
- `pm` no emojis in UI — use `[L]`/`[ ]` for lock state, plain text indicators only
- Restrained color strategy: TUI is near-black `oklch(0.08 0 0)`, brand color used sparingly
- Instant transitions. No animations in the TUI. Reduced-motion handled by design
- All keyboard-navigable. No mouse dependency
- `--no-color` flag must degrade gracefully to hex-only output
- Color values always paired with hex text — meaning never conveyed by hue alone

## Verification

No test suite exists yet. After any Go code is added, verify with `go build ./...` and manually check output with `go run . generate --mood dark --count 5`.
