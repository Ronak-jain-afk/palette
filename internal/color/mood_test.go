package color

import (
	"testing"
)

func TestGenerateFromMoodCount(t *testing.T) {
	tests := []struct {
		mood   string
		scheme string
		count  int
	}{
		{"dark", "analogous", 5},
		{"dark", "analogous", 3},
		{"dark", "analogous", 8},
		{"vintage", "complementary", 4},
		{"minimal", "triadic", 3},
		{"nature", "tetradic", 4},
		{"pastel", "monochromatic", 6},
	}
	for _, tt := range tests {
		t.Run(tt.mood+"-"+tt.scheme, func(t *testing.T) {
			p := GenerateFromMood(tt.mood, tt.count, tt.scheme)
			if len(p.Colors) != tt.count {
				t.Errorf("GenerateFromMood(%q, %d, %q) returned %d colors; want %d", tt.mood, tt.count, tt.scheme, len(p.Colors), tt.count)
			}
			if len(p.Locked) != tt.count {
				t.Errorf("GenerateFromMood locked length = %d; want %d", len(p.Locked), tt.count)
			}
			if p.Mood != tt.mood {
				t.Errorf("GenerateFromMood mood = %q; want %q", p.Mood, tt.mood)
			}
			if p.Scheme != tt.scheme {
				t.Errorf("GenerateFromMood scheme = %q; want %q", p.Scheme, tt.scheme)
			}
		})
	}
}

func TestGenerateFromMoodInvalidFallback(t *testing.T) {
	p := GenerateFromMood("nonexistent", 5, "analogous")
	if p.Mood != "nonexistent" {
		t.Errorf("mood should be passed through; got %q", p.Mood)
	}
	if len(p.Colors) != 5 {
		t.Errorf("expected 5 colors from fallback; got %d", len(p.Colors))
	}
}

func TestGenerateFromMoodRandoms(t *testing.T) {
	// Run multiple times to ensure no panics or degenerate output
	for i := 0; i < 50; i++ {
		p := GenerateFromMood("pastel", 5, "analogous")
		if len(p.Colors) != 5 {
			t.Fatalf("iteration %d: got %d colors", i, len(p.Colors))
		}
		for j, c := range p.Colors {
			if c.R == 0 && c.G == 0 && c.B == 0 {
				t.Errorf("iteration %d, color %d: black color", i, j)
			}
		}
	}
}

func TestPresetsDefined(t *testing.T) {
	wanted := []string{"dark", "vintage", "minimal", "nature", "pastel"}
	for _, name := range wanted {
		m, ok := Presets[name]
		if !ok {
			t.Errorf("preset %q not found", name)
			continue
		}
		if len(m.HueRange) == 0 {
			t.Errorf("preset %q has no hue ranges", name)
		}
		if m.Saturation[0] < 0 || m.Saturation[0] > 1 || m.Saturation[1] < 0 || m.Saturation[1] > 1 {
			t.Errorf("preset %q saturation out of range: %v", name, m.Saturation)
		}
		if m.Lightness[0] < 0 || m.Lightness[0] > 1 || m.Lightness[1] < 0 || m.Lightness[1] > 1 {
			t.Errorf("preset %q lightness out of range: %v", name, m.Lightness)
		}
	}
}
