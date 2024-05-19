package slice

import (
	"reflect"
	"testing"
)

func TestFilterPositiveInt(t *testing.T) {
	inputs := []int{-1, -2, 3, 4}
	expected := []int{3, 4}

	result := Filter(inputs, func(input int) bool {
		return input > 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}

func TestFilterEmptySlice(t *testing.T) {
	inputs := []int{}
	expected := []int{}

	result := Filter(inputs, func(input int) bool {
		return input > 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}

func TestFilterAllMatch(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	result := Filter(inputs, func(input int) bool {
		return input > 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}

func TestFilterNoneMatch(t *testing.T) {
	inputs := []int{-1, -2, -3, -4}
	expected := []int{}

	result := Filter(inputs, func(input int) bool {
		return input > 0
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}

func TestFilterStrings(t *testing.T) {
	inputs := []string{"apple", "banana", "cherry", "date"}
	expected := []string{"apple"}

	result := Filter(inputs, func(input string) bool {
		return len(input) == 5
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}
