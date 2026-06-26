package colorspace

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RGBToHSL(r, g, b uint8) (h, s, l float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	delta := max - min

	l = (max + min) / 2.0

	if delta == 0 {
		return 0, 0, l
	}

	s = delta / (1 - math.Abs(2*l-1))

	switch max {
	case rf:
		h = math.Mod((gf-bf)/delta, 6.0)
	case gf:
		h = (bf-rf)/delta + 2
	case bf:
		h = (rf-gf)/delta + 4
	}

	h *= 60.0
	if h < 0 {
		h += 360.0
	}

	return h, s, l
}

func HSLToRGB(h, s, l float64) (uint8, uint8, uint8) {
	if s == 0 {
		v := uint8(math.Round(l * 255.0))
		return v, v, v
	}

	h = math.Mod(h, 360.0)
	if h < 0 {
		h += 360.0
	}

	var q float64
	if l < 0.5 {
		q = l * (1.0 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q
	hk := h / 360.0

	tr := hk + 1.0/3.0
	tg := hk
	tb := hk - 1.0/3.0

	r := hueToRGB(p, q, tr)
	g := hueToRGB(p, q, tg)
	b := hueToRGB(p, q, tb)

	return uint8(math.Round(r * 255.0)), uint8(math.Round(g * 255.0)), uint8(math.Round(b * 255.0))
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1.0
	}
	if t > 1.0 {
		t -= 1.0
	}
	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6.0*t
	case t < 0.5:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*6.0*(2.0/3.0-t)
	default:
		return p
	}
}

func HexToRGB(hex string) (uint8, uint8, uint8, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	return uint8(r), uint8(g), uint8(b), nil
}

func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
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
