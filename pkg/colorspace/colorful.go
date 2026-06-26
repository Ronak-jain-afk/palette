package colorspace

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

func toColorful(r, g, b uint8) colorful.Color {
	return colorful.Color{R: float64(r) / 255.0, G: float64(g) / 255.0, B: float64(b) / 255.0}
}

func DeltaE(r1, g1, b1, r2, g2, b2 uint8) float64 {
	c1 := toColorful(r1, g1, b1)
	c2 := toColorful(r2, g2, b2)
	return c1.DistanceLab(c2)
}

func Blend(r1, g1, b1, r2, g2, b2 uint8, t float64) (uint8, uint8, uint8) {
	c1 := toColorful(r1, g1, b1)
	c2 := toColorful(r2, g2, b2)
	blended := c1.BlendHcl(c2, t)
	r, g, b := blended.RGB255()
	return r, g, b
}

func linearize(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

func relativeLuminance(r, g, b uint8) float64 {
	rl := linearize(float64(r) / 255.0)
	gl := linearize(float64(g) / 255.0)
	bl := linearize(float64(b) / 255.0)
	return 0.2126*rl + 0.7152*gl + 0.0722*bl
}

func ContrastRatio(r1, g1, b1, r2, g2, b2 uint8) float64 {
	l1 := relativeLuminance(r1, g1, b1)
	l2 := relativeLuminance(r2, g2, b2)
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05)
}
