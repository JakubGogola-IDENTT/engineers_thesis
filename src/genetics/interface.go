package genetics

import "image"

// Speciment represents image as speciment of generation
type Speciment struct {
	file image.Image
}

// Genetics is interface for genetics module
type Genetics interface {
	mutate()
	score()
	cross()
}
