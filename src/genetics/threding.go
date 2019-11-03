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

func (d *DNA) copier(spec *Specimen, specToCopy Specimen, wg *sync.WaitGroup) {
	defer wg.Done()
	CopySpecimen(spec, specToCopy)
}

func (d *DNA) dispatcher(copy bool) {
	var wg sync.WaitGroup

	for i := range d.specimens {
		wg.Add(1)

		if copy {
			go d.copier(&d.specimens[i], d.bestSpecs[uint(i)%d.config.NumOfBest], &wg)
		} else {
			go d.worker(&d.specimens[i], &wg)
		}
	}

	wg.Wait()
}
