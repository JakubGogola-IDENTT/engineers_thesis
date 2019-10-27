package genetics

import (
	"fmt"
	"image"
	"math/rand"
	"os"
)

type specToProcess struct {
	spec chan Specimen
}

func (d *DNA) worker(input chan<- specToProcess, output chan<- Specimen) {
	for {
		// make and send request for speciment to process
		request := specToProcess{make(chan Specimen)}
		input <- request

		// get spec
		spec := <-request.spec

		if spec.Spec.Bounds() == image.Rect(0, 0, 0, 0) {
			continue
		}

		spec.Mutate()
		// spec.Fitness(d.originalImage)
		output <- spec
	}
}

func (d *DNA) dispatcher(requestChan <-chan specToProcess, responseChan <-chan Specimen) {
	// counter of generations
	generation := uint(1)

	var processed []Specimen

	for {
		select {
		case request := <-requestChan:
			request.spec <- d.specimens[rand.Intn(len(d.specimens))]
		case response := <-responseChan:
			processed = append(processed, response)
		default:
		}

		if uint(len(processed)) < 2 {
			continue
		}

		fmt.Printf("Len of processed: %d\n", len(processed))

		processed = processed[:0]

		fmt.Printf("Generation: %d\n", generation)

		// exit when all iterations were done
		if generation == d.config.NumOfIterations {
			os.Exit(1)
		}

		generation++
	}
}
