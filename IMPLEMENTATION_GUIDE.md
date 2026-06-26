# Implementation Guide: Moody CLI Color Palette Generator

A practical, step-by-step build guide in Go. Each phase builds on the last. Complete one phase fully before moving to the next.

---

## PHASE 0: Scaffolding

### Step 0.1 — Initialize the module

```bash
mkdir -p palette/{cmd,internal/{color,display,tui,export},pkg/{colorspace,terminal}}
cd palette
go mod init github.com/<you>/palette
```

### Step 0.2 — Install dependencies (install all now, never revisit)

```bash
go get github.com/spf13/cobra
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/lucasb-eyer/go-colorful
go get github.com/muesli/termenv
go get github.com/spf13/viper
go get github.com/atotto/clipboard
go get github.com/mattn/go-runewidth
```

### Step 0.3 — Wire up `main.go`

**File: `main.go`**

Import and execute `cmd.Execute()`. Keep it a single line of real logic.

```go
package main

import "github.com/<you>/palette/cmd"

func main() {
    cmd.Execute()
}
```

### Step 0.4 — Create the root command

**File: `cmd/root.go`**

- Define a root `cobra.Command` named `palette`
- Add global flags: `--no-color`, `--config`, `--verbose`
- Call `cmd.AddCommand()` for each subcommand (generate, interactive, export, configure)
- Stub subcommands inline with `Run: func(...) { ... }` that prints "not yet implemented"
- Bind viper to config file path (`~/.palette/config.yaml` or `$XDG_CONFIG_HOME/palette/config.yaml`)

**Key detail:** Use `cobra.OnInitialize(initConfig)` to load config before any command runs.

---

## PHASE 1: Color Foundations

### Step 1.1 — Color struct and constructors

**File: `internal/color/palette.go`**

Define the `Color`, `Palette`, and `Config` structs from the spec.

`Color` needs:
- `R, G, B uint8`
- `Hex string` (lazy — compute when accessed)
- Methods: `ToHSL()`, `HexString()`, `String()`

`Palette` needs:
- `Name string`, `Colors []Color`, `Mood string`, `GeneratedAt time.Time`, `Locked []bool`

**Strict rule:** Store RGB as the canonical representation. Compute HSL and Hex on demand. Never store redundant copies that can drift.

### Step 1.2 — RGB ↔ HSL ↔ Hex conversions

**File: `pkg/colorspace/conversions.go`**

Implement these functions:

```go
func RGBToHSL(r, g, b uint8) (h, s, l float64)
func HSLToRGB(h, s, l float64) (r, g, b uint8)
func HexToRGB(hex string) (uint8, uint8, uint8, error)
func RGBToHex(r, g, b uint8) string
```

**Algorithm sources:**
- RGB→HSL: Standard cylindrical conversion (0-360°, 0-1, 0-1)
- HSL→RGB: Chroma-based inverse
- `HexToRGB`: Strip `#`, parse two-char hex per component with `strconv.ParseUint`
- `RGBToHex`: `fmt.Sprintf("#%02X%02X%02X", r, g, b)`

**Test immediately:** Write a table test that round-trips 20 known colors (red, green, blue, black, white, gray, random hex values). If `HexToRGB(RGBToHex(x)) != x`, the conversion is wrong.

### Step 1.3 — Delegate to go-colorful for advanced color math

**File: `pkg/colorspace/colorful.go`**

Thin wrappers around `go-colorful` for operations that would be complex to hand-roll:

```go
func DeltaE(lab1, lab2 Color) float64        // CIE76, good enough
func Blend(c1, c2 Color, t float64) Color     // Linear interpolation in RGB
func ContrastRatio(c1, c2 Color) float64      // WCAG relative luminance
```

**Why a wrapper:** Isolates the dependency. If you swap color libraries later, you change one file.

### Step 1.4 — Quick self-check

```bash
go test ./pkg/colorspace/... -v
```

At this point every conversion must round-trip correctly.

---

## PHASE 2: CLI — `generate` Command

### Step 2.1 — color generation: random

**File: `internal/color/palette.go`**

Add `GenerateRandom() Color`:

```go
func GenerateRandom() Color {
    return Color{
        R: uint8(rand.Intn(256)),
        G: uint8(rand.Intn(256)),
        B: uint8(rand.Intn(256)),
    }
}
```

That's it. Random generation is trivial.

### Step 2.2 — Harmony algorithms

**File: `internal/color/harmony.go`**

Implement these functions. Each takes a base `Color` and returns `[]Color`:

