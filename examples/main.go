package main

import (
	"fmt"
	"strings"

	"github.com/benetis/blitz"
)

func main() {
	inputs := []int{1, 2, 3, 4, 5}

	doubled := blitz.Map(inputs, func(input int) int {
		return input * 2
	})
	fmt.Println("Doubled:", doubled)

	strs := []string{"a", "b", "c"}
	uppercased := blitz.Map(strs, func(input string) string {
		return strings.ToUpper(input)
	}, 2)
	fmt.Println("Uppercased:", uppercased)

	// Filter
	inputs = []int{-1, -2, 3, 4}
	positive := blitz.Filter(inputs, func(input int) bool {
		return input > 0
	})
	fmt.Println("Positive:", positive)

}
