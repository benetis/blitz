package slice

import (
	"reflect"
	"strings"
	"testing"
)

func TestMapInt(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	result := Map(inputs, func(input int) int {
		return input * 2
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v; want %v", result, expected)
	}
}

func TestMapString(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	expected := []string{"A", "B", "C"}

	result := Map(inputs, func(input string) string {
		return strings.ToUpper(input)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v; want %v", result, expected)
	}
}

func TestMapWithMaxWorkers(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	result := Map(inputs, func(input int) int {
		return input * 2
	}, 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v; want %v", result, expected)
	}
}
