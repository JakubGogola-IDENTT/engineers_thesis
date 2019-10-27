package genetics

import (
	"image"
	"math"
	"math/rand"
)

// Mutate mutates random rectangle in speciment image
func (s *Specimen) Mutate() {
	if rand.Float32() > 0.3 {
		return
	}

	// get bounds of random rectangle
	randomRect := GetRandomRectBounds(s.Spec.Bounds())

	// get random color for rectangle
	randomColor := GetRandomColor()

	for x := randomRect.Min.X; x <= randomRect.Max.X; x++ {
		for y := randomRect.Min.Y; y <= randomRect.Max.Y; y++ {

			s.Spec.Set(x, y, MixColors(randomColor, s.Spec.At(x, y)))
		}
	}
}

// Fitness return fitness of speciment to original image
func (s *Specimen) Fitness(originalImage image.Image) {
	var score float64

	// bounds of images
	bounds := image.Pt(originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y)

	// Iterate over all pixels of speciment
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			score += math.Pow(CompareColors(s.Spec.At(x, y), originalImage.At(x, y)), 2.0)
		}
	}

	// s.Score = 1.0 - score/(float64(bounds.X*bounds.Y*4*256))

	s.Score = score
}

// Cross crosses speciment with antoher image
func (s *Specimen) Cross(spec Specimen) {

	randomRect := GetRandomRectBounds(s.Spec.Bounds())

	for x := randomRect.Min.X; x <= randomRect.Max.X; x++ {
		for y := randomRect.Min.Y; y <= randomRect.Max.Y; y++ {
			s.Spec.Set(x, y, spec.Spec.At(x, y))
		}
	}

	// // maximal values of x and y in given Image with border
	// maxPt := image.Pt(s.Spec.Bounds().Max.X, s.Spec.Bounds().Max.Y)

	// for x := 0; x < maxPt.X; x++ {
	// 	for y := 0; y < maxPt.Y; y++ {
	// 		s.Spec.Set(x, y, MixColors(s.Spec.At(x, y), spec.Spec.At(x, y)))
	// 	}
	// }
}
