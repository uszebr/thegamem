package coordinate

import (
	"fmt"
	"strings"
	"testing"
)

func TestCoordinateTransform(t *testing.T) {
	tests := []struct {
		coordinate int
		boardSize  int
		direction  bool
		result     int
	}{

		{coordinate: 2, boardSize: 4, direction: true, result: 3},
		{coordinate: 3, boardSize: 4, direction: true, result: 0},
		{coordinate: 0, boardSize: 4, direction: false, result: 3},
		{coordinate: 0, boardSize: 1, direction: true, result: 0},
		{coordinate: 0, boardSize: 40, direction: false, result: 39},
		{coordinate: 22, boardSize: 40, direction: false, result: 21},
		{coordinate: 22, boardSize: 40, direction: true, result: 23},
		{coordinate: 39, boardSize: 40, direction: true, result: 0},

		//tofail
		//{coordinate: 39, boardSize: 40, direction: true, result: 2},
	}
	for index, testl := range tests {
		test := testl // for multithreading safe before 1.22
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			actual := coordinateTransform(test.coordinate, test.boardSize, test.direction)
			if actual != test.result {
				t.Errorf("i: %d, Actual coordinate: %v not as expected x: %v: ", index, actual, test.result)
			}
		})
	}
}

func TestPositionInPositions(t *testing.T) {
	tests := []struct {
		p         Position
		positions []Position
		result    bool
	}{
		{p: Position{2, 3}, positions: []Position{{4, 5}, {2, 4}, {2, 3}}, result: true},
		{p: Position{2, 3}, positions: []Position{{4, 5}, {2, 4}, {2, 6}}, result: false},
		{p: Position{0, 0}, positions: []Position{{4, 5}, {2, 4}, {2, 6}}, result: false},
		{p: Position{0, 0}, positions: []Position{{4, 5}, {2, 4}, {0, 0}, {2, 6}}, result: true},
		{p: Position{2, 0}, positions: []Position{{2, 0}, {2, 4}, {0, 0}, {2, 6}}, result: true},
		{p: Position{2, 0}, positions: []Position{{6, 4}, {2, 4}, {0, 0}, {2, 6}}, result: false},
		//tofail
		//{p: Position{0, 0}, positions: []Position{{4, 5}, {2, 4}, {0, 0}, {2, 6}}, result: false},
	}
	for index, testl := range tests {
		test := testl // for multithreading safe before 1.22
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			actual := test.p.In(test.positions)
			if actual != test.result {
				t.Errorf("i: %d, Actual result: %v not as expected x: %v: ", index, actual, test.result)
			}
		})
	}
}

func TestPositionsArround(t *testing.T) {
	tests := []struct {
		p               Position
		boardCols       int
		boardRows       int
		someOfPositions []Position
		errMes          string
	}{
		{p: Position{1, 1}, boardCols: 3, boardRows: 3, someOfPositions: []Position{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, errMes: ""},
		{p: Position{2, 2}, boardCols: 3, boardRows: 3, someOfPositions: []Position{{1, 1}, {1, 2}, {1, 0}, {2, 1}, {2, 0}, {0, 2}, {0, 0}, {0, 1}}, errMes: ""},
		{p: Position{0, 0}, boardCols: 5, boardRows: 2, someOfPositions: []Position{{4, 1}, {0, 1}, {1, 1}}, errMes: ""},

		//small board with duplicates
		{p: Position{0, 0}, boardCols: 2, boardRows: 2, someOfPositions: []Position{{0, 1}, {1, 0}, {1, 1}}, errMes: ""},
		{p: Position{1, 0}, boardCols: 2, boardRows: 2, someOfPositions: []Position{{0, 1}, {0, 0}, {1, 1}}, errMes: ""},

		//negative
		{p: Position{0, -1}, boardCols: 5, boardRows: 2, errMes: "players arround issue: y position is wrong:"},
		{p: Position{-4, 10}, boardCols: 5, boardRows: 20, errMes: "players arround issue: x position is wrong:"},
		{p: Position{7, 10}, boardCols: 5, boardRows: 20, errMes: "players arround issue: x position is wrong:"},
		{p: Position{4, 22}, boardCols: 5, boardRows: 20, errMes: "players arround issue: y position is wrong:"},
	}
	for index, testl := range tests {
		test := testl // for multithreading safe before 1.22
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			actual, err := test.p.GetPositinsAround(test.boardCols, test.boardRows)
			if test.errMes != "" {
				if err == nil {
					t.Errorf("i: %d, expecting error %v but not occured ", index, test.errMes)
				}
				if !strings.Contains(err.Error(), test.errMes) {
					t.Errorf("i: %d, expecting error %v but does not contain message %v", index, err, test.errMes)
				}
				return
			}
			if err != nil {
				t.Errorf("i: %d, unexpected error %v ", index, err)
			}

			if len(actual) > 8 {
				t.Errorf("i: %d, actual size larger then 8: %v", index, actual)
			}

			for _, expectedPosition := range test.someOfPositions {
				if !expectedPosition.In(actual) {
					t.Errorf("i: %d, Actual Positions: %v not contain expected: %v: ", index, actual, expectedPosition)
				}
			}
		})

	}
}
