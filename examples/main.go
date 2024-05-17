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
	})
	fmt.Println("Uppercased:", uppercased)
}
