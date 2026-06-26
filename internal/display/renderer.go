package display

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

type Renderer struct {
	SupportsTrueColor bool
	Supports256       bool
	NoColor           bool
}

func NewRenderer(noColor bool) *Renderer {
	r := &Renderer{}
	if noColor {
		r.NoColor = true
		return r
	}

	cterm := os.Getenv("COLORTERM")
	if cterm == "truecolor" || cterm == "24bit" {
		r.SupportsTrueColor = true
		return r
	}

	term := os.Getenv("TERM")
	if strings.Contains(term, "256color") || strings.HasPrefix(term, "xterm") {
		r.Supports256 = true
		return r
	}

	r.NoColor = true
	return r
}

func (r *Renderer) backgroundEscape(c color.Color) string {
	if r.NoColor {
		return ""
	}
	if r.SupportsTrueColor {
		return fmt.Sprintf("\033[48;2;%d;%d;%dm", c.R, c.G, c.B)
	}
	return fmt.Sprintf("\033[48;5;%dm", nearest256(c))
}

func (r *Renderer) Reset() string {
	if r.NoColor {
		return ""
	}
	return "\033[0m"
}

var xtermColors [256][3]uint8

func init() {
	base16 := [...][3]uint8{
		{0, 0, 0}, {128, 0, 0}, {0, 128, 0}, {128, 128, 0},
		{0, 0, 128}, {128, 0, 128}, {0, 128, 128}, {192, 192, 192},
		{128, 128, 128}, {255, 0, 0}, {0, 255, 0}, {255, 255, 0},
		{0, 0, 255}, {255, 0, 255}, {0, 255, 255}, {255, 255, 255},
	}
	copy(xtermColors[:], base16[:])
	for i := 16; i < 232; i++ {
		j := i - 16
		r := uint8(j/36) * 255 / 5
		g := uint8((j%36)/6) * 255 / 5
		b := uint8(j%6) * 255 / 5
		xtermColors[i] = [3]uint8{r, g, b}
	}
	for i := 232; i < 256; i++ {
		v := uint8((i-232)*10 + 8)
		xtermColors[i] = [3]uint8{v, v, v}
	}
}

func nearest256(c color.Color) uint8 {
	best := uint8(0)
	bestDist := float64(3 * 255 * 255)
	for i, xc := range xtermColors {
		dr := int(c.R) - int(xc[0])
		dg := int(c.G) - int(xc[1])
		db := int(c.B) - int(xc[2])
		dist := float64(dr*dr + dg*dg + db*db)
		if dist < bestDist {
			bestDist = dist
			best = uint8(i)
		}
	}
	return best
}
