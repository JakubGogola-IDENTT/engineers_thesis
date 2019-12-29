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

	d.randomlySelected = make([]int, d.config.SizeOfGeneration)

	d.evolve()
}

func (d *DNA) chooseBestSpecimens() {
	d.specimens = sortSpeciments(d.specimens, true)
	d.nextGenerationSpecs = make([]Specimen, d.config.NumOfBest)
	copy(d.nextGenerationSpecs, d.specimens[:d.config.NumOfBest])
}

func (d *DNA) initFirstGeneration() {
	// size of every specimen
	minRect := d.originalImage.Bounds().Min
	maxRect := d.originalImage.Bounds().Max

	c := color.RGBA{0, 0, 0, 255}

	for i := range d.specimens {
		img := image.NewRGBA(image.Rect(minRect.X, minRect.Y, maxRect.X, maxRect.Y))
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)

		d.specimens[i].Spec = *img
		d.specimens[i].config = d.config
	}
}

func (d *DNA) saveImages(iteration uint) {
	for i := uint(0); i < d.config.NumOfBest; i++ {
		fileName := fmt.Sprintf("img_%d_it_%d_best.png", i, iteration)

		imageToSave := d.specimens[i].Spec

		d.saveImage(imageToSave, fileName)
	}
}

func (d *DNA) evolve() {
	fmt.Printf("Num of generations: %d\n", d.config.NumOfIterations)

	for i := uint(0); i < d.config.NumOfIterations; i++ {
		d.dispatcher(EVOLVE)

		d.chooseBestSpecimens()

		d.dispatcher(COPY)

		if d.config.WithCrossing {
			d.dispatcher(CROSS)
		}

		fmt.Printf("Generation: %d\n", i)

		switch i {
		case 10, 50, 100, 500, 1000, 5000, 10000:
			d.saveImages(i)
		}

	}

	d.saveImages(d.config.NumOfIterations)
}