```
Analogous(base, count)      → shift hue ±30°/count steps
Complementary(base)         → base + 180° hue shift
Complementary(base, count)  → split complement if count > 2
Triadic(base)               → base, base+120°, base+240°
Tetradic(base)              → base, base+90°, base+180°, base+270°
Monochromatic(base, count)  → same hue, vary S and L in steps
```

**Implementation pattern:** Convert to HSL, manipulate H/S/L values, convert back to RGB.

```go
func Analogous(base Color, count int) []Color {
    h, s, l := RGBToHSL(base.R, base.G, base.B)
    step := 30.0 / float64(count-1)
    colors := make([]Color, count)
    for i := range colors {
        offset := step * float64(i) - 15.0 // center on base
        rh, rl := h+offset, l
        // clamp h to [0, 360), l to [0, 1], s to [0, 1]
        colors[i] = HSLToRGB(normalizeHue(rh), clamp(s, 0, 1), clamp(rl, 0, 1))
    }
    return colors
}
```

### Step 2.3 — Mood presets

**File: `internal/color/mood.go`**

Define a `Mood` type and a preset registry:

```go
type Mood struct {
    Name        string
    Saturation  [2]float64  // min, max
    Lightness   [2]float64  // min, max
    HueRange    [][2]float64 // list of (start, end) hue ranges
}

var Presets = map[string]Mood{
    "dark": {
        Saturation: [2]float64{0.2, 0.45},
        Lightness:  [2]float64{0.15, 0.40},
        HueRange:   [][2]float64{{200, 280}, {300, 360}, {0, 20}, {170, 200}},
    },
    "vintage": {
        Saturation: [2]float64{0.3, 0.6},
        Lightness:  [2]float64{0.3, 0.6},
        HueRange:   [][2]float64{{10, 40}, {30, 50}, {40, 60}},
    },
    "minimal": {
        Saturation: [2]float64{0.05, 0.20},
        Lightness:  [2]float64{0.3, 0.8},
        HueRange:   [][2]float64{{0, 360}},
    },
    // nature, pastel — same pattern, values from spec
}
```

Also add `GenerateFromMood(mood string, count int, scheme string) Palette` that:
1. Picks a random hue from the mood's hue ranges
2. Generates base color with random S/L within mood constraints
3. Applies harmony algorithm
4. Returns a `Palette`

### Step 2.4 — Terminal renderer: ANSI escape sequences

**File: `internal/display/renderer.go`**

Define a `Renderer` struct that detects terminal capabilities:

```go
type Renderer struct {
    SupportsTrueColor bool
    Supports256       bool
    NoColor           bool
}
```

Detection logic:
- If `--no-color` → `NoColor = true`
- Check `os.Getenv("COLORTERM")` — if `truecolor` or `24bit`, support truecolor
- Check `os.Getenv("TERM")` — if not "xterm" or "xterm-256color", degrade

Rendering methods:

```go
func (r *Renderer) ColorBlock(c color.Color, width, height int) string
func (r *Renderer) ColorLine(colors []color.Color) string
func (r *Renderer) HexLabel(hex string) string
func (r *Renderer) Reset() string
```

For truecolor: `fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)`
For 256-color: map RGB to nearest xterm-256 index (pre-compute lookup table at startup)

### Step 2.5 — Color swatch display

**File: `internal/display/swatches.go`**

```go
func RenderSwatches(palette Palette, renderer *Renderer) string
```

Layout calculation:
- Terminal width → how many swatches fit side by side
- Each swatch: a colored block with the hex code underneath
- Minimum swatch size: 12 chars wide × 3 lines tall

### Step 2.6 — Wire up `generate` command

**File: `cmd/generate.go`**

```go
var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a color palette",
    RunE: func(cmd *cobra.Command, args []string) error {
        mood, _ := cmd.Flags().GetString("mood")
        count, _ := cmd.Flags().GetInt("count")
        scheme, _ := cmd.Flags().GetString("scheme")

        palette := color.GenerateFromMood(mood, count, scheme)

        renderer := display.NewRenderer(cmd.Flags())
        output := display.RenderSwatches(palette, renderer)
        fmt.Println(output)

        return nil
    },
}
```

Register flags in `init()`:
- `--mood` (default: "dark")
- `--count` (default: 5)
- `--scheme` (default: "analogous")
- `--base-color` (default: "random")

### Step 2.7 — Test it

```bash
go run . generate
go run . generate --mood vintage --count 3
go run . generate --mood nature --scheme triadic
```

