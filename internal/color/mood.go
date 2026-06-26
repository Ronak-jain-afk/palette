package color

import (
	"math/rand"
	"time"

	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
)

type Mood struct {
	Name       string
	Saturation [2]float64
	Lightness  [2]float64
	HueRange   [][2]float64
}

var Presets = map[string]Mood{
	"dark": {
		Saturation: [2]float64{0.20, 0.45},
		Lightness:  [2]float64{0.15, 0.40},
		HueRange:   [][2]float64{{200, 280}, {300, 360}, {0, 20}, {170, 200}},
	},
	"vintage": {
		Saturation: [2]float64{0.30, 0.60},
		Lightness:  [2]float64{0.30, 0.60},
		HueRange:   [][2]float64{{10, 40}, {30, 50}, {40, 60}},
	},
	"minimal": {
		Saturation: [2]float64{0.05, 0.20},
		Lightness:  [2]float64{0.30, 0.80},
		HueRange:   [][2]float64{{0, 360}},
	},
	"nature": {
		Saturation: [2]float64{0.40, 0.70},
		Lightness:  [2]float64{0.25, 0.60},
		HueRange:   [][2]float64{{80, 160}, {30, 60}, {190, 230}, {0, 30}},
	},
	"pastel": {
		Saturation: [2]float64{0.20, 0.35},
		Lightness:  [2]float64{0.65, 0.85},
		HueRange:   [][2]float64{{0, 360}},
	},
}

func GenerateFromMood(mood string, count int, scheme string) Palette {
	m, ok := Presets[mood]
	if !ok {
		m = Presets["dark"]
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	hueRange := m.HueRange[rng.Intn(len(m.HueRange))]
	h := hueRange[0] + rng.Float64()*(hueRange[1]-hueRange[0])
	s := m.Saturation[0] + rng.Float64()*(m.Saturation[1]-m.Saturation[0])
	l := m.Lightness[0] + rng.Float64()*(m.Lightness[1]-m.Lightness[0])

	r, g, b := colorspace.HSLToRGB(h, s, l)
	base := Color{r, g, b}

	colors := GenerateScheme(base, scheme, count)

	locked := make([]bool, len(colors))

	return Palette{
		Name:        mood + " " + scheme + " palette",
		Colors:      colors,
		Mood:        mood,
		Scheme:      scheme,
		GeneratedAt: time.Now(),
		Locked:      locked,
	}
}
