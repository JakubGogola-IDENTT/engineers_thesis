package genetics

import (
	"fmt"
	"image"
	"os"
	"sync"
)

type specToProcess struct {
	spec chan Specimen
}

func (d *DNA) worker(input chan<- specToProcess, output chan<- Specimen) {
	var mu sync.Mutex

	for {
		request := specToProcess{make(chan Specimen)}

		input <- request
		spec := <-request.spec
		mu.Lock()
		if spec.Spec.Bounds() == image.Rect(0, 0, 0, 0) {
			continue
		}

		// process specimen
		spec.Mutate()
		// spec.Fitness(originalImage)

		mu.Unlock()

		// resend specimen
		output <- spec
	}
}

func (d *DNA) dispatcher(requestChan <-chan specToProcess, responseChan <-chan Specimen) {
	// specimens to process
	toProcess := d.specimens
	d.specimens = d.specimens[:0]

	iteration := uint(1)

	for {
		select {
		case request := <-requestChan:
			if len(toProcess) <= 0 {
				request.spec <- Specimen{}
			} else {
				request.spec <- toProcess[0]
				toProcess = toProcess[1:]
			}
		case response := <-responseChan:
			d.specimens = append(d.specimens, response)
		default:
		}

		if uint(len(d.specimens)) < d.config.SizeOfGeneration {
			continue
		}

		// d.findBestSpecimens()

		// for i, spec := range d.specimens {
		// 	spec.Cross(d.bestSpecs[i%int(d.config.NumOfBest)])
		// 	d.specimens[i] = spec
		// }

		// d.specimens = append(d.bestSpecs, d.specimens...)

		fmt.Printf("Generation: %d\n", iteration)
		iteration++

		toProcess = d.specimens
		d.specimens = d.specimens[:0]

		if iteration == d.config.NumOfIterations {
			for i, spec := range d.bestSpecs {
				imageName := fmt.Sprintf("img_%d.png", i)
				d.saveImage(spec.Spec, imageName)
			}

			os.Exit(1)
		}
	}
}
