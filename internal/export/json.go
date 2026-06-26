package export

import (
	"encoding/json"

	"github.com/Ronak-jain-afk/palette/internal/color"
)

type jsonPalette struct {
	Name   string   `json:"name"`
	Mood   string   `json:"mood"`
	Scheme string   `json:"scheme"`
	Colors []string `json:"colors"`
	// ponytail: generated_at and locked omitted from JSON for now — add when consumers need them
}

func ToJSON(p color.Palette, includeMetadata bool) (string, error) {
	hexes := hexSlice(p)
	if !includeMetadata {
		b, err := json.MarshalIndent(hexes, "", "  ")
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	jp := jsonPalette{
		Name:   p.Name,
		Mood:   p.Mood,
		Scheme: p.Scheme,
		Colors: hexes,
	}

	b, err := json.MarshalIndent(jp, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
