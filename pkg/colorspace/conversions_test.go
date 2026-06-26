package colorspace

import (
	"testing"
)

func TestRGBToHex(t *testing.T) {
	tests := []struct {
		r, g, b uint8
		want    string
	}{
		{0, 0, 0, "#000000"},
		{255, 255, 255, "#FFFFFF"},
		{255, 0, 0, "#FF0000"},
		{0, 255, 0, "#00FF00"},
		{0, 0, 255, "#0000FF"},
		{26, 26, 46, "#1A1A2E"},
		{22, 33, 62, "#16213E"},
		{15, 52, 96, "#0F3460"},
		{83, 52, 131, "#533483"},
		{233, 69, 96, "#E94560"},
		{128, 128, 128, "#808080"},
		{192, 192, 192, "#C0C0C0"},
	}
	for _, tt := range tests {
		got := RGBToHex(tt.r, tt.g, tt.b)
		if got != tt.want {
			t.Errorf("RGBToHex(%d,%d,%d) = %s; want %s", tt.r, tt.g, tt.b, got, tt.want)
		}
	}
}

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		hex       string
		r, g, b   uint8
		shouldErr bool
	}{
		{"#000000", 0, 0, 0, false},
		{"#FFFFFF", 255, 255, 255, false},
		{"#FF0000", 255, 0, 0, false},
		{"#00FF00", 0, 255, 0, false},
		{"#0000FF", 0, 0, 255, false},
		{"#1A1A2E", 26, 26, 46, false},
		{"#808080", 128, 128, 128, false},
		{"invalid", 0, 0, 0, true},
		{"#FFF", 0, 0, 0, true},
		{"", 0, 0, 0, true},
	}
	for _, tt := range tests {
		r, g, b, err := HexToRGB(tt.hex)
		if tt.shouldErr {
			if err == nil {
				t.Errorf("HexToRGB(%q) expected error", tt.hex)
			}
			continue
		}
		if err != nil {
			t.Errorf("HexToRGB(%q) unexpected error: %v", tt.hex, err)
			continue
		}
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("HexToRGB(%q) = (%d,%d,%d); want (%d,%d,%d)", tt.hex, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}

func TestRoundTripHex(t *testing.T) {
	tests := []struct {
		r, g, b uint8
	}{
		{0, 0, 0},
		{255, 255, 255},
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{128, 128, 128},
		{26, 26, 46},
		{233, 69, 96},
		{15, 52, 96},
		{200, 100, 50},
		{75, 130, 200},
		{255, 128, 64},
		{192, 64, 128},
		{100, 200, 150},
		{40, 60, 80},
		{220, 220, 220},
		{10, 20, 30},
		{250, 240, 230},
		{80, 160, 240},
		{180, 90, 45},
	}
	for _, tt := range tests {
		hex := RGBToHex(tt.r, tt.g, tt.b)
		r, g, b, err := HexToRGB(hex)
		if err != nil {
			t.Errorf("HexToRGB(%s) error: %v", hex, err)
			continue
		}
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("RoundTrip (%d,%d,%d) -> %s -> (%d,%d,%d)", tt.r, tt.g, tt.b, hex, r, g, b)
		}
	}
}

func TestHSLIdentity(t *testing.T) {
	tests := []struct {
		r, g, b uint8
	}{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{255, 255, 0},
		{0, 255, 255},
		{255, 0, 255},
		{128, 128, 128},
		{0, 0, 0},
		{255, 255, 255},
		{26, 26, 46},
		{233, 69, 96},
		{200, 100, 50},
		{75, 130, 200},
	}
	for _, tt := range tests {
		h, s, l := RGBToHSL(tt.r, tt.g, tt.b)
		r, g, b := HSLToRGB(h, s, l)
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("HSL round-trip (%d,%d,%d) -> (%.2f,%.4f,%.4f) -> (%d,%d,%d); want (%d,%d,%d)",
				tt.r, tt.g, tt.b, h, s, l, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}
