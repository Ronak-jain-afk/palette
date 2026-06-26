# Detailed Development Plan: Moody CLI Color Palette Generator in Go

This is a comprehensive, code-free roadmap for building your CLI tool. It's structured like a product specification document to guide your development process from start to finish.

---

## 1. PROJECT OVERVIEW

### Core Concept
A terminal-based tool that generates professional, mood-driven color palettes. Users can generate palettes, lock colors, and interactively adjust them—all within the terminal using a Text User Interface (TUI).

### Target User Experience
- **Quick generation**: `palette generate --mood dark --count 5`
- **Interactive mode**: `palette interactive` opens a TUI similar to Coolors
- **Export**: `palette export --format css` outputs the current palette

---

## 2. PROJECT ARCHITECTURE

### Directory Structure
```
palette/
├── cmd/
│   ├── root.go           # Main command entry point
│   ├── generate.go       # Generate subcommand
│   ├── interactive.go    # TUI subcommand
│   └── export.go         # Export subcommand
├── internal/
│   ├── color/            # Color manipulation logic
│   │   ├── palette.go    # Palette generation
│   │   ├── harmony.go    # Color theory algorithms
│   │   └── mood.go       # Mood presets
│   ├── display/          # Terminal rendering
│   │   ├── renderer.go   # ANSI escape sequences
│   │   └── swatches.go   # Color block display
│   ├── tui/              # Interactive interface
│   │   ├── model.go      # State management
│   │   ├── update.go     # Event handling
│   │   └── view.go       # UI rendering
│   └── export/           # Export functionality
│       ├── css.go
│       ├── json.go
│       └── clipboard.go
├── pkg/                  # Reusable utilities
│   ├── colorspace/       # Color space conversions
│   └── terminal/         # Terminal capability detection
├── go.mod
└── main.go
```

---

## 3. DATA STRUCTURES

### Color Representation
```
struct Color {
    R, G, B     uint8    // RGB values 0-255
    Hex          string   // #RRGGBB format
    HSL          struct {
        H, S, L  float64  // 0-360, 0-1, 0-1
    }
}
```

### Palette Structure
```
struct Palette {
    Name        string
    Colors      []Color
    Mood        string
    GeneratedAt time.Time
    Locked      []bool    // Tracks which colors are locked
}
```

### Configuration
```
struct Config {
    DefaultMood     string
    DefaultCount    int
    OutputFormat    string
    TerminalColor   bool   // Truecolor vs 256-color
}
```

---

## 4. CORE FUNCTIONALITY: COLOR GENERATION

### Color Space Utilities
- Convert between RGB, HSL, HEX, and CIE-LAB
- Implement perceptual color differences (Delta E)
- Ensure color consistency across conversions

### Generation Algorithms

**Base Generator Functions:**
1. **Random**: Generate completely random colors with constraints
2. **Mood Presets**: Adjust generation parameters based on mood
3. **Harmony Rules**: Apply color theory transformations

**Harmony Algorithms:**
- **Analogous**: Take base color ± 30° in HSL wheel, generate 3-5 colors
- **Complementary**: Base + 180° hue shift, with possible split complements
- **Triadic**: Base + 120°, Base + 240°
- **Tetradic**: Two complementary pairs (rectangle on color wheel)
- **Monochromatic**: Vary lightness and saturation at fixed hue

**Mood Presets:**
```
Dark Moody:
  - Saturation: 20-45%
  - Lightness: 15-40%
  - Hues: Blue, Purple, Teal, Deep Red

Warm Vintage:
  - Saturation: 30-60%
  - Lightness: 30-60%
  - Hues: Orange, Brown, Ochre, Mustard

Minimal Corporate:
  - Saturation: 5-20%
  - Lightness: 30-80%
  - Hues: Grays with one accent color

Nature Inspired:
  - Saturation: 40-70%
  - Lightness: 25-60%
  - Hues: Green, Brown, Sky Blue, Earth Tones

Pastel Soft:
  - Saturation: 20-35%
  - Lightness: 65-85%
  - Hues: Any, but with low saturation
```

### Generation Process
1. Select base color (random or user-provided)
2. Apply mood constraints (saturation/lightness ranges)
3. Apply harmony rules
4. Validate colors meet accessibility standards (optional)
5. Generate final palette

