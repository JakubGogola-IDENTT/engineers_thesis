package genetics

import "image"

// Genetics is interface for genetics module
type Genetics interface {
	mutate()
	fitness()
	cross()
}

// Generation represents set of speciments (images) in single iteration (= generation)
type Generation struct {
	speciments []image.Image
}
