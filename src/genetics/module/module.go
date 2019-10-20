package module

import (
	"image"
	"image/color"
	"math/rand"
)

func mutate(spec image.RGBA) {
	// maximal values of x and y in given Image with border
	borderX := spec.Bounds().Max.X - 1
	borderY := spec.Bounds().Max.Y - 1

	// x and y of left upper corner of mutated rectangle
	minX := rand.Intn(borderX)
	minY := rand.Intn(borderY)

	// width and height of mutated rectangle
	width := rand.Intn(borderX - minX - 1)
	height := rand.Intn(borderY - minY + 1)

	for x := minX; x <= minX+width; x++ {
		for y := minY; y <= minY+height; y++ {
			spec.Set(x, y, getRandomColor())
		}
	}
}

func getRandomColor() color.Color {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	return color.RGBA{R: r, G: g, B: b, A: 255}
}