---

## 5. COMMAND-LINE INTERFACE (CLI)

### Root Command
```
palette [command] [flags]
```

### Subcommands

**1. Generate**
```
palette generate [flags]
Flags:
  --mood string       dark, vintage, minimal, nature, pastel (default: "dark")
  --count int         Number of colors (default: 5)
  --scheme string     analogous, complementary, triadic, monochromatic (default: "analogous")
  --base-color string #RRGGBB or "random" (default: "random")
  --accessibility     Check WCAG contrast ratios
  --output string     Display or export immediately
```

**2. Interactive**
```
palette interactive [flags]
Flags:
  --initial-mood string    Starting mood preset
  --initial-count int      Starting color count
```

**3. Export**
```
palette export [flags]
Flags:
  --format string      json, css, yaml, hex (default: "json")
  --output-file string File to write to (default: stdout)
  --clipboard          Copy to clipboard
```

**4. Configure**
```
palette configure [key] [value]
Sets default preferences in ~/.palette/config.yaml
```

### Global Flags
```
--no-color          Disable color output
--config string     Custom config file path
--verbose           Enable detailed logging
```

---

## 6. INTERACTIVE TUI DESIGN

### Bubble Tea Model

**State Management:**
```
Model state contains:
  - Palette (current colors)
  - Focus index (which color is selected)
  - Locked colors (bool array)
  - Mode (view, adjust, export)
  - Key bindings
  - Viewport size
  - Error messages
```

**Key Bindings:**
```
General:
  q / Ctrl+C     Quit
  ?              Show help overlay
  h              Show keyboard shortcuts

Navigation:
  ← / →          Move focus between colors
  Tab            Jump to next unblocked color
  /              Search/filter (future)

Color Actions:
  Space          Lock/unlock current color
  r              Regenerate current color
  R              Regenerate entire palette (preserve locks)
  up/down        Fine-tune lightness
  [ / ]          Shift hue ±5°
  { / }          Shift hue ±30° (full harmony step)
  < / >          Adjust saturation ±5%
  c              Copy current color HEX to clipboard
  a              Add new color
  d              Delete current color (if more than 2)

Mood & Scheme:
  m              Cycle through moods
  s              Cycle through harmony schemes
  p              Show palette info (mood, scheme, generation date)

Export & Sharing:
  e              Export palette (opens export view)
  Ctrl+C         Copy palette to clipboard (JSON format)

History:
  Ctrl+Z         Undo last change
  Ctrl+Y         Redo
```

**UI Layout:**
```
╭────────────────────────────────────────────────────────────────────────────╮
│ Palette Generator                                       ● Dark • Analog   │  
├────────────────────────────────────────────────────────────────────────────┤
│                                                                            │
│  ████████  ████████  ████████  ████████  ████████                          │
│  #1A1A2E  #16213E  #0F3460  #533483  #E94560                               │
│     [L]         •          •          •           •                        │
│                                                                            │
├────────────────────────────────────────────────────────────────────────────┤
│ Palette                                                                    │
│   Colors     5        Locked      1        Harmony     Analog              │
│   Mood       Dark     Contrast    AAA      Terminal    TrueColor           │
├────────────────────────────────────────────────────────────────────────────┤
│ r Regenerate    Space Lock    m Mood    s Scheme    e Export    ? Help     │
╰────────────────────────────────────────────────────────────────────────────╯

Selected → #1A1A2E  RGB(26,26,46)  HSL(240°,27%,14%)
```

**Color Focus Rendering:**
- Active color gets a bright border or reverse video
- Locked colors show a lock indicator (`[L]`)
- Show RGB/HSL values for focused color
- Display contrast ratios with adjacent colors (WCAG)

---

## 7. TERMINAL RENDERING ENGINE

### Capability Detection
- Check `$COLORTERM` environment variable
- Detect truecolor support (24-bit)
- Fallback to 256-color support
- Fallback to 16-color support
- Disable colors with `--no-color`

### ANSI Escape Code Generation

**Truecolor (24-bit):**
- Foreground: `\033[38;2;R;G;Bm`
- Background: `\033[48;2;R;G;Bm`
- Bold text: `\033[1m`
- Reset: `\033[0m`

