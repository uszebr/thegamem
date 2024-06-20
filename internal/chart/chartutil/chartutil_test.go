package chartutil

import (
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
