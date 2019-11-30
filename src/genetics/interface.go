package genetics

import (
	"image"
	"thesis/config"
)

// Genetics is interface for genetics module
type Genetics interface {
	Mutate()
	Fitness(originalImage image.Image)
	Cross(spec image.RGBA)
}

// Specimen represents set of specimens (images) in single iteration (= generation)
type Specimen struct {
	Spec   image.RGBA
	Score  float64
	config config.Config
}

// DNA contains config, speciments set and root image for algorithm
type DNA struct {
	originalImage       image.Image
	specimens           []Specimen
	nextGenerationSpecs []Specimen
	randomlySelected    []int
	crossed             []Specimen
	config              config.Config
}
