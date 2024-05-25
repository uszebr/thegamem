package pair

import "github.com/uszebr/thegamem/play/board/coordinate"

// to fit interface for creating pairs
// player should play with all other players on the board
type PairAll struct{}

func (pairAll PairAll) CreatePairs(boardCols, boardRows int) ([]coordinate.PositionPair, error) {
	result := make([]coordinate.PositionPair, 0)
	calculatedPositions := make([]coordinate.Position, 0, boardCols*boardRows)
	for i := 0; i < boardCols; i++ {
		for j := 0; j < boardRows; j++ {
			calculatedPositions = append(calculatedPositions, coordinate.Position{X: i, Y: j})
		}
	}
	for i := 0; i < len(calculatedPositions); i++ {
		for j := i + 1; j < len(calculatedPositions); j++ {
			pair := coordinate.PositionPair{
				Left:  calculatedPositions[i],
				Right: calculatedPositions[j],
			}
			result = append(result, pair)
		}
	}
	return result, nil
}
