package export

import (
	"strings"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

func ToHexLines(p color.Palette) string {
	var b strings.Builder
	for _, c := range p.Colors {
		b.WriteString(c.HexString())
		b.WriteByte('\n')
	}
	return b.String()
}
