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
