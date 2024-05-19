package slice

import (
	"testing"
)

func SequentialFilter[T any](inputs []T, fn func(T) bool) []T {
	var outputs []T
	for _, input := range inputs {
		if fn(input) {
			outputs = append(outputs, input)
		}
	}
	return outputs
}

func BenchmarkParallelFilterNonIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	fn := func(input int) bool {
		return input%2 == 0
	}

	for n := 0; n < b.N; n++ {
		Filter(inputs, fn)
	}
}

func BenchmarkSequentialFilterNonIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	fn := func(input int) bool {
		return input%2 == 0
	}

	for n := 0; n < b.N; n++ {
		SequentialFilter(inputs, fn)
	}
}

func BenchmarkParallelFilterIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	fn := func(input int) bool {
		sum := 0
		for i := 0; i < 1000; i++ {
			sum += input * i
		}
		return sum%2 == 0
	}

	for n := 0; n < b.N; n++ {
		Filter(inputs, fn)
	}
}

func BenchmarkSequentialFilterIntensive(b *testing.B) {
	inputs := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		inputs[i] = i
	}

	fn := func(input int) bool {
		sum := 0
		for i := 0; i < 1000; i++ {
			sum += input * i
		}
		return sum%2 == 0
	}

	for n := 0; n < b.N; n++ {
		SequentialFilter(inputs, fn)
	}
}
