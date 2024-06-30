package chartutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIntToString(t *testing.T) {
	tests := []struct {
		input  []int
		output string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 2, 3}, "[1, 2, 3]"},
		{[]int{-1, 0, 1}, "[-1, 0, 1]"},
		{[]int{10, 20, 30}, "[10, 20, 30]"},
	}

	for _, test := range tests {
		result := SliceIntToString(test.input)
		assert.Equal(t, test.output, result, "For input %v, expected %s, but got %s", test.input, test.output, result)
	}
}

func TestSliceBorderToString(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected string
	}{
		{5, 5, "[]"},
		{6, 5, "[]"},
		{5, 6, "[5]"},
		{1, 10, "[1, 2, 3, 4, 5, 6, 7, 8, 9]"},
		{0, 5, "[0, 1, 2, 3, 4]"},
		{-3, 3, "[-3, -2, -1, 0, 1, 2]"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("left=%d right=%d", test.left, test.right), func(t *testing.T) {
			result := SliceBorderToString(test.left, test.right)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSliceStringToString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, "[]"},
		{[]string{"one"}, "['one']"},
		{[]string{"one", "two", "three"}, "['one', 'two', 'three']"},
		{[]string{"a", "b", "c"}, "['a', 'b', 'c']"},
		{[]string{"hello", "world"}, "['hello', 'world']"},
		{[]string{"with", "emtpy", ""}, "['with', 'emtpy', '']"},
	}

	for _, test := range tests {
		result := SliceStringToString(test.input)
		assert.Equal(t, test.expected, result, "should be equal")
	}
}
