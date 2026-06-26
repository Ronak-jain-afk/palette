package color

import (
	"math/rand"
	"time"

	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
)

type Color struct {
	R, G, B uint8
}

type Palette struct {
	Name        string
	Colors      []Color
	Mood        string
	Scheme      string
	GeneratedAt time.Time
	Locked      []bool
}

type Config struct {
	DefaultMood   string
	DefaultCount  int
	OutputFormat  string
	TerminalColor bool
}

func (c Color) HexString() string {
	return colorspace.RGBToHex(c.R, c.G, c.B)
}

func (c Color) String() string {
	return c.HexString()
}

func (c Color) ToHSL() (h, s, l float64) {
	return colorspace.RGBToHSL(c.R, c.G, c.B)
}

func GenerateRandom() Color {
	return Color{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
	}
}
