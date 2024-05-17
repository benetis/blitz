package slice

import (
	"sync"
)

const defaultWorkers = 8

func Map[T any, R any](inputs []T, fn func(T) R, maxWorkers ...int) []R {
	numWorkers := defaultWorkers
	if len(maxWorkers) > 0 && maxWorkers[0] > 0 {
		numWorkers = maxWorkers[0]
	}

	var wg sync.WaitGroup
	outputs := make([]R, len(inputs))
	jobs := make(chan int, len(inputs))

	worker := func() {
		defer wg.Done()
		for i := range jobs {
			outputs[i] = fn(inputs[i])
		}
	}

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker()
	}

	for i := range inputs {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	return outputs
}
