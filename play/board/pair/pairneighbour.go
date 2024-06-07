package pair

import "github.com/uszebr/thegamem/play/board/coordinate"

// to fit interface for creating pairs for neighbour players up down, left right and diagonal
// board considered as endless
type PairsNeighbour struct{}

const (
	descriptionNeighbour = "Pair only neighbour players on the board"
)

func (pairsNeighbour PairsNeighbour) CreatePairs(boardCols, boardRows int) ([]coordinate.PositionPair, error) {
	result := make([]coordinate.PositionPair, 0)
	calculatedPositions := make([]coordinate.Position, 0, boardCols*boardRows)
	for x := range boardCols {
		for y := range boardRows {
			position := coordinate.Position{X: x, Y: y}
			calculatedPositions = append(calculatedPositions, position)
			partners, err := position.GetPositinsAround(boardCols, boardRows)
			if err != nil {
				return []coordinate.PositionPair{}, err
			}
			for _, partner := range partners {
				if !partner.In(calculatedPositions) {
					result = append(result, coordinate.PositionPair{Left: position, Right: partner})
				}
			}
		}
	}
	return result, nil
}

func (pairsNeighbour PairsNeighbour) GetDescription() string {
	return descriptionNeighbour
}
