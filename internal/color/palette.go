package color

import (
	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
)

type Color struct {
	R, G, B uint8
}

type Palette struct {
	Name   string  `json:"name"`
	Colors []Color `json:"colors"`
	Mood   string  `json:"mood"`
	Scheme string  `json:"scheme"`
	Locked []bool  `json:"locked"`
}

func (c Color) HexString() string {
	return colorspace.RGBToHex(c.R, c.G, c.B)
}
