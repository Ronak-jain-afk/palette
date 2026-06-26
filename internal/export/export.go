package export

import (
	"fmt"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

func hexSlice(p color.Palette) []string {
	hexes := make([]string, len(p.Colors))
	for i, c := range p.Colors {
		hexes[i] = c.HexString()
	}
	return hexes
}

func Export(p color.Palette, format string, includeMetadata bool) (string, error) {
	switch format {
	case "json":
		return ToJSON(p, includeMetadata)
	case "css":
		return ToCSS(p), nil
	case "hex":
		return ToHexLines(p), nil
	case "yaml":
		return ToYAML(p, includeMetadata)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
