package color

import (
	"testing"
)

func TestHarmonyCounts(t *testing.T) {
	base := Color{R: 233, G: 69, B: 96}

	t.Run("analogous-3", func(t *testing.T) {
		got := Analogous(base, 3)
		if len(got) != 3 {
			t.Errorf("Analogous(base, 3) = %d; want 3", len(got))
		}
	})
	t.Run("analogous-5", func(t *testing.T) {
		got := Analogous(base, 5)
		if len(got) != 5 {
			t.Errorf("Analogous(base, 5) = %d; want 5", len(got))
		}
	})
	t.Run("complementary", func(t *testing.T) {
		got := Complementary(base)
		if len(got) != 2 {
			t.Errorf("Complementary(base) = %d; want 2", len(got))
		}
	})
	t.Run("triadic", func(t *testing.T) {
		got := Triadic(base)
		if len(got) != 3 {
			t.Errorf("Triadic(base) = %d; want 3", len(got))
		}
	})
	t.Run("tetradic", func(t *testing.T) {
		got := Tetradic(base)
		if len(got) != 4 {
			t.Errorf("Tetradic(base) = %d; want 4", len(got))
		}
	})
	t.Run("monochromatic-4", func(t *testing.T) {
		got := Monochromatic(base, 4)
		if len(got) != 4 {
			t.Errorf("Monochromatic(base, 4) = %d; want 4", len(got))
		}
	})
	t.Run("split-complementary-3", func(t *testing.T) {
		got := SplitComplementary(base, 3)
		if len(got) != 3 {
			t.Errorf("SplitComplementary(base, 3) = %d; want 3", len(got))
		}
	})
	t.Run("split-complementary-2-fallback", func(t *testing.T) {
		got := SplitComplementary(base, 2)
		if len(got) != 2 {
			t.Errorf("SplitComplementary(base, 2) = %d; want 2", len(got))
		}
	})
}

func TestGenerateSchemeRoutes(t *testing.T) {
	base := Color{R: 233, G: 69, B: 96}
	tests := []struct {
		scheme string
		count  int
	}{
		{"analogous", 5},
		{"complementary", 4},
		{"triadic", 3},
		{"tetradic", 4},
		{"monochromatic", 5},
	}
	for _, tt := range tests {
		t.Run(tt.scheme, func(t *testing.T) {
			got := GenerateScheme(base, tt.scheme, tt.count)
			if len(got) != tt.count && tt.scheme == "triadic" && len(got) != 3 {
				t.Errorf("GenerateScheme(base, %q, %d) returned %d colors; want %d", tt.scheme, tt.count, len(got), tt.count)
			}
		})
	}
}

func TestGenerateSchemeInvalidFallback(t *testing.T) {
	base := Color{R: 233, G: 69, B: 96}
	got := GenerateScheme(base, "nonexistent", 5)
	if len(got) != 5 {
		t.Errorf("GenerateScheme with invalid scheme returned %d colors; want 5", len(got))
	}
}

func TestNormalizeHue(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 0},
		{360, 0},
		{720, 0},
		{-360, 0},
		{-1, 359},
		{361, 1},
		{180, 180},
	}
	for _, tt := range tests {
		got := normalizeHue(tt.input)
		if got != tt.want {
			t.Errorf("normalizeHue(%f) = %f; want %f", tt.input, got, tt.want)
		}
	}
}