You should see colored blocks in your terminal.

---

## PHASE 3: Export

### Step 3.1 — JSON export

**File: `internal/export/json.go`**

```go
func ToJSON(p Palette, includeMetadata bool) (string, error)
```

Use `encoding/json`. The `Palette` struct needs JSON tags (`json:"name"` etc.).

### Step 3.2 — CSS export

**File: `internal/export/css.go`**

```go
func ToCSS(p Palette) string
```

Generate `:root { --color-1: #...; --color-2: #...; }`.

### Step 3.3 — Plain hex export

**File: `internal/export/hex.go`**

```go
func ToHexLines(p Palette) string
```

One hex per line, no decoration.

### Step 3.4 — YAML export

**File: `internal/export/yaml.go`**

Use `gopkg.in/yaml.v3` (add to go.mod if needed).

### Step 3.5 — Export router

**File: `internal/export/export.go`**

```go
func Export(p Palette, format string, includeMetadata bool) (string, error) {
    switch format {
    case "json": return ToJSON(p, includeMetadata)
    case "css":  return ToCSS(p)
    case "hex":  return ToHexLines(p)
    case "yaml": return ToYAML(p, includeMetadata)
    default:     return "", fmt.Errorf("unknown format: %s", format)
    }
}
```

### Step 3.6 — Clipboard support

**File: `internal/export/clipboard.go`**

```go
func CopyToClipboard(text string) error
```

Wrap `github.com/atotto/clipboard`. If it returns an error (no xclip, no pbcopy), return a user-friendly message telling them to pipe the output instead.

### Step 3.7 — Wire up `export` command

**File: `cmd/export.go`**

Reads the current palette (from a temp file or in-memory during a session) and writes it to stdout or clipboard.

Flags: `--format`, `--clipboard`, `--output-file`.

### Step 3.8 — Wire `generate --output`

Add `--output` flag to the generate command. After generation, if `--output=json`, call the export function instead of the renderer.

---

## PHASE 4: Interactive TUI

### Step 4.1 — Bubble Tea model

**File: `internal/tui/model.go`**

```go
type model struct {
    palette     color.Palette
    focusIndex  int
    mode        string  // "view", "adjust", "export"
    renderer    *display.Renderer
    // ... keybindings, viewport, etc.
}
```

Implement the `tea.Model` interface:
- `Init() tea.Cmd` — returns `tea.EnterAltScreen`
- `Update(msg tea.Msg) (tea.Model, tea.Cmd)` — handle key presses
- `View() string` — render the UI

### Step 4.2 — Update loop

**File: `internal/tui/update.go`**

Handle these key events at minimum:

| Key | Action |
|-----|--------|
| `q` / `ctrl+c` | `tea.Quit` |
| `left` / `right` | Move focus (wrap around) |
| `space` | Toggle lock on focused color |
| `r` | Regenerate focused color |
| `R` | Regenerate all (respect locks) |
| `up` / `down` | Fine-tune lightness ±5% |
| `[` / `]` | Shift hue ±5° |
| `{` / `}` | Shift hue ±30° |
| `<` / `>` | Adjust saturation ±5% |
| `m` | Cycle mood |
| `s` | Cycle scheme |
| `e` | Switch to export mode |
| `h` / `?` | Show help overlay |

### Step 4.3 — View rendering

**File: `internal/tui/view.go`**

Build the UI string:

1. **Header bar**: "Palette Generator" + current mood/scheme
2. **Color swatches row**: Colored blocks with hex codes. Focused swatch gets a highlighted border. Locked swatches show a lock indicator.
3. **Info panel**: Color count, lock count, harmony, mood, contrast
4. **Footer**: Available keybindings
5. **Status line**: Current color's RGB/HSL details

**Layout approach:** Use `lipgloss` for composing styled strings with proper alignment and borders.

```go
func (m model) View() string {
    header := m.renderHeader()
    swatches := m.renderSwatches()
    info := m.renderInfo()
    footer := m.renderFooter()
    status := m.renderStatus()

    return lipgloss.JoinVertical(lipgloss.Top,
        header,
        swatches,
        info,
        footer,
        status,
    )
}
```

### Step 4.4 — Wire up `interactive` command

**File: `cmd/interactive.go`**

```go
var interactiveCmd = &cobra.Command{
    Use:   "interactive",
    Short: "Launch interactive palette editor",
    RunE: func(cmd *cobra.Command, args []string) error {
        p := tea.NewProgram(tui.InitialModel())
        if _, err := p.Run(); err != nil {
            return err
        }
        return nil
    },
}
```

