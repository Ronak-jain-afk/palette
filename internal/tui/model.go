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
	showHelp   bool
	renderer   *display.Renderer

	moodKeys   []string
	schemeKeys []string
	moodIdx    int
	schemeIdx  int
}

func InitialModel(noColor bool) model {
	keys := make([]string, 0, len(color.Presets))
	for k := range color.Presets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	schemes := []string{"analogous", "complementary", "triadic", "tetradic", "monochromatic"}

	moodIdx := 0
	for i, k := range keys {
		if k == "dark" {
			moodIdx = i
			break
		}
	}
	m := model{
		palette:    color.GenerateFromMood("dark", 5, "analogous"),
		focusIndex: 0,
		renderer:   display.NewRenderer(noColor),
		moodKeys:   keys,
		schemeKeys: schemes,
		moodIdx:    moodIdx,
		schemeIdx:  0,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}
