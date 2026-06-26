package display

import (
	"fmt"
	"strings"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

const (
	blockWidth  = 12
	blockHeight = 3
)

func RenderSwatches(p color.Palette, r *Renderer) string {
	if r.NoColor {
		return renderNoColor(p)
	}
	return renderColor(p, r)
}

func renderNoColor(p color.Palette) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Palette: %s\n", p.Name)
	for _, c := range p.Colors {
		fmt.Fprintf(&b, "  %s\n", c.HexString())
	}
	return b.String()
}

func renderColor(p color.Palette, r *Renderer) string {
	var b strings.Builder

	for _, c := range p.Colors {
		bg := r.backgroundEscape(c)
		reset := r.Reset()

		b.WriteString(bg)
		b.WriteString(strings.Repeat(" ", blockWidth))
		b.WriteString(reset)
	}

	b.WriteByte('\n')

	for _, c := range p.Colors {
		bg := r.backgroundEscape(c)
		reset := r.Reset()
		hex := c.HexString()

		padding := blockWidth - len(hex)
		left := padding / 2
		right := padding - left

		b.WriteString(bg)
		b.WriteString(strings.Repeat(" ", left))
		b.WriteString(hex)
		b.WriteString(strings.Repeat(" ", right))
		b.WriteString(reset)
	}

	b.WriteByte('\n')

	for _, c := range p.Colors {
		bg := r.backgroundEscape(c)
		reset := r.Reset()

		b.WriteString(bg)
		b.WriteString(strings.Repeat(" ", blockWidth))
		b.WriteString(reset)
	}

	return b.String()
}
