package tui

import (
	"sort"

	"github.com/Ronak-jain-afk/palette/internal/color"
	"github.com/Ronak-jain-afk/palette/internal/display"
	"github.com/charmbracelet/bubbletea"
)

type model struct {
	palette    color.Palette
	focusIndex int
	mode       string // "view" or "export"
	showHelp   bool
	renderer   *display.Renderer

	moodKeys  []string
	schemeKeys []string
	moodIdx   int
	schemeIdx int
}

func InitialModel(noColor bool) model {
	keys := make([]string, 0, len(color.Presets))
	for k := range color.Presets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	schemes := []string{"analogous", "complementary", "triadic", "tetradic", "monochromatic"}

	m := model{
		palette:    color.GenerateFromMood("dark", 5, "analogous"),
		focusIndex: 0,
		mode:       "view",
		renderer:   display.NewRenderer(noColor),
		moodKeys:   keys,
		schemeKeys: schemes,
		moodIdx:    indexOf("dark", keys),
		schemeIdx:  0,
	}
	return m
}

func indexOf(s string, xs []string) int {
	for i, x := range xs {
		if x == s {
			return i
		}
	}
	return 0
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}
