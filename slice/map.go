package slice

import (
	"sync"
)

const defaultWorkers = 8

func Map[A any, B any](inputs []A, fn func(A) B, maxWorkers ...int) []B {
	numWorkers := assignWorkers(maxWorkers, len(inputs))

	var wg sync.WaitGroup
	outputs := make([]B, len(inputs))

	chunkSize := calcChunkSize(len(inputs), numWorkers)

	worker := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			outputs[i] = fn(inputs[i])
		}
	}

	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := (w + 1) * chunkSize
		if end > len(inputs) {
			end = len(inputs)
		}
		wg.Add(1)
		go worker(start, end)
	}

	wg.Wait()
	return outputs
}
