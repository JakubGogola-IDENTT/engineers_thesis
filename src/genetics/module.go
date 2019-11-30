package genetics

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// Mutate mutates random rectangle in speciment image
func (s *Specimen) Mutate() {
	if rand.Float64() > s.config.MutationChance {
		return
	}

	// get bounds of random rectangle
	randomRect := GetRandomRectBounds(s.Spec.Bounds())

	// get random color for rectangle
	var randomColor color.Color

	if s.config.GrayScale {
		randomColor = GetRandomColor(true)
	} else {
		randomColor = GetRandomColor(false)
	}

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
			score += math.Pow(CompareColors(s.Spec.At(x, y), originalImage.At(x, y), s.config.GrayScale), 2.0)
		}
	}

	s.Score = score
}

// Cross crosses speciment with antoher image
func (s *Specimen) Cross(spec *Specimen) {
	// // get random part of first image
	// randomRect := GetRandomRectBounds(s.Spec.Bounds())

	// minPixOffset := s.Spec.PixOffset(randomRect.Min.X, randomRect.Min.Y)
	// maxPixOffset := s.Spec.PixOffset(randomRect.Max.X, randomRect.Max.Y)

	// // copy pixels of random area chosen above
	// tmp := s.Spec.Pix[minPixOffset:maxPixOffset]

	// // TODO: ugly and slow solution but, probaby, works
	// for pix := range tmp {
	// 	spec.P
	// }

}
