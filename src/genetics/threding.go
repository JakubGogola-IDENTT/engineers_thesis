package genetics

import (
	"sync"
)

type specToProcess struct {
	spec chan *Specimen
}

func (d *DNA) worker(spec *Specimen, wg *sync.WaitGroup) {
	defer wg.Done()

	spec.Mutate()
	spec.Fitness(d.originalImage)
}

func (d *DNA) dispatcher() {
	var wg sync.WaitGroup

	for i := range d.specimens {
		wg.Add(1)
		go d.worker(&d.specimens[i], &wg)
	}

	wg.Wait()
}
