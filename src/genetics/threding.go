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

func (d *DNA) crossWorker(spec1, spec2 *Specimen, wg *sync.WaitGroup) {
	defer wg.Done()
	CrossSpecimens(spec1, spec2)
}

func (d *DNA) dispatchEvolveWorkers() {
	var wg sync.WaitGroup

	for i := range d.specimens {
		wg.Add(1)

		go d.evolveWorker(&d.specimens[i], &wg)
	}

	wg.Wait()
}

func (d *DNA) dispatchCopyWorkers() {
	var wg sync.WaitGroup

	for i := uint(0); i < d.config.SizeOfGeneration; i++ {
		wg.Add(1)

		go d.copyWorker(&d.specimens[i], d.nextGenerationSpecs[uint(i)%d.config.NumOfBest], &wg)
		continue
	}

	wg.Wait()
}

func (d *DNA) dispatchCrossingWorkers() {
	var wg sync.WaitGroup

	for i := uint(0); i < d.config.SizeOfGeneration; i += 2 {
		wg.Add(1)

		go d.crossWorker(&d.specimens[i], &d.specimens[i+1], &wg)
	}

	wg.Wait()
}

func (d *DNA) dispatcher(workerType string) {
	switch workerType {
	case EVOLVE:
		d.dispatchCopyWorkers()
	case CROSS:
		d.dispatchCrossingWorkers()
	default:
		d.dispatchEvolveWorkers()
	}
}
