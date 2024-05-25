package pair

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePairsAll(t *testing.T) {
	tests := []struct {
		cols   int
		rows   int
		result int
	}{
		{cols: 1, rows: 1, result: 0}, // no Rounds
		{cols: 1, rows: 2, result: 1},
		{cols: 1, rows: 3, result: 3},
		{cols: 2, rows: 2, result: 6},
		{cols: 2, rows: 3, result: 15},
		{cols: 3, rows: 3, result: 36},
		{cols: 10, rows: 18, result: 16110},
		//tofail{cols: 15, rows: 4, result: 120},
	}
	for index, testl := range tests {
		test := testl // for multithreading safe before 1.22
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			pairAll := new(PairAll)
			pairs, err := pairAll.CreatePairs(test.cols, test.rows)
			assert.Nil(t, err, "Creating All pairs")
			assert.Len(t, pairs, test.result)
		})
	}
}
