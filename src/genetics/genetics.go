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
	d.bestSpecs = make([]Specimen, d.config.NumOfBest)
	copy(d.bestSpecs, d.specimens[:d.config.NumOfBest])
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
	}
}

func (d *DNA) evolve() {
	fmt.Printf("Num of generations: %d\n", d.config.NumOfIterations)

	for i := uint(0); i < d.config.NumOfIterations; i++ {
		d.dispatcher()

		d.findBestSpecimens()

		for j := range d.specimens {
			d.specimens[j] = d.copySpecimen(d.bestSpecs[uint(j)%d.config.NumOfBest])
		}

		fmt.Printf("Generation: %d\n", i)
	}

	for i := range d.bestSpecs {
		fileName := fmt.Sprintf("img_%d_best.png", i)

		imageToSave := d.specimens[i].Spec

		d.saveImage(imageToSave, fileName)
	}
}

func (d *DNA) copySpecimen(s Specimen) (spec Specimen) {
	img := *image.NewRGBA(d.originalImage.Bounds())

	maxRect := d.originalImage.Bounds().Max

	for x := 0; x <= maxRect.X; x++ {
		for y := 0; y <= maxRect.Y; y++ {
			img.Set(x, y, s.Spec.At(x, y))
		}
	}

	spec.Spec = img

	return spec
}
