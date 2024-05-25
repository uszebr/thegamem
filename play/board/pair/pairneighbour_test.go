package pair

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePairsNeighbours(t *testing.T) {
	tests := []struct {
		cols   int
		rows   int
		result int
	}{
		//if board wide/tall more then 2 rounds = cols*rows * 4
		{cols: 1, rows: 1, result: 0}, // no Rounds
		{cols: 1, rows: 2, result: 1}, // no Rounds
		{cols: 2, rows: 2, result: 6}, // 6 rounds
		{cols: 6, rows: 4, result: 96},
		{cols: 5, rows: 3, result: 60},
		{cols: 15, rows: 4, result: 240},
		//tofail{cols: 15, rows: 4, result: 120},
	}
	for index, testl := range tests {
		test := testl // for multithreading safe before 1.22
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			pairNeighbour := new(PairsNeighbour)
			pairs, err := pairNeighbour.CreatePairs(test.cols, test.rows)
			assert.Nil(t, err, "Creating neighbour pairs")
			assert.Len(t, pairs, test.result)

		})
	}
}
