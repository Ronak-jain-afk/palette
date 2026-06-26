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
	buildXtermTable()
}

func buildXtermTable() {
	for i := 0; i < 16; i++ {
		var r, g, b uint8
		switch i {
		case 0:
			r, g, b = 0, 0, 0
		case 1:
			r, g, b = 128, 0, 0
		case 2:
			r, g, b = 0, 128, 0
		case 3:
			r, g, b = 128, 128, 0
		case 4:
			r, g, b = 0, 0, 128
		case 5:
			r, g, b = 128, 0, 128
		case 6:
			r, g, b = 0, 128, 128
		case 7:
			r, g, b = 192, 192, 192
		case 8:
			r, g, b = 128, 128, 128
		case 9:
			r, g, b = 255, 0, 0
		case 10:
			r, g, b = 0, 255, 0
		case 11:
			r, g, b = 255, 255, 0
		case 12:
			r, g, b = 0, 0, 255
		case 13:
			r, g, b = 255, 0, 255
		case 14:
			r, g, b = 0, 255, 255
		case 15:
			r, g, b = 255, 255, 255
		}
		xtermColors[i] = [3]uint8{r, g, b}
	}
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