**256-Color Fallback:**
- Map RGB to nearest xterm-256 color index
- Use `\033[38;5;%dm` and `\033[48;5;%dm`

### Color Block Rendering
- Calculate block size based on terminal width
- Ensure minimum block size (e.g., 6x3 characters)
- Render hex codes above or inside blocks
- Handle terminal resizing (SIGWINCH)

### Performance Optimization
- Cache rendered swatches
- Batch escape sequence output
- Use buffer pooling for large outputs
- Minimize flicker with double buffering

---

## 8. EXPORT FUNCTIONALITY

### Supported Formats

**JSON (Default):**
```json
{
  "name": "Dark Moody Palette",
  "mood": "dark",
  "scheme": "analogous",
  "colors": ["#1A1A2E", "#16213E", "#0F3460"],
  "metadata": {
    "generated": "2026-06-26T10:00:00Z",
    "total_colors": 3
  }
}
```

**CSS Variables:**
```css
:root {
  --color-1: #1A1A2E;
  --color-2: #16213E;
  --color-3: #0F3460;
}
```

**SCSS/Sass Map:**
```scss
$palette: (
  'dark': (
    'color-1': #1A1A2E,
    'color-2': #16213E
  )
);
```

**YAML:**
```yaml
palette:
  mood: dark
  colors:
    - "#1A1A2E"
    - "#16213E"
```

**Plain HEX:**
```
#1A1A2E
#16213E
#0F3460
```

**Tailwind Config:**
```javascript
module.exports = {
  theme: {
    extend: {
      colors: {
        palette: {
          1: '#1A1A2E',
          2: '#16213E'
        }
      }
    }
  }
}
```

### Clipboard Integration
- Use `xclip` (Linux), `pbcopy` (macOS), or `clip` (Windows)
- Copy in multiple formats (JSON, HEX, etc.)
- Fallback to displaying copy instructions

---

## 9. CONFIGURATION MANAGEMENT

### Config File Location
- `~/.palette/config.yaml`
- `$XDG_CONFIG_HOME/palette/config.yaml` (Linux)
- Cross-platform handling with `os.UserConfigDir()`

### Config Structure
```yaml
defaults:
  mood: "dark"
  count: 5
  scheme: "analogous"
  accessibility: false

preferences:
  color_format: "hex"  # hex, rgb, hsl
  terminal_color: "auto"  # auto, truecolor, 256
  clipboard_format: "json"
  history_enabled: true
  history_size: 20

export:
  default_format: "json"
  include_metadata: true

tui:
  show_help_on_start: false
  color_preview_size: "large"
  animation_enabled: true
```

### Configuration Loading Priority
1. Command-line flags (highest)
2. Environment variables (`PALETTE_MOOD`, `PALETTE_COUNT`)
3. Config file
4. Defaults (lowest)

---

## 10. ERROR HANDLING & VALIDATION

