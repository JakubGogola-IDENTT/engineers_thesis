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

// Speciment represents set of speciments (images) in single iteration (= generation)
type Speciment struct {
	Spec  image.RGBA
	Score float64
}

// DNA contains config, speciments set and root image for algorithm
type DNA struct {
	originalImage image.Image
	speciments    []Speciment
	config        config.Config
}
