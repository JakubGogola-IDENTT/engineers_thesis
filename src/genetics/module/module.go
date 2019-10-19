package module

import (
	"fmt"
	"image/draw"
	"math/rand"
)

func mutate(spec draw.Image) {
	// maximal values of x and y in given Image with border
	borderX := spec.Bounds().Max.X - 1
	borderY := spec.Bounds().Max.Y - 1

	// x of left upper corner of mutated rectangle
	minX := rand.Intn(borderX)

	// y of left upper corner of mutated rectangle
	minY := rand.Intn(borderY)

	// width of mutated rectangle
	width := rand.Intn(borderX - minX - 1)

	// height of mutated rectangle
	height := rand.Intn(borderY - minY + 1)

	// color of mutated rectangle
	fmt.Println(width, height)

	for x := minX; x <= minX+width; x++ {
		for y := minY; y <= minY+height; y++ {
			spec.Set()
		}
	}
}
