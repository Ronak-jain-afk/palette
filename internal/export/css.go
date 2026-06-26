package export

import (
	"fmt"
	"strings"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

func ToCSS(p color.Palette) string {
	var b strings.Builder
	b.WriteString(":root {\n")
	for i, c := range p.Colors {
		fmt.Fprintf(&b, "  --color-%d: %s;\n", i+1, c.HexString())
	}
	b.WriteString("}\n")
	return b.String()
}
