package slice

import "sync"

func Filter[T any](inputs []T, fn func(T) bool, maxWorkers ...int) []T {
	numWorkers := assignWorkers(maxWorkers, len(inputs))

	if numWorkers < 1 {
		return []T{}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	filterFlags := make([]bool, len(inputs))
	chunkSize := calcChunkSize(len(inputs), numWorkers)

	worker := func(start int, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			if fn(inputs[i]) {
				mu.Lock()
				filterFlags[i] = true
				mu.Unlock()
			}
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

	outputs := make([]T, 0, len(inputs))

	for i, flag := range filterFlags {
		if flag {
			outputs = append(outputs, inputs[i])
		}
	}

	return outputs
}
