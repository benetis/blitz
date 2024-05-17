package slice

import (
	"sync"
)

const maxWorkers = 8

func Map[T any, R any](inputs []T, fn func(T) R) []R {
	var wg sync.WaitGroup
	outputs := make([]R, len(inputs))
	jobs := make(chan int, len(inputs))

	worker := func() {
		defer wg.Done()
		for i := range jobs {
			outputs[i] = fn(inputs[i])
		}
	}

	for w := 0; w < maxWorkers; w++ {
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