### Input Validation
- Color codes must be valid hex (#RRGGBB)
- Count between 2-10 colors
- Valid mood names
- Valid harmony schemes
- Valid export formats

### Graceful Failure
- Terminal unsupported: fallback to no-color output
- Missing dependencies (xclip): show installation instructions
- Invalid config: warn and use defaults
- Panic recovery in TUI (don't crash on invalid state)

### User Feedback
- Error messages with suggestions
- Warning for low contrast ratios
- Info messages for actions (e.g., "Copied #1A1A2E to clipboard")
- Progress indicators for expensive operations

---

## 11. TESTING STRATEGY

### Unit Tests
- Color conversion functions (RGB ↔ HSL ↔ HEX)
- Harmony generation algorithms (deterministic outputs)
- Mood preset validation
- Export format generation

### Integration Tests
- CLI flag parsing and validation
- End-to-end generation commands
- Config file loading
- Terminal capability detection

### Manual Testing
- TUI interactions (keyboard navigation)
- Visual inspection of colors
- Cross-terminal compatibility (iTerm, Terminal.app, Alacritty, etc.)
- Cross-platform testing (Linux, macOS, Windows)

---

## 12. PERFORMANCE CONSIDERATIONS

### Optimization Targets
- Palette generation: < 50ms
- TUI render: < 16ms (60fps)
- Export: < 100ms for 100 colors
- Startup time: < 100ms

### Memory Usage
- Avoid storing full history in memory for large sessions
- Use streaming for large exports
- Cache frequently used calculations (e.g., nearest 256-color)

### Concurrency
- Use goroutines for clipboard operations
- Parallel generation for multiple palettes
- Background calculations for contrast ratios

---

## 13. DEPENDENCY MANAGEMENT

### Standard Library Usage
- `flag` package for CLI parsing (or `cobra`)
- `fmt` for output formatting
- `os` for file operations
- `time` for timestamps
- `encoding/json` for JSON export
- `image/color` for color handling

### External Libraries (select one from each category)

**CLI Framework:**
- `cobra` (standard, feature-rich)
- `urfave/cli` (simpler alternative)

**TUI Framework:**
- `bubbletea` (primary recommendation)
- `tview` (alternative)

**Color Manipulation:**
- `go-colorful` (primary, comprehensive)

**Terminal Rendering:**
- `charmbracelet/lipgloss` (style composition)
- `termenv` (ANSI escape generation)

**Configuration:**
- `spf13/viper` (full-featured config manager)

**Miscellaneous:**
- `atotto/clipboard` (cross-platform clipboard)
- `mattn/go-runewidth` (Unicode width handling)

---

## 14. RELEASE & DISTRIBUTION

### Build Process
- Use `go build` with optimization flags (`-ldflags="-s -w"`)
- Cross-compile for multiple platforms:
  ```
  GOOS=linux GOARCH=amd64 go build
  GOOS=darwin GOARCH=arm64 go build
  GOOS=windows GOARCH=amd64 go build
  ```

### Installation Methods
1. **Pre-built binaries**: Provide downloads for each platform
2. **Homebrew**: Create a formula for macOS
3. **Go install**: `go install github.com/username/palette@latest`
4. **Package managers**: Add to APT, Snap, or Chocolatey

### CI/CD Pipeline
- GitHub Actions or GitLab CI
- Run tests on each commit
- Build binaries for all platforms on tag
- Auto-release to GitHub Releases

### Documentation
- `README.md`: Installation, usage examples
- `--help` flag: Full command documentation
- Man pages: `man palette`
- Example output: Show generated palettes

---

## 15. FUTURE ENHANCEMENTS (Phase 2)

### Advanced Features
- **Image extraction**: Generate palette from images (`palette extract image.jpg`)
- **Color blindness simulation**: Preview palettes for protanopia/deuteranopia
- **Gradient generation**: Create smooth color gradients
- **Semantic naming**: Name colors (e.g., "Midnight Blue")
- **Web import**: Import palettes from Coolors, Adobe Color, etc.
- **VSCode integration**: Export as VS Code theme
- **API mode**: Serve palettes via HTTP (decorative)
- **AI generation**: Use ML for more sophisticated palettes
- **Palette search**: Save and search previously generated palettes

### Plugin System
- Custom harmony algorithms
- Custom export formats
- Custom mood definitions

### Interactive Enhancements
- Mouse support in TUI
- Live preview of color changes
- Undo/redo stack
- Palette history browser

---

## 16. DEVELOPMENT PHASES

### Phase 1: MVP (Week 1-2)
- [ ] Basic CLI with generate command
- [ ] RGB ↔ HSL conversions
- [ ] Simple random palette generation
- [ ] Terminal output with color blocks
- [ ] Export to HEX and JSON

### Phase 2: Core Features (Week 3-4)
- [ ] All harmony algorithms
- [ ] Mood presets
- [ ] Color locking
- [ ] Export to CSS, YAML
- [ ] Clipboard support

### Phase 3: Interactive TUI (Week 5-6)
- [ ] Bubble Tea implementation
- [ ] Navigation and focusing
- [ ] Regenerate individual colors
- [ ] Locking/unlocking
- [ ] Mood/scheme switching

### Phase 4: Polish & Release (Week 7-8)
- [ ] Config file management
- [ ] Comprehensive testing
- [ ] Cross-platform builds
- [ ] Documentation
- [ ] Release packaging

### Phase 5: Advanced Features (Optional)
- [ ] Image extraction
- [ ] Color blindness simulation
- [ ] Web import/export
- [ ] AI generation


