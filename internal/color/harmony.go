package color

import (
	"github.com/Ronak-jain-afk/palette/pkg/colorspace"
)

func normalizeHue(h float64) float64 {
	h = h
	for h < 0 {
		h += 360
	}
	for h >= 360 {
		h -= 360
	}
	return h
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func Analogous(base Color, count int) []Color {
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	step := 30.0 / float64(count-1)
	colors := make([]Color, count)
	for i := range colors {
		offset := step*float64(i) - 15.0
		rh := normalizeHue(h + offset)
		r, g, b := colorspace.HSLToRGB(rh, clamp(s, 0, 1), clamp(l, 0, 1))
		colors[i] = Color{r, g, b}
	}
	return colors
}

func Complementary(base Color) []Color {
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	rh := normalizeHue(h + 180)
	r, g, b := colorspace.HSLToRGB(rh, clamp(s, 0, 1), clamp(l, 0, 1))
	return []Color{base, Color{r, g, b}}
}

func SplitComplementary(base Color, count int) []Color {
	if count < 3 {
		return Complementary(base)
	}
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	colors := make([]Color, count)
	colors[0] = base
	angleStep := 180.0 / float64(count-1)
	for i := 1; i < count; i++ {
		rh := normalizeHue(h + 150.0 + angleStep*float64(i-1))
		r, g, b := colorspace.HSLToRGB(rh, clamp(s, 0, 1), clamp(l, 0, 1))
		colors[i] = Color{r, g, b}
	}
	return colors
}

func Triadic(base Color) []Color {
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	colors := make([]Color, 3)
	for i := range colors {
		rh := normalizeHue(h + float64(i)*120.0)
		r, g, b := colorspace.HSLToRGB(rh, clamp(s, 0, 1), clamp(l, 0, 1))
		colors[i] = Color{r, g, b}
	}
	return colors
}

func Tetradic(base Color) []Color {
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	colors := make([]Color, 4)
	for i := range colors {
		rh := normalizeHue(h + float64(i)*90.0)
		r, g, b := colorspace.HSLToRGB(rh, clamp(s, 0, 1), clamp(l, 0, 1))
		colors[i] = Color{r, g, b}
	}
	return colors
}

func Monochromatic(base Color, count int) []Color {
	h, s, l := colorspace.RGBToHSL(base.R, base.G, base.B)
	colors := make([]Color, count)
	for i := range colors {
		t := float64(i) / float64(count-1)
		sl := clamp(s*(0.3+t*0.7), 0, 1)
		ll := clamp(l*(0.5+t*0.5), 0, 1)
		r, g, b := colorspace.HSLToRGB(h, sl, ll)
		colors[i] = Color{r, g, b}
	}
	return colors
}

var schemeFuncs = map[string]func(Color, int) []Color{
	"analogous":      func(c Color, n int) []Color { return Analogous(c, n) },
	"complementary":  func(c Color, n int) []Color { return SplitComplementary(c, n) },
	"triadic":        func(c Color, n int) []Color { return Triadic(c) },
	"tetradic":       func(c Color, n int) []Color { return Tetradic(c) },
	"monochromatic":  func(c Color, n int) []Color { return Monochromatic(c, n) },
}

func GenerateScheme(base Color, scheme string, count int) []Color {
	fn, ok := schemeFuncs[scheme]
	if !ok {
		fn = schemeFuncs["analogous"]
	}
	return fn(base, count)
}
