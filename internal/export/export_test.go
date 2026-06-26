package export

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Ronak-jain-afk/palette/internal/color"
	"gopkg.in/yaml.v3"
)

func testPalette() color.Palette {
	return color.Palette{
		Name:   "test palette",
		Mood:   "dark",
		Scheme: "analogous",
		Colors: []color.Color{
			{R: 233, G: 69, B: 96},
			{R: 26, G: 26, B: 46},
		},
		Locked: []bool{false, false},
	}
}

func TestToJSON(t *testing.T) {
	p := testPalette()

	t.Run("with metadata", func(t *testing.T) {
		out, err := ToJSON(p, true)
		if err != nil {
			t.Fatal(err)
		}
		var parsed map[string]interface{}
		if err := json.Unmarshal([]byte(out), &parsed); err != nil {
			t.Fatalf("invalid JSON: %v\noutput: %s", err, out)
		}
		if parsed["name"] != p.Name {
			t.Errorf("name = %q; want %q", parsed["name"], p.Name)
		}
		colors := parsed["colors"].([]interface{})
		if len(colors) != len(p.Colors) {
			t.Errorf("colors length = %d; want %d", len(colors), len(p.Colors))
		}
		if colors[0] != "#E94560" {
			t.Errorf("first color = %q; want %q", colors[0], "#E94560")
		}
	})

	t.Run("without metadata", func(t *testing.T) {
		out, err := ToJSON(p, false)
		if err != nil {
			t.Fatal(err)
		}
		var parsed []string
		if err := json.Unmarshal([]byte(out), &parsed); err != nil {
			t.Fatalf("invalid JSON: %v\noutput: %s", err, out)
		}
		if len(parsed) != len(p.Colors) {
			t.Errorf("colors length = %d; want %d", len(parsed), len(p.Colors))
		}
	})
}

func TestToCSS(t *testing.T) {
	p := testPalette()
	out := ToCSS(p)

	if !strings.HasPrefix(out, ":root {\n") {
		t.Errorf("CSS should start with :root {; got %q", out[:20])
	}
	if !strings.HasSuffix(strings.TrimSpace(out), "}") {
		t.Errorf("CSS should end with }")
	}
	if !strings.Contains(out, "--color-1: #E94560;") {
		t.Errorf("CSS missing --color-1: #E94560;\n%s", out)
	}
	if !strings.Contains(out, "--color-2: #1A1A2E;") {
		t.Errorf("CSS missing --color-2: #1A1A2E;\n%s", out)
	}
}

func TestToHexLines(t *testing.T) {
	p := testPalette()
	out := ToHexLines(p)

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != len(p.Colors) {
		t.Errorf("hex lines = %d; want %d", len(lines), len(p.Colors))
	}
	if lines[0] != "#E94560" {
		t.Errorf("first line = %q; want %q", lines[0], "#E94560")
	}
	if lines[1] != "#1A1A2E" {
		t.Errorf("second line = %q; want %q", lines[1], "#1A1A2E")
	}
}

func TestToYAML(t *testing.T) {
	p := testPalette()

	t.Run("with metadata", func(t *testing.T) {
		out, err := ToYAML(p, true)
		if err != nil {
			t.Fatal(err)
		}
		var parsed yamlPalette
		if err := yaml.Unmarshal([]byte(out), &parsed); err != nil {
			t.Fatalf("invalid YAML: %v\noutput: %s", err, out)
		}
		if parsed.Name != p.Name {
			t.Errorf("name = %q; want %q", parsed.Name, p.Name)
		}
		if len(parsed.Colors) != len(p.Colors) {
			t.Errorf("colors length = %d; want %d", len(parsed.Colors), len(p.Colors))
		}
	})

	t.Run("without metadata", func(t *testing.T) {
		out, err := ToYAML(p, false)
		if err != nil {
			t.Fatal(err)
		}
		var parsed []string
		if err := yaml.Unmarshal([]byte(out), &parsed); err != nil {
			t.Fatalf("invalid YAML: %v\noutput: %s", err, out)
		}
		if len(parsed) != len(p.Colors) {
			t.Errorf("colors length = %d; want %d", len(parsed), len(p.Colors))
		}
	})
}

func TestExportRouter(t *testing.T) {
	p := testPalette()
	formats := []string{"json", "css", "hex", "yaml"}
	for _, f := range formats {
		t.Run(f, func(t *testing.T) {
			out, err := Export(p, f, true)
			if err != nil {
				t.Fatalf("Export(%q) error: %v", f, err)
			}
			if out == "" {
				t.Errorf("Export(%q) returned empty string", f)
			}
		})
	}

	t.Run("unknown format", func(t *testing.T) {
		_, err := Export(p, "unknown", true)
		if err == nil {
			t.Error("Export with unknown format should return error")
		}
	})
}