Initial model generates a starting palette from the `--initial-mood` and `--initial-color` flags.

---

## PHASE 5: Configuration & Polish

### Step 5.1 — Viper-based config

**File: `cmd/root.go`** (extend `initConfig`)

```go
func initConfig() {
    cfgFile, _ := rootCmd.PersistentFlags().GetString("config")
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        configDir, _ := os.UserConfigDir()
        viper.AddConfigPath(filepath.Join(configDir, "palette"))
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
    }
    viper.SetEnvPrefix("PALETTE")
    viper.AutomaticEnv()
    viper.ReadInConfig() // silent fail — defaults if no config
}
```

### Step 5.2 — `configure` command

**File: `cmd/configure.go`**

```go
palette configure <key> <value>
```

Parses key-value pairs and writes to the config YAML file. Create the config directory if it doesn't exist.

Supported keys: `defaults.mood`, `defaults.count`, `defaults.scheme`, `preferences.color_format`, etc.

### Step 5.3 — Error recovery for TUI

Wrap `tea.Program.Run()` in a recovery block. If the TUI panics, print a friendly message and exit with code 1 instead of dumping a stack trace.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Fprintf(os.Stderr, "Palette encountered an error and needs to close.\n")
        os.Exit(1)
    }
}()
```

### Step 5.4 — Terminal fallback

In `Renderer`, when detected terminal doesn't support colors:
- No ANSI escape sequences at all
- Output hex codes as plain text lines
- Block characters become `████` using literal chars

---

## PHASE 6: Testing & Release

### Step 6.1 — What to test

| Package | Tests | Priority |
|---------|-------|----------|
| `pkg/colorspace` | Round-trip conversions, edge cases (black, white, pure hues) | **Critical** |
| `internal/color` | Harmony produces correct count, mood constraints respected | **Critical** |
| `internal/export` | Each format produces valid output | **High** |
| `cmd` | Flag parsing, config loading, help text | **Medium** |

### Step 6.2 — Table-driven tests pattern

```go
// pkg/colorspace/conversions_test.go
func TestRGBToHex(t *testing.T) {
    tests := []struct {
        r, g, b uint8
        want    string
    }{
        {0, 0, 0, "#000000"},
        {255, 255, 255, "#FFFFFF"},
        {255, 0, 0, "#FF0000"},
        {26, 26, 46, "#1A1A2E"},
    }
    for _, tt := range tests {
        got := RGBToHex(tt.r, tt.g, tt.b)
        if got != tt.want {
            t.Errorf("RGBToHex(%d,%d,%d) = %s; want %s", tt.r, tt.g, tt.b, got, tt.want)
        }
    }
}
```

### Step 6.3 — Build for distribution

```bash
go build -ldflags="-s -w" -o palette .
```

Cross-compile:

```bash
GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o palette-linux-amd64 .
GOOS=darwin  GOARCH=arm64 go build -ldflags="-s -w" -o palette-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o palette-windows-amd64.exe .
```

---

## APPENDIX A: File Creation Order

Build files in this exact sequence to minimize compilation errors:

1. `main.go`
2. `cmd/root.go`
3. `pkg/colorspace/conversions.go` + `conversions_test.go`
4. `internal/color/palette.go`
5. `internal/color/harmony.go`
6. `internal/color/mood.go`
7. `internal/display/renderer.go`
8. `internal/display/swatches.go`
9. `cmd/generate.go`
10. `internal/export/*.go`
11. `cmd/export.go`
12. `internal/tui/model.go`
13. `internal/tui/update.go`
14. `internal/tui/view.go`
15. `cmd/interactive.go`
16. `cmd/configure.go`

## APPENDIX B: First Command to Make It Work

After Phase 2, before anything else is built:

```bash
go run . generate --mood dark --count 5
# You should see 5 colored blocks in your terminal
```

If that works, the foundation is solid. If not, fix the conversion functions first — everything depends on them.

## APPENDIX C: Dependency Map

```
main.go → cmd/*.go
cmd/generate.go → internal/color/*, internal/display/*
cmd/export.go → internal/export/*
cmd/interactive.go → internal/tui/*
internal/color/palette.go → pkg/colorspace/*
internal/color/harmony.go → pkg/colorspace/*
internal/color/mood.go → pkg/colorspace/*
internal/tui/* → internal/color/*, internal/display/*, internal/export/*
internal/display/* → internal/color/*, pkg/colorspace/*
```
