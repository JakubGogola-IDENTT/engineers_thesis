package genetics

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
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

func (d *DNA) chooseRandomSpecimens() {
	copy(d.nextGenerationSpecs, d.specimens)

	for i := uint(0); i < d.config.SizeOfGeneration; i++ {
		d.randomlySelected[i] = rand.Intn(int(d.config.SizeOfGeneration))
	}

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

func (d *DNA) evolve() {
	fmt.Printf("Num of generations: %d\n", d.config.NumOfIterations)

	for i := uint(0); i < d.config.NumOfIterations; i++ {
		d.dispatcher(COPY)

		if d.config.SelectionType == config.STRONGEST {
			d.chooseBestSpecimens()
		} else {
			d.chooseRandomSpecimens()
		}

		d.dispatcher(EVOLVE)

		fmt.Printf("Generation: %d\n", i)
	}

	for i := uint(0); i < d.config.NumOfBest; i++ {
		fileName := fmt.Sprintf("img_%d_best.png", i)

		imageToSave := d.specimens[i].Spec

		d.saveImage(imageToSave, fileName)
	}
}
