package genetics

import (
	"image"
	"sync"
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
	Spec  image.RGBA
	Score float64
}

// DNA contains config, speciments set and root image for algorithm
type DNA struct {
	originalImage image.Image
	specimens     []Specimen
	bestSpecs     []Specimen
	crossed       []Specimen
	config        config.Config
	mu            sync.Mutex
}
