package genetics

import (
	"sync"
)

type specToProcess struct {
	spec chan *Specimen
}

func (d *DNA) evolveWorker(spec *Specimen, wg *sync.WaitGroup) {
	defer wg.Done()

	spec.Mutate()
	spec.Fitness(d.originalImage)
}

func (d *DNA) copyWorker(spec *Specimen, specToCopy Specimen, wg *sync.WaitGroup) {
	defer wg.Done()
	CopySpecimen(spec, specToCopy)
}

func (d *DNA) dispatchCopyWorkers() {
	var wg sync.WaitGroup

	for i := range d.specimens {
		wg.Add(1)

		go d.evolveWorker(&d.specimens[i], &wg)
	}

	wg.Wait()
}

func (d *DNA) dispatchEvolveWorkers() {
	var wg sync.WaitGroup

	for i := uint(0); i < d.config.SizeOfGeneration; i++ {
		wg.Add(1)

		go d.copyWorker(&d.specimens[i], d.nextGenerationSpecs[uint(i)%d.config.NumOfBest], &wg)
		continue
	}

	wg.Wait()
}

func (d *DNA) dispatcher(workerType string) {
	if workerType == EVOLVE {
		d.dispatchEvolveWorkers()
	} else {
		d.dispatchCopyWorkers()
	}
}
