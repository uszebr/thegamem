package chartutil

import (
	"fmt"
	"strings"
)

func SliceIntToString(nums []int) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, num := range nums {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", num))
	}
	sb.WriteString("]")
	return sb.String()
}

// result "[1, 2, 3, 4, 5, 6, 7, 8, 9]"
func SliceBorderToString(leftIncluded int, rightNotIncluded int) string {
	var result []string
	for i := leftIncluded; i < rightNotIncluded; i++ {
		result = append(result, fmt.Sprintf("%d", i))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

// result "['Team A', 'Team B', 'Team C', 'Team D', 'Team E']"
func SliceStringToString(strs []string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, s := range strs {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("'%s'", s))
	}
	sb.WriteString("]")
	return sb.String()
}
