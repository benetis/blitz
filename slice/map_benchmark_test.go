package slice

import "testing"

func SequentialMap[T any, R any](inputs []T, fn func(T) R) []R {
	outputs := make([]R, len(inputs))
	for i, input := range inputs {
		outputs[i] = fn(input)
	}
	return outputs
}

func intensiveTask(input int) int {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += input * i
	}
	return sum
}

func nonIntensiveTask(input int) int {
	return input * 2
}

func BenchmarkParallelMapNonIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	for n := 0; n < b.N; n++ {
		Map(inputs, nonIntensiveTask)
	}
}

func BenchmarkSequentialMapNonIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	for n := 0; n < b.N; n++ {
		SequentialMap(inputs, nonIntensiveTask)
	}
}

func BenchmarkParallelMapIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	for n := 0; n < b.N; n++ {
		Map(inputs, intensiveTask)
	}
}

func BenchmarkSequentialMapIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	for n := 0; n < b.N; n++ {
		SequentialMap(inputs, intensiveTask)
	}
}
