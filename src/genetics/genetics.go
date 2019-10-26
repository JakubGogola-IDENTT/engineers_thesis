package genetics

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"thesis/config"
)

// Run - initializes algorithm
func (d *DNA) Run() {
	d.init()
}

func (d *DNA) init() {
	// parse config
	d.config = config.Config{}
	d.config.Init(false)

	// read image
	d.readImage(d.config.PathToImage)

	// init first generation of specimens
	d.specimens = make([]Specimen, d.config.SizeOfGeneration)

	// init specimens in first generation
	d.initFirstGeneration()

	d.evolve()
}

func (d *DNA) findBestSpecimens() {
	d.specimens = sortSpeciments(d.specimens, true)
	d.bestSpecs = d.specimens[:d.config.NumOfBest]

	// tricky!!!!!!
	d.specimens = d.specimens[d.config.NumOfBest:]
}

func (d *DNA) initFirstGeneration() {
	// size of every specimen
	minRect := d.originalImage.Bounds().Min
	maxRect := d.originalImage.Bounds().Max

	c := color.RGBA{255, 255, 255, 255}

	for i := range d.specimens {
		img := image.NewRGBA(image.Rect(minRect.X, minRect.Y, maxRect.X, maxRect.Y))
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)

		d.specimens[i].Spec = *img
	}
}

func (d *DNA) evolve() {
	// evelution loop
	for i := uint(0); i <= d.config.NumOfIterations; i++ {
		for i, spec := range d.specimens {
			spec.Mutate()
			spec.Fitness(d.originalImage)
			d.specimens[i] = spec // TODO: add pointer
		}

		d.findBestSpecimens()

		for i, spec := range d.specimens {
			spec.Cross(d.bestSpecs[i%int(d.config.NumOfBest)])
			d.specimens[i] = spec
		}

		d.specimens = append(d.bestSpecs, d.specimens...)
		fmt.Printf("Generation: %d\n", i)
	}

	for i, spec := range d.bestSpecs {
		imageName := fmt.Sprintf("img_%d.png", i)
		d.saveImage(spec.Spec, imageName)
	}
}
