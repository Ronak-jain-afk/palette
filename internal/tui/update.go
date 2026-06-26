package tui

import (
	"github.com/Ronak-jain-afk/palette/internal/color"
	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
	"github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case tea.WindowSizeMsg:
		return m, nil
	}
	return m, nil
}

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case "export":
		return m.handleExportKey(msg)
	default:
		return m.handleViewKey(msg)
	}
}

func (m model) handleViewKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "left":
		m.focusIndex--
		if m.focusIndex < 0 {
			m.focusIndex = len(m.palette.Colors) - 1
		}
	case "right":
		m.focusIndex++
		if m.focusIndex >= len(m.palette.Colors) {
			m.focusIndex = 0
		}

	case " ":
		m.palette.Locked[m.focusIndex] = !m.palette.Locked[m.focusIndex]

	case "r":
		m.regenerateFocused()

	case "R":
		m.regenerateAll()

	case "up":
		m.adjustLightness(0.05)
	case "down":
		m.adjustLightness(-0.05)

	case "[":
		m.shiftHue(-5)
	case "]":
		m.shiftHue(5)
	case "{":
		m.shiftHue(-30)
	case "}":
		m.shiftHue(30)

	case "<", ",":
		m.adjustSaturation(-0.05)
	case ">", ".":
		m.adjustSaturation(0.05)

	case "m":
		m.cycleMood(1)
	case "s":
		m.cycleScheme(1)

	case "e":
		m.mode = "export"

	case "h", "?":
		m.showHelp = !m.showHelp
	}

	return m, nil
}

func (m model) handleExportKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c", "e", "esc":
		m.mode = "view"
	}
	return m, nil
}

func (m *model) regenerateFocused() {
	if m.focusIndex >= len(m.palette.Colors) {
		return
	}
	if m.palette.Locked[m.focusIndex] {
		return
	}
	c := m.palette.Colors[m.focusIndex]
	h, s, l := colorspace.RGBToHSL(c.R, c.G, c.B)
	h = h + 60
	if h >= 360 {
		h -= 360
	}
	r, g, b := colorspace.HSLToRGB(h, s, l)
	m.palette.Colors[m.focusIndex] = color.Color{R: r, G: g, B: b}
}

func (m *model) regenerateAll() {
	mood := m.moodKeys[m.moodIdx]
	scheme := m.schemeKeys[m.schemeIdx]
	p := color.GenerateFromMood(mood, len(m.palette.Colors), scheme)
	for i, c := range p.Colors {
		if i < len(m.palette.Locked) && m.palette.Locked[i] {
			continue
		}
		if i < len(m.palette.Colors) {
			m.palette.Colors[i] = c
		} else {
			m.palette.Colors = append(m.palette.Colors, c)
		}
	}
	m.palette.Mood = p.Mood
	m.palette.Scheme = p.Scheme
	m.palette.Name = p.Name
}

func (m *model) adjustLightness(delta float64) {
	if m.focusIndex >= len(m.palette.Colors) {
		return
	}
	c := m.palette.Colors[m.focusIndex]
	h, s, l := colorspace.RGBToHSL(c.R, c.G, c.B)
	l += delta
	if l < 0 {
		l = 0
	}
	if l > 1 {
		l = 1
	}
	r, g, b := colorspace.HSLToRGB(h, s, l)
	m.palette.Colors[m.focusIndex] = color.Color{R: r, G: g, B: b}
}

func (m *model) shiftHue(delta float64) {
	if m.focusIndex >= len(m.palette.Colors) {
		return
	}
	c := m.palette.Colors[m.focusIndex]
	h, s, l := colorspace.RGBToHSL(c.R, c.G, c.B)
	h += delta
	for h < 0 {
		h += 360
	}
	for h >= 360 {
		h -= 360
	}
	r, g, b := colorspace.HSLToRGB(h, s, l)
	m.palette.Colors[m.focusIndex] = color.Color{R: r, G: g, B: b}
}

func (m *model) adjustSaturation(delta float64) {
	if m.focusIndex >= len(m.palette.Colors) {
		return
	}
	c := m.palette.Colors[m.focusIndex]
	h, s, l := colorspace.RGBToHSL(c.R, c.G, c.B)
	s += delta
	if s < 0 {
		s = 0
	}
	if s > 1 {
		s = 1
	}
	r, g, b := colorspace.HSLToRGB(h, s, l)
	m.palette.Colors[m.focusIndex] = color.Color{R: r, G: g, B: b}
}

func (m *model) cycleMood(dir int) {
	m.moodIdx += dir
	if m.moodIdx < 0 {
		m.moodIdx = len(m.moodKeys) - 1
	}
	if m.moodIdx >= len(m.moodKeys) {
		m.moodIdx = 0
	}
	m.regenerateAll()
}

func (m *model) cycleScheme(dir int) {
	m.schemeIdx += dir
	if m.schemeIdx < 0 {
		m.schemeIdx = len(m.schemeKeys) - 1
	}
	if m.schemeIdx >= len(m.schemeKeys) {
		m.schemeIdx = 0
	}
	m.regenerateAll()
}
