package tui

import (
	"fmt"
	"strings"

	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true)
	labelStyle  = lipgloss.NewStyle().Faint(true)
	keyStyle    = lipgloss.NewStyle().Bold(true)
)

func (m model) View() string {
	if m.showHelp {
		return m.helpView()
	}

	header := m.renderHeader()
	swatches := m.renderSwatches()
	info := m.renderInfo()
	footer := m.renderFooter()
	status := m.renderStatus()

	return lipgloss.JoinVertical(lipgloss.Top, header, swatches, info, footer, status)
}

func (m model) renderHeader() string {
	mood := m.moodKeys[m.moodIdx]
	scheme := m.schemeKeys[m.schemeIdx]
	return headerStyle.Render(fmt.Sprintf("Palette Generator  %s %s  %s %s", "·", mood, "·", scheme))
}

func (m model) renderSwatches() string {
	var rows []string
	blockW := 14

	for row := 0; row < 3; row++ {
		var line strings.Builder
		for i, c := range m.palette.Colors {
			hex := c.HexString()
			isFocused := i == m.focusIndex
			isLocked := m.palette.Locked[i]

			var content string
			switch row {
			case 0, 2:
				content = strings.Repeat(" ", blockW)
			case 1:
				label := hex
				if isLocked {
					label = "[L] " + hex
				}
				pad := blockW - len(label)
				left := pad / 2
				content = strings.Repeat(" ", left) + label + strings.Repeat(" ", pad-left)
			}

			style := lipgloss.NewStyle().Background(lipgloss.Color(hex))
			if isFocused {
				style = style.Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFFFFF"))
			}

			block := style.Render(content)
			line.WriteString("  ")
			line.WriteString(block)
		}
		rows = append(rows, line.String())
	}
	return strings.Join(rows, "\n") + "\n"
}

func (m model) renderInfo() string {
	total := len(m.palette.Colors)
	locked := 0
	for _, l := range m.palette.Locked {
		if l {
			locked++
		}
	}

	if len(m.palette.Colors) < 2 {
		return labelStyle.Render("no palette")
	}

	c1 := m.palette.Colors[0]
	c2 := m.palette.Colors[1]
	cr := colorspace.ContrastRatio(c1.R, c1.G, c1.B, c2.R, c2.G, c2.B)

	mood := m.moodKeys[m.moodIdx]
	scheme := m.schemeKeys[m.schemeIdx]

	return fmt.Sprintf("Colors %d | Locked %d | Mood %s | Scheme %s | Contrast %.1f:1",
		total, locked, mood, scheme, cr)
}

func (m model) renderFooter() string {
	keys := []string{
		"r Regenerate", "Space Lock", "m Mood",
		"s Scheme", "e Export", "? Help", "q Quit",
	}
	return strings.Join(keys, "  ")
}

func (m model) renderStatus() string {
	if m.focusIndex >= len(m.palette.Colors) {
		return ""
	}
	c := m.palette.Colors[m.focusIndex]
	h, s, l := colorspace.RGBToHSL(c.R, c.G, c.B)
	hex := c.HexString()
	isLocked := m.palette.Locked[m.focusIndex]

	lock := " "
	if isLocked {
		lock = "[L]"
	}

	return fmt.Sprintf("%s %s  RGB(%d,%d,%d)  HSL(%.0f°,%.0f%%,%.0f%%)",
		lock, hex, c.R, c.G, c.B, h, s*100, l*100)
}

func (m model) helpView() string {
	return `HELP

Navigation
  ← →      Focus prev/next color
  Space    Lock/unlock color

Color actions
  r        Regenerate focused
  R        Regenerate all (preserve locks)
  ↑ ↓      Adjust lightness
  [ ]      Shift hue ±5°
  { }      Shift hue ±30°
  < >      Adjust saturation ±5%

Mood & Scheme
  m        Cycle moods
  s        Cycle schemes

Export
  e        Enter export mode

General
  ? / h    Toggle this help
  q        Quit

Press ? or h to close.`
}


