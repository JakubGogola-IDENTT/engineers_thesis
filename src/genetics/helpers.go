package genetics

import (
	"image/color"
	"math"
	"math/rand"
)

// GetRandomColor returns random color in RGBA (without transparency)
func GetRandomColor() color.Color {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	return color.RGBA{R: r, G: g, B: b, A: 0} // TODO: check if A value is correct
}

// CompareColors compares color of two pixels
func CompareColors(c1, c2 color.Color) (diff float64) {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	// TODO: implement square root
	diff += math.Abs(float64(r1 - r2))
	diff += math.Abs(float64(g1 - g2))
	diff += math.Abs(float64(b1 - b2))
	diff += math.Abs(float64(a1 - a2))

	return diff
}

// MixColors returns average of two colors
func MixColors(c1, c2 color.Color) (c color.RGBA) {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	// Take average of every channel
	c.R = uint8(r1+r2) >> 1
	c.G = uint8(g1+g2) >> 1
	c.B = uint8(b1+b2) >> 1
	c.A = uint8(a1+a2) >> 1

	return c
}
