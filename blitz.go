package blitz

import "github.com/benetis/blitz/slice"

func Map[A any, B any](inputs []A, fn func(A) B, maxWorkers ...int) []B {
	return slice.Map(inputs, fn, maxWorkers...)
}

func Filter[T any](inputs []T, fn func(T) bool, maxWorkers ...int) []T {
	return slice.Filter(inputs, fn, maxWorkers...)
}
