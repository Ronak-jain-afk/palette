package color

import (
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

func (c Color) HexString() string {
	return colorspace.RGBToHex(c.R, c.G, c.B)
}
