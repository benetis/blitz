package slice

func assignWorkers(maxWorkers []int, inputsSize int) int {
	numWorkers := defaultWorkers
	if len(maxWorkers) > 0 && maxWorkers[0] > 0 {
		numWorkers = maxWorkers[0]
	}

	if numWorkers > inputsSize {
		numWorkers = inputsSize
	}
	return numWorkers
}

// 'Ceil' division to ensure all inputs are processed
// ceil(x/y)=(x+y-1)/y, where ceil(x/y) is calcChunkSize, x is len(inputs), y is numWorkers
func calcChunkSize(inputSize int, numWorkers int) int {
	return (inputSize + numWorkers - 1) / numWorkers
}
