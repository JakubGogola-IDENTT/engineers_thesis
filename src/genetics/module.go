package genetics

import (
	"image"
	"math/rand"
)

// Mutate mutates random rectangle in speciment image
func (s *Speciment) Mutate() {
	// maximal values of x and y in given Image with border
	maxPt := image.Pt(s.Spec.Bounds().Max.X-1, s.Spec.Bounds().Max.Y-1)

	// x and y of left upper corner of mutated rectangle
	minPt := image.Pt(rand.Intn(maxPt.X), rand.Intn(maxPt.Y))

	// width and height of mutated rectangle
	width := rand.Intn(maxPt.X - minPt.X - 1)
	height := rand.Intn(maxPt.Y - minPt.Y + 1)

	for x := minPt.X; x <= minPt.X+width; x++ {
		for y := minPt.Y; y <= minPt.Y+height; y++ {
			// get random color for rectangle
			randomColor := GetRandomColor()

			s.Spec.Set(x, y, MixColors(randomColor, s.Spec.At(x, y))) // TODO: add mixing colors
		}
	}
}

// Fitness return fitness of speciment to original image
func (s *Speciment) Fitness(originalImage image.Image) {
	var score float64

	// bounds of images
	bounds := image.Pt(originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y)

	// Iterate over all pixels of speciment
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			score += CompareColors(s.Spec.At(x, y), originalImage.At(x, y))
		}
	}

	s.Score = 1.0 - score/(float64(bounds.X*bounds.Y*4*256))
}

// Cross crosses speciment with antoher image
func (s *Speciment) Cross(spec image.RGBA) {

}
