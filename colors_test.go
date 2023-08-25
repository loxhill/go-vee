package govee_test

import (
	"testing"

	govee "github.com/loxhill/go-vee"
)

func TestColors(t *testing.T) {
	tests := []struct {
		name     string
		colorFn  func() govee.Color
		expected govee.Color
	}{
		{"Red", govee.Red, govee.Color{255, 0, 0}},
		{"Green", govee.Green, govee.Color{0, 255, 0}},
		{"Blue", govee.Blue, govee.Color{0, 0, 255}},
		{"White", govee.White, govee.Color{255, 255, 255}},
		{"Yellow", govee.Yellow, govee.Color{255, 255, 0}},
		{"Cyan", govee.Cyan, govee.Color{0, 255, 255}},
		{"Magenta", govee.Magenta, govee.Color{255, 0, 255}},
		{"Orange", govee.Orange, govee.Color{255, 165, 0}},
		{"Purple", govee.Purple, govee.Color{128, 0, 128}},
		{"Pink", govee.Pink, govee.Color{255, 192, 203}},
		{"Brown", govee.Brown, govee.Color{165, 42, 42}},
		{"Gold", govee.Gold, govee.Color{255, 215, 0}},
		{"Silver", govee.Silver, govee.Color{192, 192, 192}},
		{"Gray", govee.Gray, govee.Color{128, 128, 128}},
		{"Maroon", govee.Maroon, govee.Color{128, 0, 0}},
		{"Olive", govee.Olive, govee.Color{128, 128, 0}},
		{"GreenYellow", govee.GreenYellow, govee.Color{173, 255, 47}},
		{"Lime", govee.Lime, govee.Color{0, 255, 0}},
		{"Teal", govee.Teal, govee.Color{0, 128, 128}},
		{"Aqua", govee.Aqua, govee.Color{0, 255, 255}},
		{"Navy", govee.Navy, govee.Color{0, 0, 128}},
		{"DarkBlue", govee.DarkBlue, govee.Color{0, 0, 139}},
		{"Indigo", govee.Indigo, govee.Color{75, 0, 130}},
		{"Violet", govee.Violet, govee.Color{238, 130, 238}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			color := test.colorFn()
			if color != test.expected {
				t.Errorf("Unexpected color value: got %v, want %v", color, test.expected)
			}
			if color.R < 0 || color.R > 255 {
				t.Errorf("R value out of range: %d", color.R)
			}
			if color.G < 0 || color.G > 255 {
				t.Errorf("G value out of range: %d", color.G)
			}
			if color.B < 0 || color.B > 255 {
				t.Errorf("B value out of range: %d", color.B)
			}
		})
	}
}
