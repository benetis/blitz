package blitz

import "github.com/benetis/blitz/slice"

// Map I for Input, O for Output
func Map[I any, O any](inputs []I, fn func(I) O) []O {
	return slice.Map(inputs, fn)
}