# Palette

A mood-driven color palette generator for the terminal. Generate harmonious color schemes from mood presets — export to CSS, JSON, YAML, or hex.

```bash
palette generate --mood nature --count 5
```

## Install

```bash
go install github.com/Ronak-jain-afk/palette@latest
```

Or download a prebuilt binary from the [releases page](https://github.com/Ronak-jain-afk/palette/releases).

## Usage

### Generate

```bash
palette generate                          # defaults: dark, 5 colors, analogous
palette generate --mood nature --count 3
palette generate --mood pastel --scheme complementary
palette generate --mood dark --count 4 --output json
```

| Flag | Default | Values |
|------|---------|--------|
| `--mood` | `dark` | dark, vintage, minimal, nature, pastel |
| `--count` | `5` | 2–10 |
| `--scheme` | `analogous` | analogous, complementary, triadic, tetradic, monochromatic |
| `--output` | (terminal display) | json, css, hex, yaml |

### Export

```bash
palette export --format css --output palette.css
palette export --format json --clipboard
palette export --mood nature --count 4 --format yaml
```

| Flag | Default | Values |
|------|---------|--------|
| `--format` | `json` | json, css, hex, yaml |
| `--mood` | `dark` | mood preset |
| `--count` | `5` | 2–10 |
| `--scheme` | `analogous` | harmony scheme |
| `--output` | (stdout) | file path |
| `--clipboard` | `false` | copy to clipboard |

### Configure defaults

```bash
palette configure defaults.mood nature
palette configure defaults.count 3
palette configure defaults.scheme triadic
```

Config is stored at `~/.config/palette/config.yaml`. Command-line flags override defaults.

### Interactive TUI

```bash
palette interactive
```

| Key | Action |
|-----|--------|
| `←` `→` | Focus color |
| `Space` | Lock/unlock |
| `r` | Regenerate focused |
| `R` | Regenerate all |
| `↑` `↓` | Adjust lightness |
| `[` `]` | Shift hue ±5° |
| `{` `}` | Shift hue ±30° |
| `<` `>` | Adjust saturation |
| `m` | Cycle mood |
| `s` | Cycle scheme |
| `h` `?` | Help overlay |
| `q` | Quit |

## Config

`~/.config/palette/config.yaml`:

```yaml
defaults:
  mood: dark
  count: 5
  scheme: analogous
```

## Build from source

```bash
git clone https://github.com/Ronak-jain-afk/palette.git
cd palette
go build -o palette .
```

## About

Palette generates colors using HSL math with mood-constrained saturation, lightness, and hue ranges. Each harmony scheme shifts hues according to color theory rules. Export helpers format the output for use in web projects, design tools, or further processing.

## License

MIT
