package export

import (
	"github.com/Ronak-jain-afk/palette/internal/color"
	"gopkg.in/yaml.v3"
)

type yamlPalette struct {
	Name   string   `yaml:"name"`
	Mood   string   `yaml:"mood"`
	Scheme string   `yaml:"scheme"`
	Colors []string `yaml:"colors"`
}

func ToYAML(p color.Palette, includeMetadata bool) (string, error) {
	hexes := hexSlice(p)
	if !includeMetadata {
		b, err := yaml.Marshal(hexes)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	yp := yamlPalette{
		Name:   p.Name,
		Mood:   p.Mood,
		Scheme: p.Scheme,
		Colors: hexes,
	}

	b, err := yaml.Marshal(yp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
