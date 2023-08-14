package govee

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

func Red() Color {
	return Color{255, 0, 0}
}

func Green() Color {
	return Color{0, 255, 0}
}

func Blue() Color {
	return Color{0, 0, 255}
}

func White() Color {
	return Color{255, 255, 255}
}

func Yellow() Color {
	return Color{255, 255, 0}
}

func Cyan() Color {
	return Color{0, 255, 255}
}

func Magenta() Color {
	return Color{255, 0, 255}
}

func Orange() Color {
	return Color{255, 165, 0}
}

func Purple() Color {
	return Color{128, 0, 128}
}

func Pink() Color {
	return Color{255, 192, 203}
}

func Brown() Color {
	return Color{165, 42, 42}
}

func Gold() Color {
	return Color{255, 215, 0}
}

func Silver() Color {
	return Color{192, 192, 192}
}

func Gray() Color {
	return Color{128, 128, 128}
}

func Maroon() Color {
	return Color{128, 0, 0}
}

func Olive() Color {
	return Color{128, 128, 0}
}

func GreenYellow() Color {
	return Color{173, 255, 47}
}

func Lime() Color {
	return Color{0, 255, 0}
}

func Teal() Color {
	return Color{0, 128, 128}
}

func Aqua() Color {
	return Color{0, 255, 255}
}

func Navy() Color {
	return Color{0, 0, 128}
}

func DarkBlue() Color {
	return Color{0, 0, 139}
}

func Indigo() Color {
	return Color{75, 0, 130}
}

func Violet() Color {
	return Color{238, 130, 238}
}
